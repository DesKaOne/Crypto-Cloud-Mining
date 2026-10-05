package domain

import (
	"errors"
	"math/big"
	"strings"
)

var (
	ErrInvalidRewardInput = errors.New("invalid mining reward input")
	ErrNegativeDuration   = errors.New("mining duration cannot be negative")
)

type MiningRewardInput struct {
	Hashrate         string
	DurationSeconds  int64
	RewardPerHashSec string
}

type MiningRewardPolicy struct {
	Version int
}

// CalculateMiningReward deterministically calculates reward = hashrate * duration * rate.
// All arithmetic uses rational numbers; no floating-point arithmetic is used.
func (p MiningRewardPolicy) CalculateMiningReward(input MiningRewardInput) (string, error) {
	if p.Version < 1 || input.Hashrate == "" || input.RewardPerHashSec == "" {
		return "", ErrInvalidRewardInput
	}
	if input.DurationSeconds < 0 {
		return "", ErrNegativeDuration
	}

	hashrate, ok := new(big.Rat).SetString(strings.TrimSpace(input.Hashrate))
	if !ok || hashrate.Sign() < 0 {
		return "", ErrInvalidRewardInput
	}
	rate, ok := new(big.Rat).SetString(strings.TrimSpace(input.RewardPerHashSec))
	if !ok || rate.Sign() < 0 {
		return "", ErrInvalidRewardInput
	}

	duration := new(big.Rat).SetInt64(input.DurationSeconds)
	reward := new(big.Rat).Mul(hashrate, duration)
	reward.Mul(reward, rate)

	return decimalFromRat(reward), nil
}

func decimalFromRat(value *big.Rat) string {
	if value.Sign() == 0 {
		return "0"
	}

	s := value.FloatString(36)
	s = strings.TrimRight(strings.TrimRight(s, "0"), ".")
	return s
}
