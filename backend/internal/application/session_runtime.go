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

// MiningSessionRuntimeStore owns the transaction boundary for session lifecycle changes.
// Implementations must make each operation atomic with its corresponding interval mutation.
type MiningSessionRuntimeStore interface {
	Start(ctx context.Context, session domain.MiningSession, interval domain.MiningInterval) error
	Pause(ctx context.Context, sessionID domain.MiningSessionID, endedAt time.Time) error
	Resume(ctx context.Context, sessionID domain.MiningSessionID, interval domain.MiningInterval) error
	Finish(ctx context.Context, sessionID domain.MiningSessionID, status string, endedAt time.Time) error
}

type SessionRuntimeCommand struct {
	Sessions MiningSessionRuntimeStore
	Clock    Clock
}

func (uc SessionRuntimeCommand) Pause(ctx context.Context, sessionID domain.MiningSessionID) error {
	if sessionID == "" {
		return ErrInvalidMiningTransition
	}
	return uc.Sessions.Pause(ctx, sessionID, uc.Clock.Now().UTC())
}

func (uc SessionRuntimeCommand) Resume(ctx context.Context, sessionID domain.MiningSessionID, intervalID domain.MiningIntervalID) error {
	if sessionID == "" || intervalID == "" {
		return ErrInvalidMiningTransition
	}
	now := uc.Clock.Now().UTC()
	return uc.Sessions.Resume(ctx, sessionID, domain.MiningInterval{ID: intervalID, SessionID: sessionID, StartedAt: now})
}

func (uc SessionRuntimeCommand) Complete(ctx context.Context, sessionID domain.MiningSessionID) error {
	return uc.finish(ctx, sessionID, "completed")
}

func (uc SessionRuntimeCommand) Cancel(ctx context.Context, sessionID domain.MiningSessionID) error {
	return uc.finish(ctx, sessionID, "cancelled")
}

func (uc SessionRuntimeCommand) finish(ctx context.Context, sessionID domain.MiningSessionID, status string) error {
	if sessionID == "" {
		return ErrInvalidMiningTransition
	}
	return uc.Sessions.Finish(ctx, sessionID, status, uc.Clock.Now().UTC())
}

func CalculateSessionPeriodReward(ctx context.Context, input SessionRewardInput, policy domain.MiningRewardPolicy, now time.Time) (domain.Reward, error) {
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
	quantity, err := policy.CalculateMiningReward(domain.MiningRewardInput{Hashrate: input.Hashrate, DurationSeconds: seconds, RewardPerHashSec: input.RewardPerHash})
	if err != nil {
		return domain.Reward{}, err
	}
	key := string(input.Session.ID) + ":" + input.PeriodKey
	return domain.Reward{ID: domain.RewardID(key), MiningSessionID: input.Session.ID, Quantity: quantity, PeriodKey: input.PeriodKey, IdempotencyKey: key, PolicyVersion: policy.Version, CreatedAt: now.UTC()}, nil
}
