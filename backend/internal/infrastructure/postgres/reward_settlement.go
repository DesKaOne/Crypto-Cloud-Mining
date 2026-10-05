package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

func (s RewardSettlementStore) Begin(ctx context.Context) (*sql.Tx, error) {
	if s.DB == nil {
		return nil, ErrTransactionBegin
	}
	return s.DB.BeginTx(ctx, nil)
}

type RewardSettlementStore struct {
	DB                     DB
	UserAccount            string
	MiningLiabilityAccount string
}

func (s RewardSettlementStore) SettleReward(ctx context.Context, reward domain.Reward) error {
	if reward.MiningSessionID == "" || reward.AssetID == "" || reward.IdempotencyKey == "" {
		return domain.ErrInvalidRewardInput
	}
	tx, err := s.Begin(ctx)
	if err != nil {
		return err
	}
	defer rollback(tx)

	var rewardID string
	err = tx.QueryRowContext(ctx, "INSERT INTO rewards (mining_session_id, asset_id, quantity, period_key, idempotency_key, policy_version, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7) ON CONFLICT (idempotency_key) DO NOTHING RETURNING id",
		reward.MiningSessionID, reward.AssetID, reward.Quantity, reward.PeriodKey, reward.IdempotencyKey, reward.PolicyVersion, reward.CreatedAt.UTC()).Scan(&rewardID)
	if err == sql.ErrNoRows {
		return tx.Commit()
	}
	if err != nil {
		return fmt.Errorf("insert reward: %w", err)
	}

	if s.UserAccount == "" || s.MiningLiabilityAccount == "" {
		return fmt.Errorf("reward settlement accounts are not configured")
	}
	if _, err = tx.ExecContext(ctx, "INSERT INTO ledger_entries (asset_id, reference_id, debit_account, credit_account, quantity, idempotency_key, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7)",
		reward.AssetID, rewardID, s.MiningLiabilityAccount, s.UserAccount, reward.Quantity, reward.IdempotencyKey, reward.CreatedAt.UTC()); err != nil {
		return fmt.Errorf("insert reward ledger entry: %w", err)
	}
	return tx.Commit()
}
