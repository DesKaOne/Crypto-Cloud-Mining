package application

import (
	"context"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

// Clock makes authoritative server time injectable and testable.
type Clock interface {
	Now() time.Time
}

// MiningPlanReader supplies plan state to use cases.
type MiningPlanReader interface {
	GetByID(ctx context.Context, id domain.MiningPlanID) (*domain.MiningPlan, error)
}

// MiningSessionStore creates and updates mining sessions atomically.
type MiningSessionStore interface {
	Create(ctx context.Context, session domain.MiningSession) error
	GetByID(ctx context.Context, id domain.MiningSessionID) (*domain.MiningSession, error)
	Update(ctx context.Context, session domain.MiningSession) error
}

// RewardCalculator calculates rewards from authoritative server-side inputs.
type RewardCalculator interface {
	Calculate(ctx context.Context, session domain.MiningSession, periodKey string) (domain.Reward, error)
}

// RewardStore persists idempotent reward records.
type RewardStore interface {
	CreateIfAbsent(ctx context.Context, reward domain.Reward) (created bool, err error)
}

// LedgerPoster atomically posts a financial movement.
type LedgerPoster interface {
	PostReward(ctx context.Context, reward domain.Reward) error
}
