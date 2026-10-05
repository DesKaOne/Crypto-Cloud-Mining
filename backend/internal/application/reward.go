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

// PostReward persists the reward and its financial effect through one atomic boundary.
type PostReward struct {
	Settlement RewardSettlement
}

func (uc PostReward) Execute(ctx context.Context, reward domain.Reward) error {
	return uc.Settlement.SettleReward(ctx, reward)
}
