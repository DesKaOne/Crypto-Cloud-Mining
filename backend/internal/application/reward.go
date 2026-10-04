package application

import (
	"context"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

// CalculateReward delegates deterministic calculation to a domain-specific policy.
type CalculateReward struct {
	Calculator RewardCalculator
}

func (uc CalculateReward) Execute(ctx context.Context, session domain.MiningSession, periodKey string) (domain.Reward, error) {
	return uc.Calculator.Calculate(ctx, session, periodKey)
}

// PostReward persists the reward identity and posts its financial effect.
// The infrastructure implementation must make the persistence boundary atomic.
type PostReward struct {
	Rewards RewardStore
	Ledger  LedgerPoster
}

func (uc PostReward) Execute(ctx context.Context, reward domain.Reward) error {
	created, err := uc.Rewards.CreateIfAbsent(ctx, reward)
	if err != nil {
		return err
	}
	if !created {
		return nil
	}
	return uc.Ledger.PostReward(ctx, reward)
}
