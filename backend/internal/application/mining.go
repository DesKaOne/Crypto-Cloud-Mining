package application

import (
	"context"
	"errors"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

var (
	ErrMiningPlanInactive      = errors.New("mining plan is inactive")
	ErrInvalidMiningTransition = errors.New("invalid mining session transition")
)

// StartMiningSession creates an active session using server time.
type StartMiningSession struct {
	Plans    MiningPlanReader
	Sessions MiningSessionStore
	Clock    Clock
}

func (uc StartMiningSession) Execute(ctx context.Context, userID domain.UserID, sessionID domain.MiningSessionID, planID domain.MiningPlanID) error {
	plan, err := uc.Plans.GetByID(ctx, planID)
	if err != nil {
		return err
	}
	if plan.Status != "active" {
		return ErrMiningPlanInactive
	}

	now := uc.Clock.Now().UTC()
	session := domain.MiningSession{
		ID:           sessionID,
		UserID:       userID,
		MiningPlanID: plan.ID,
		PlanVersion:  plan.Version,
		Status:       "active",
		StartedAt:    &now,
		CreatedAt:    now,
	}
	return uc.Sessions.Create(ctx, session)
}

// ChangeMiningSessionStatus centralizes lifecycle validation.
func ChangeMiningSessionStatus(session *domain.MiningSession, next string, now time.Time) error {
	switch session.Status {
	case "active":
		if next != "paused" && next != "completed" && next != "cancelled" {
			return ErrInvalidMiningTransition
		}
	case "paused":
		if next != "active" {
			return ErrInvalidMiningTransition
		}
	case "pending":
		if next != "active" && next != "cancelled" {
			return ErrInvalidMiningTransition
		}
	default:
		return ErrInvalidMiningTransition
	}

	session.Status = next
	now = now.UTC()
	if next == "completed" || next == "cancelled" {
		session.EndedAt = &now
	}
	return nil
}
