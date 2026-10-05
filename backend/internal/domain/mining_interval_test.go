package domain

import (
	"testing"
	"time"
)

func TestActiveDurationExcludesPausedTime(t *testing.T) {
	start := time.Date(2026, 10, 4, 10, 0, 0, 0, time.UTC)
	pause := start.Add(10 * time.Minute)
	resume := start.Add(20 * time.Minute)
	end := start.Add(30 * time.Minute)

	seconds, err := ActiveDurationSeconds([]MiningInterval{
		{StartedAt: start, EndedAt: &pause},
		{StartedAt: resume, EndedAt: &end},
	}, end)
	if err != nil {
		t.Fatal(err)
	}
	if seconds != 20*60 {
		t.Fatalf("active duration = %d seconds, want %d", seconds, 20*60)
	}
}

func TestActiveDurationRejectsOpenInterval(t *testing.T) {
	start := time.Date(2026, 10, 4, 10, 0, 0, 0, time.UTC)
	if _, err := ActiveDurationSeconds([]MiningInterval{{StartedAt: start}}, start.Add(time.Minute)); err != ErrOpenMiningInterval {
		t.Fatalf("expected ErrOpenMiningInterval, got %v", err)
	}
}

func TestActiveDurationRejectsNegativeInterval(t *testing.T) {
	start := time.Date(2026, 10, 4, 10, 0, 0, 0, time.UTC)
	end := start.Add(-time.Second)
	if _, err := ActiveDurationSeconds([]MiningInterval{{StartedAt: start, EndedAt: &end}}, start); err != ErrInvalidMiningInterval {
		t.Fatalf("expected ErrInvalidMiningInterval, got %v", err)
	}
}
