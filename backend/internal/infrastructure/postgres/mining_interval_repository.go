package postgres

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

var ErrMiningIntervalNotFound = errors.New("mining interval not found")

type MiningIntervalRepository struct {
	DB *sql.DB
}

func (r MiningIntervalRepository) Open(ctx context.Context, interval domain.MiningInterval) error {
	_, err := r.DB.ExecContext(ctx, "INSERT INTO mining_session_intervals (id, session_id, started_at, ended_at) VALUES ($1, $2, $3, $4)", string(interval.ID), string(interval.SessionID), interval.StartedAt.UTC(), interval.EndedAt)
	return err
}

func (r MiningIntervalRepository) CloseOpen(ctx context.Context, sessionID domain.MiningSessionID, endedAt time.Time) error {
	result, err := r.DB.ExecContext(ctx, "UPDATE mining_session_intervals SET ended_at = $2 WHERE session_id = $1 AND ended_at IS NULL", string(sessionID), endedAt.UTC())
	if err != nil {
		return err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rows != 1 {
		return ErrMiningIntervalNotFound
	}
	return nil
}

func (r MiningIntervalRepository) ListBySession(ctx context.Context, sessionID domain.MiningSessionID) ([]domain.MiningInterval, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id::text, session_id::text, started_at, ended_at FROM mining_session_intervals WHERE session_id = $1 ORDER BY started_at ASC, id ASC", string(sessionID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var intervals []domain.MiningInterval
	for rows.Next() {
		var id, session string
		var startedAt time.Time
		var endedAt *time.Time
		if err := rows.Scan(&id, &session, &startedAt, &endedAt); err != nil {
			return nil, err
		}
		intervals = append(intervals, domain.MiningInterval{
			ID:        domain.MiningIntervalID(id),
			SessionID: domain.MiningSessionID(session),
			StartedAt: startedAt.UTC(),
			EndedAt:   endedAt,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return intervals, nil
}