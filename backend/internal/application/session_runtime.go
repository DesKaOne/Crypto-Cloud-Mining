package application

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

var ErrRewardPeriodAlreadyClosed = errors.New("reward period is already closed")

type SessionRewardInput struct {
	Session       domain.MiningSession
	Intervals     []domain.MiningInterval
	Hashrate      string
	RewardPerHash string
	PolicyVersion int
	PeriodKey     string
}

func CalculateSessionPeriodReward(
	ctx context.Context,
	input SessionRewardInput,
	policy domain.MiningRewardPolicy,
	now time.Time,
) (domain.Reward, error) {
	if err := ctx.Err(); err != nil {
		return domain.Reward{}, err
	}
	if input.PeriodKey == "" || input.Session.ID == "" {
		return domain.Reward{}, domain.ErrInvalidRewardInput
	}
	if input.PolicyVersion != policy.Version {
		return domain.Reward{}, fmt.Errorf("policy version mismatch: session=%d policy=%d", input.PolicyVersion, policy.Version)
	}

	seconds, err := domain.ActiveDurationSeconds(input.Intervals, now)
	if err != nil {
		return domain.Reward{}, err
	}

	quantity, err := policy.CalculateMiningReward(domain.MiningRewardInput{
		Hashrate:         input.Hashrate,
		DurationSeconds:  seconds,
		RewardPerHashSec: input.RewardPerHash,
	})
	if err != nil {
		return domain.Reward{}, err
	}

	key := string(input.Session.ID) + ":" + input.PeriodKey
	return domain.Reward{
		ID:              domain.RewardID(key),
		MiningSessionID: input.Session.ID,
		Quantity:        quantity,
		PeriodKey:       input.PeriodKey,
		IdempotencyKey:  key,
		PolicyVersion:   policy.Version,
	}, nil
}
