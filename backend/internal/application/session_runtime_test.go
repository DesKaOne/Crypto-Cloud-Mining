package application

import (
	"context"
	"testing"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

type runtimeStore struct {
	started, paused, resumed, finished bool
	interval                           domain.MiningInterval
	endedAt                            time.Time
	status                             string
}

func (s *runtimeStore) Start(_ context.Context, _ domain.MiningSession, interval domain.MiningInterval) error {
	s.started = true
	s.interval = interval
	return nil
}

func (s *runtimeStore) Pause(_ context.Context, _ domain.MiningSessionID, endedAt time.Time) error {
	s.paused = true
	s.endedAt = endedAt
	return nil
}

func (s *runtimeStore) Resume(_ context.Context, _ domain.MiningSessionID, interval domain.MiningInterval) error {
	s.resumed = true
	s.interval = interval
	return nil
}

func (s *runtimeStore) Finish(_ context.Context, _ domain.MiningSessionID, status string, endedAt time.Time) error {
	s.finished = true
	s.status = status
	s.endedAt = endedAt
	return nil
}

func TestSessionRuntimeCommandUsesServerTimeForPause(t *testing.T) {
	now := time.Date(2026, 10, 4, 12, 0, 0, 0, time.FixedZone("WIB", 7*60*60))
	store := &runtimeStore{}
	err := (SessionRuntimeCommand{Sessions: store, Clock: fixedClock{now: now}}).Pause(context.Background(), "session-1")
	if err != nil {
		t.Fatal(err)
	}
	if !store.paused || !store.endedAt.Equal(now.UTC()) {
		t.Fatalf("unexpected pause timestamp: %v", store.endedAt)
	}
}

func TestSessionRuntimeCommandResumeCreatesOpenInterval(t *testing.T) {
	now := time.Date(2026, 10, 4, 12, 1, 0, 0, time.UTC)
	store := &runtimeStore{}
	err := (SessionRuntimeCommand{Sessions: store, Clock: fixedClock{now: now}}).Resume(context.Background(), "session-1", "interval-2")
	if err != nil {
		t.Fatal(err)
	}
	if !store.resumed || store.interval.EndedAt != nil {
		t.Fatal("expected open resumed interval")
	}
	if !store.interval.StartedAt.Equal(now) {
		t.Fatalf("unexpected start time: %v", store.interval.StartedAt)
	}
}

func TestSessionRuntimeCommandFinishRejectsMissingSession(t *testing.T) {
	store := &runtimeStore{}
	err := (SessionRuntimeCommand{Sessions: store, Clock: fixedClock{now: time.Now()}}).Complete(context.Background(), "")
	if err != ErrInvalidMiningTransition {
		t.Fatalf("expected invalid transition, got %v", err)
	}
}

func TestCalculateSessionPeriodRewardUsesOnlyActiveIntervals(t *testing.T) {
	start := time.Date(2026, 10, 4, 10, 0, 0, 0, time.UTC)
	pause := start.Add(10 * time.Minute)
	resume := start.Add(20 * time.Minute)
	end := start.Add(30 * time.Minute)
	reward, err := CalculateSessionPeriodReward(context.Background(), SessionRewardInput{
		Session: domain.MiningSession{ID: "session-1"},
		Intervals: []domain.MiningInterval{
			{StartedAt: start, EndedAt: &pause},
			{StartedAt: resume, EndedAt: &end},
		},
		Hashrate: "100", RewardPerHash: "0.000001", PolicyVersion: 1,
		PeriodKey: "2026-10-04T10:00:00Z/2026-10-04T10:30:00Z",
	}, domain.MiningRewardPolicy{Version: 1}, end)
	if err != nil {
		t.Fatal(err)
	}
	if reward.Quantity != "0.12" {
		t.Fatalf("reward = %q, want %q", reward.Quantity, "0.12")
	}
}

func TestCalculateSessionPeriodRewardRejectsPolicyMismatch(t *testing.T) {
	end := time.Unix(1, 0).UTC()
	start := time.Unix(0, 0).UTC()
	_, err := CalculateSessionPeriodReward(context.Background(), SessionRewardInput{
		Session:   domain.MiningSession{ID: "session-1"},
		Intervals: []domain.MiningInterval{{StartedAt: start, EndedAt: &end}},
		Hashrate:  "1", RewardPerHash: "1", PolicyVersion: 2, PeriodKey: "period-1",
	}, domain.MiningRewardPolicy{Version: 1}, end)
	if err == nil {
		t.Fatal("expected policy mismatch error")
	}
}
