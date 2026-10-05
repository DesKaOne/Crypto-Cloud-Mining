package domain

import (
	"errors"
	"time"
)

var (
	ErrInvalidMiningInterval = errors.New("invalid mining interval")
	ErrOpenMiningInterval    = errors.New("mining interval is still open")
)

type MiningInterval struct {
	ID        MiningIntervalID
	SessionID MiningSessionID
	StartedAt time.Time
	EndedAt   *time.Time
}

func (i MiningInterval) DurationUntil(now time.Time) (time.Duration, error) {
	start := i.StartedAt.UTC()
	end := now.UTC()
	if i.EndedAt != nil {
		end = i.EndedAt.UTC()
	}
	if end.Before(start) {
		return 0, ErrInvalidMiningInterval
	}
	return end.Sub(start), nil
}

func ActiveDuration(intervals []MiningInterval, now time.Time) (time.Duration, error) {
	var total time.Duration
	for _, interval := range intervals {
		duration, err := interval.DurationUntil(now)
		if err != nil {
			return 0, err
		}
		if interval.EndedAt == nil {
			return 0, ErrOpenMiningInterval
		}
		total += duration
	}
	return total, nil
}

func ActiveDurationSeconds(intervals []MiningInterval, now time.Time) (int64, error) {
	duration, err := ActiveDuration(intervals, now)
	if err != nil {
		return 0, err
	}
	return int64(duration / time.Second), nil
}
