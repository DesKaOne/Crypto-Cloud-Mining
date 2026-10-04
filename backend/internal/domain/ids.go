package domain

// Domain IDs are opaque strings at the domain boundary.
// Persistence can use UUID values without coupling domain types to a UUID library.
type UserID string
type AssetID string
type MiningPlanID string
type MiningSessionID string
type RewardID string
type WalletID string
type LedgerEntryID string
type TransactionID string
