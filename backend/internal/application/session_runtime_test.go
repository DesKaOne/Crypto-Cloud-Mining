package application

import (
	"context"
	"testing"
	"time"

	"github.com/DesKaOne/Crypto-Cloud-Mining/backend/internal/domain"
)

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
		Hashrate:      "100",
		RewardPerHash: "0.000001",
		PolicyVersion: 1,
		PeriodKey:     "2026-10-04T10:00:00Z/2026-10-04T10:30:00Z",
	}, domain.MiningRewardPolicy{Version: 1}, end)
	if err != nil {
		t.Fatal(err)
	}
	if reward.Quantity != "0.12" {
		t.Fatalf("reward = %q, want %q", reward.Quantity, "0.12")
	}
	if reward.IdempotencyKey != "session-1:2026-10-04T10:00:00Z/2026-10-04T10:30:00Z" {
		t.Fatalf("unexpected idempotency key: %q", reward.IdempotencyKey)
	}
}

func TestCalculateSessionPeriodRewardRejectsPolicyMismatch(t *testing.T) {
	end := time.Unix(1, 0).UTC()
	start := time.Unix(0, 0).UTC()
	_, err := CalculateSessionPeriodReward(context.Background(), SessionRewardInput{
		Session:       domain.MiningSession{ID: "session-1"},
		Intervals:     []domain.MiningInterval{{StartedAt: start, EndedAt: &end}},
		Hashrate:      "1",
		RewardPerHash: "1",
		PolicyVersion: 2,
		PeriodKey:     "period-1",
	}, domain.MiningRewardPolicy{Version: 1}, end)
	if err == nil {
		t.Fatal("expected policy mismatch error")
	}
}
