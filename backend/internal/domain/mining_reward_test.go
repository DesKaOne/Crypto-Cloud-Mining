package domain

import "testing"

func TestMiningRewardPolicyCalculate(t *testing.T) {
	policy := MiningRewardPolicy{Version: 1}

	got, err := policy.CalculateMiningReward(MiningRewardInput{
		Hashrate:         "100",
		DurationSeconds:  60,
		RewardPerHashSec: "0.000001",
	})
	if err != nil {
		t.Fatalf("CalculateMiningReward() error = %v", err)
	}
	if got != "0.006" {
		t.Fatalf("CalculateMiningReward() = %q, want %q", got, "0.006")
	}
}

func TestMiningRewardPolicyAvoidsFloatingPointLoss(t *testing.T) {
	policy := MiningRewardPolicy{Version: 1}

	got, err := policy.CalculateMiningReward(MiningRewardInput{
		Hashrate:         "0.1",
		DurationSeconds:  3,
		RewardPerHashSec: "0.1",
	})
	if err != nil {
		t.Fatalf("CalculateMiningReward() error = %v", err)
	}
	if got != "0.03" {
		t.Fatalf("CalculateMiningReward() = %q, want %q", got, "0.03")
	}
}

func TestMiningRewardPolicyRejectsInvalidInput(t *testing.T) {
	policy := MiningRewardPolicy{Version: 1}

	cases := []MiningRewardInput{
		{Hashrate: "-1", DurationSeconds: 1, RewardPerHashSec: "1"},
		{Hashrate: "1", DurationSeconds: -1, RewardPerHashSec: "1"},
		{Hashrate: "not-a-number", DurationSeconds: 1, RewardPerHashSec: "1"},
		{Hashrate: "1", DurationSeconds: 1, RewardPerHashSec: "-0.1"},
	}

	for _, input := range cases {
		if _, err := policy.CalculateMiningReward(input); err == nil {
			t.Fatalf("expected invalid input error for %+v", input)
		}
	}
}

func TestMiningRewardPolicyIsDeterministic(t *testing.T) {
	policy := MiningRewardPolicy{Version: 7}
	input := MiningRewardInput{
		Hashrate:         "123456789.123456789",
		DurationSeconds:  86400,
		RewardPerHashSec: "0.000000000000000123",
	}

	first, err := policy.CalculateMiningReward(input)
	if err != nil {
		t.Fatalf("first calculation error = %v", err)
	}
	second, err := policy.CalculateMiningReward(input)
	if err != nil {
		t.Fatalf("second calculation error = %v", err)
	}
	if first != second {
		t.Fatalf("calculation is not deterministic: %q != %q", first, second)
	}
}
