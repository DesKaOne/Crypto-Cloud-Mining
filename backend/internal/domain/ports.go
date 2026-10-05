package domain

import (
	"context"
	"time"
)

type UserRepository interface {
	GetByID(ctx context.Context, id UserID) (*User, error)
}

type AssetRepository interface {
	GetByID(ctx context.Context, id AssetID) (*Asset, error)
	GetByCode(ctx context.Context, code string) (*Asset, error)
}

type MiningPlanRepository interface {
	GetByID(ctx context.Context, id MiningPlanID) (*MiningPlan, error)
}

type MiningSessionRepository interface {
	GetByID(ctx context.Context, id MiningSessionID) (*MiningSession, error)
}

type MiningIntervalRepository interface {
	Open(ctx context.Context, interval MiningInterval) error
	CloseOpen(ctx context.Context, sessionID MiningSessionID, endedAt time.Time) error
	ListBySession(ctx context.Context, sessionID MiningSessionID) ([]MiningInterval, error)
}

type LedgerRepository interface {
	Append(ctx context.Context, entry LedgerEntry) error
}
