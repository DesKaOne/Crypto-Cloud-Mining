package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

var ErrTransactionBegin = errors.New("postgres transaction begin failed")

type DB interface {
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
}

type MiningRuntimeStore struct {
	DB DB
}

func (s MiningRuntimeStore) Begin(ctx context.Context) (*sql.Tx, error) {
	if s.DB == nil {
		return nil, ErrTransactionBegin
	}
	return s.DB.BeginTx(ctx, nil)
}

func (s MiningRuntimeStore) GetByID(ctx context.Context, sessionID domain.MiningSessionID) (*domain.MiningSession, error) {
	if sessionID == "" {
		return nil, domain.ErrInvalidMiningInterval
	}
	var session domain.MiningSession
	var startedAt, endedAt sql.NullTime
	err := s.DB.QueryRowContext(ctx, "SELECT id, user_id, mining_plan_id, plan_version, status, started_at, ended_at, created_at FROM mining_sessions WHERE id = $1", sessionID).
		Scan(&session.ID, &session.UserID, &session.MiningPlanID, &session.PlanVersion, &session.Status, &startedAt, &endedAt, &session.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("get mining session: %w", err)
	}
	if startedAt.Valid {
		value := startedAt.Time.UTC()
		session.StartedAt = &value
	}
	if endedAt.Valid {
		value := endedAt.Time.UTC()
		session.EndedAt = &value
	}
	return &session, nil
}

func (s MiningRuntimeStore) ListIntervals(ctx context.Context, sessionID domain.MiningSessionID) ([]domain.MiningInterval, error) {
	rows, err := s.DB.QueryContext(ctx, "SELECT id, mining_session_id, started_at, ended_at FROM mining_intervals WHERE mining_session_id = $1 ORDER BY started_at, id", sessionID)
	if err != nil {
		return nil, fmt.Errorf("list mining intervals: %w", err)
	}
	defer rows.Close()

	var intervals []domain.MiningInterval
	for rows.Next() {
		var interval domain.MiningInterval
		var endedAt sql.NullTime
		if err := rows.Scan(&interval.ID, &interval.SessionID, &interval.StartedAt, &endedAt); err != nil {
			return nil, fmt.Errorf("scan mining interval: %w", err)
		}
		interval.StartedAt = interval.StartedAt.UTC()
		if endedAt.Valid {
			value := endedAt.Time.UTC()
			interval.EndedAt = &value
		}
		intervals = append(intervals, interval)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("iterate mining intervals: %w", err)
	}
	return intervals, nil
}

func (s MiningRuntimeStore) Start(ctx context.Context, session domain.MiningSession, interval domain.MiningInterval) error {
	if session.ID == "" || interval.ID == "" || interval.SessionID != session.ID {
		return domain.ErrInvalidMiningInterval
	}
	tx, err := s.Begin(ctx)
	if err != nil {
		return err
	}
	defer rollback(tx)

	if _, err = tx.ExecContext(ctx, "INSERT INTO mining_sessions (id, user_id, mining_plan_id, plan_version, status, started_at, ended_at, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		session.ID, session.UserID, session.MiningPlanID, session.PlanVersion, session.Status, session.StartedAt, session.EndedAt, session.CreatedAt); err != nil {
		return fmt.Errorf("insert mining session: %w", err)
	}
	if _, err = tx.ExecContext(ctx, "INSERT INTO mining_intervals (id, mining_session_id, started_at, ended_at) VALUES ($1, $2, $3, $4)",
		interval.ID, interval.SessionID, interval.StartedAt.UTC(), interval.EndedAt); err != nil {
		return fmt.Errorf("insert mining interval: %w", err)
	}
	return tx.Commit()
}

func (s MiningRuntimeStore) Pause(ctx context.Context, sessionID domain.MiningSessionID, endedAt time.Time) error {
	if sessionID == "" {
		return domain.ErrInvalidMiningInterval
	}
	tx, err := s.Begin(ctx)
	if err != nil {
		return err
	}
	defer rollback(tx)
	if err = lockSession(ctx, tx, sessionID); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, "UPDATE mining_intervals SET ended_at = $2 WHERE mining_session_id = $1 AND ended_at IS NULL", sessionID, endedAt.UTC()); err != nil {
		return fmt.Errorf("close mining interval: %w", err)
	}
	return tx.Commit()
}

func (s MiningRuntimeStore) Resume(ctx context.Context, sessionID domain.MiningSessionID, interval domain.MiningInterval) error {
	if sessionID == "" || interval.ID == "" || interval.SessionID != sessionID {
		return domain.ErrInvalidMiningInterval
	}
	tx, err := s.Begin(ctx)
	if err != nil {
		return err
	}
	defer rollback(tx)
	if err = lockSession(ctx, tx, sessionID); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, "INSERT INTO mining_intervals (id, mining_session_id, started_at, ended_at) VALUES ($1, $2, $3, NULL)",
		interval.ID, sessionID, interval.StartedAt.UTC()); err != nil {
		return fmt.Errorf("insert resumed mining interval: %w", err)
	}
	return tx.Commit()
}

func (s MiningRuntimeStore) Finish(ctx context.Context, sessionID domain.MiningSessionID, status string, endedAt time.Time) error {
	if sessionID == "" || status == "" {
		return domain.ErrInvalidMiningInterval
	}
	tx, err := s.Begin(ctx)
	if err != nil {
		return err
	}
	defer rollback(tx)
	if err = lockSession(ctx, tx, sessionID); err != nil {
		return err
	}
	if _, err = tx.ExecContext(ctx, "UPDATE mining_intervals SET ended_at = $2 WHERE mining_session_id = $1 AND ended_at IS NULL", sessionID, endedAt.UTC()); err != nil {
		return fmt.Errorf("close final mining interval: %w", err)
	}
	if _, err = tx.ExecContext(ctx, "UPDATE mining_sessions SET status = $2, ended_at = $3, updated_at = $3 WHERE id = $1",
		sessionID, status, endedAt.UTC()); err != nil {
		return fmt.Errorf("finish mining session: %w", err)
	}
	return tx.Commit()
}

func lockSession(ctx context.Context, tx *sql.Tx, sessionID domain.MiningSessionID) error {
	var status string
	if err := tx.QueryRowContext(ctx, "SELECT status FROM mining_sessions WHERE id = $1 FOR UPDATE", sessionID).Scan(&status); err != nil {
		return fmt.Errorf("lock mining session: %w", err)
	}
	return nil
}

func rollback(tx *sql.Tx) {
	_ = tx.Rollback()
}
