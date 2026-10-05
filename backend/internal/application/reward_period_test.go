package application

import (
    "context"
    "testing"
    "time"

    "github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
 )

type rewardPeriodSessionStore struct { session domain.MiningSession }
func (s rewardPeriodSessionStore) Create(context.Context, domain.MiningSession) error { return nil }
func (s rewardPeriodSessionStore) GetByID(context.Context, domain.MiningSessionID) (*domain.MiningSession, error) { return &s.session, nil }
func (s rewardPeriodSessionStore) Update(context.Context, domain.MiningSession) error { return nil }

type rewardPeriodInputs struct { input RewardPeriodInput }
func (p rewardPeriodInputs) Load(context.Context, domain.MiningSessionID, RewardPeriod) (RewardPeriodInput, error) { return p.input, nil }

type rewardSettlement struct { calls int; reward domain.Reward }
func (s *rewardSettlement) SettleReward(_ context.Context, reward domain.Reward) error { s.calls++; s.reward = reward; return nil }

func TestFixedRewardPeriodUsesDeterministicBoundaries(t *testing.T) {
    at := time.Date(2026, 10, 5, 8, 37, 42, 0, time.FixedZone("WIB", 7*60*60))
    period, err := FixedRewardPeriod(at, time.Hour); if err != nil { t.Fatal(err) }
    wantStart := time.Date(2026, 10, 5, 8, 0, 0, 0, time.UTC)
    if !period.Start.Equal(wantStart) || !period.End.Equal(wantStart.Add(time.Hour)) { t.Fatalf("period = %#v", period) }
    if period.Key() != "2026-10-05T08:00:00Z/2026-10-05T09:00:00Z" { t.Fatalf("unexpected period key: %q", period.Key()) }
}

func TestRewardPeriodWorkerRequiresClosedPeriod(t *testing.T) {
    now := time.Date(2026, 10, 5, 9, 0, 0, 0, time.UTC)
    period := RewardPeriod{Start: now.Add(-time.Hour), End: now}
    worker := RewardPeriodWorker{Clock: fixedClock{now: now}, PeriodWidth: time.Hour}
    err := worker.Process(context.Background(), "session-1", period)
    if err != ErrRewardPeriodNotClosed { t.Fatalf("expected closed-period error, got %v", err) }
}

func TestRewardPeriodWorkerSettlesIdempotentRewardIdentity(t *testing.T) {
    start := time.Date(2026, 10, 5, 7, 0, 0, 0, time.UTC); end := start.Add(time.Hour)
    session := domain.MiningSession{ID: "session-1"}; settlement := &rewardSettlement{}
    worker := RewardPeriodWorker{
        Sessions: rewardPeriodSessionStore{session: session},
        Inputs: rewardPeriodInputs{input: RewardPeriodInput{
            Session: session, AssetID: "asset-1",
            Intervals: []domain.MiningInterval{{StartedAt: start, EndedAt: &end}},
            Hashrate: "100", RewardPerHash: "0.000001", Policy: domain.MiningRewardPolicy{Version: 1},
        }},
        Settlement: settlement, Clock: fixedClock{now: end.Add(time.Minute)}, PeriodWidth: time.Hour,
    }
    if err := worker.Process(context.Background(), session.ID, RewardPeriod{Start: start, End: end}); err != nil { t.Fatal(err) }
    if settlement.calls != 1 { t.Fatalf("settlement calls = %d, want 1", settlement.calls) }
    if settlement.reward.AssetID != "asset-1" || settlement.reward.Quantity != "0.36" { t.Fatalf("unexpected reward: %#v", settlement.reward) }
    if settlement.reward.IdempotencyKey != "session-1:2026-10-05T07:00:00Z/2026-10-05T08:00:00Z" { t.Fatalf("unexpected idempotency key: %q", settlement.reward.IdempotencyKey) }
    if settlement.reward.ID != "" { t.Fatalf("reward ID should be persistence-generated, got %q", settlement.reward.ID) }
}