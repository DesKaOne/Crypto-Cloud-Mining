package application

import (
	"context"
	"errors"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

const DefaultRewardPeriod = time.Hour

var (
	ErrInvalidRewardPeriod   = errors.New("invalid reward period")
	ErrRewardPeriodNotClosed = errors.New("reward period is not closed")
)

type RewardPeriod struct {
	Start time.Time
	End   time.Time
}

func (p RewardPeriod) Validate() error {
	if p.Start.IsZero() || p.End.IsZero() || !p.End.After(p.Start) {
		return ErrInvalidRewardPeriod
	}
	return nil
}

func (p RewardPeriod) Key() string {
	return p.Start.UTC().Format(time.RFC3339) + "/" + p.End.UTC().Format(time.RFC3339)
}

func FixedRewardPeriod(at time.Time, width time.Duration) (RewardPeriod, error) {
	if width <= 0 || width%time.Second != 0 {
		return RewardPeriod{}, ErrInvalidRewardPeriod
	}
	at = at.UTC()
	seconds := int64(width / time.Second)
	startUnix := at.Unix() - at.Unix()%seconds
	start := time.Unix(startUnix, 0).UTC()
	return RewardPeriod{Start: start, End: start.Add(width)}, nil
}

type RewardPeriodInput struct {
	Session       domain.MiningSession
	AssetID       domain.AssetID
	Intervals     []domain.MiningInterval
	Hashrate      string
	RewardPerHash string
	Policy        domain.MiningRewardPolicy
}

type RewardPeriodInputProvider interface {
	Load(ctx context.Context, sessionID domain.MiningSessionID, period RewardPeriod) (RewardPeriodInput, error)
}

type RewardSettlement interface {
	SettleReward(ctx context.Context, reward domain.Reward) error
}

type RewardPeriodWorker struct {
	Sessions    MiningSessionStore
	Inputs      RewardPeriodInputProvider
	Settlement  RewardSettlement
	Clock       Clock
	PeriodWidth time.Duration
}

func (w RewardPeriodWorker) Process(ctx context.Context, sessionID domain.MiningSessionID, period RewardPeriod) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	if sessionID == "" {
		return domain.ErrInvalidRewardInput
	}
	if err := period.Validate(); err != nil {
		return err
	}
	now := w.Clock.Now().UTC()
	if !now.After(period.End) {
		return ErrRewardPeriodNotClosed
	}

	session, err := w.Sessions.GetByID(ctx, sessionID)
	if err != nil {
		return err
	}
	input, err := w.Inputs.Load(ctx, sessionID, period)
	if err != nil {
		return err
	}
	if input.Session.ID != session.ID || input.AssetID == "" {
		return domain.ErrInvalidRewardInput
	}

	reward, err := CalculateSessionPeriodReward(ctx, SessionRewardInput{
		Session:       input.Session,
		Intervals:     input.Intervals,
		Hashrate:      input.Hashrate,
		RewardPerHash: input.RewardPerHash,
		PolicyVersion: input.Policy.Version,
		PeriodKey:     period.Key(),
	}, input.Policy, period.End)
	if err != nil {
		return err
	}
	reward.AssetID = input.AssetID
	return w.Settlement.SettleReward(ctx, reward)
}
