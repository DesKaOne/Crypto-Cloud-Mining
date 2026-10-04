package domain

import "time"

type User struct {
	ID        UserID
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Asset struct {
	ID        AssetID
	Code      string
	Name      string
	Decimals  int
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type MiningPlan struct {
	ID          MiningPlanID
	AssetID     AssetID
	Version     int
	Status      string
	CreatedAt   time.Time
	EffectiveAt time.Time
}

type MiningSession struct {
	ID           MiningSessionID
	UserID       UserID
	MiningPlanID MiningPlanID
	PlanVersion  int
	Status       string
	StartedAt    *time.Time
	EndedAt      *time.Time
	CreatedAt    time.Time
}

type Reward struct {
	ID              RewardID
	MiningSessionID MiningSessionID
	AssetID         AssetID
	Quantity        string
	PeriodKey       string
	IdempotencyKey  string
	PolicyVersion   int
	CreatedAt       time.Time
}

type LedgerEntry struct {
	ID             LedgerEntryID
	AssetID        AssetID
	ReferenceID    string
	DebitAccount   string
	CreditAccount  string
	Quantity       string
	IdempotencyKey string
	CreatedAt      time.Time
}
