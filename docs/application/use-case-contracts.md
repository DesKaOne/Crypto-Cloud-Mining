# Use-case Contracts

## StartMiningSession

Input:
- user ID
- mining plan ID

Server resolves:
- asset
- current plan version
- eligibility
- session start time

Output:
- created mining session ID
- effective start timestamp

The client cannot select the authoritative start timestamp.

## CalculateReward

Input:
- mining session ID
- calculation period/reference

Server resolves:
- session
- plan/version
- asset
- effective hashrate
- reward policy/version
- authoritative active interval

Output:
- deterministic reward quantity
- idempotency key
- policy version

The client never submits the reward quantity.

## PostReward

Input:
- calculated reward identity
- target financial account

Rules:
- reward identity is idempotent;
- ledger posting is atomic;
- a previously posted reward is never silently duplicated.

## GetWalletBalance

Input:
- user ID
- asset ID

Output:
- balance projection/reconciled balance

The balance is derived from ledger state.

## RequestWithdrawal

Input:
- user ID
- asset/network
- quantity
- destination

Server validates:
- asset/network enabled
- available balance
- withdrawal policy
- idempotency key

No blockchain broadcast occurs inside the request validation transaction itself.
