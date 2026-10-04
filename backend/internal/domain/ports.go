package domain

import "context"

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

type LedgerRepository interface {
	Append(ctx context.Context, entry LedgerEntry) error
}
