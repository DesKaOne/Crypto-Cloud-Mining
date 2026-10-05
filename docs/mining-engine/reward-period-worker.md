# Reward Period Worker

## Fixed period boundaries

Reward periods are deterministic windows. The worker must process a period only after its end timestamp has passed.
The default period width is one hour. The period key is the UTC start and end timestamps in RFC3339 form.
The key is derived from the boundary, not from worker execution time.

## Processing flow

    closed period
        |
        v
    load session
        |
        v
    load authoritative asset/policy/interval inputs
        |
        v
    calculate deterministic reward
        |
        v
    SettleReward
        |
        +--> reward record
        |
        +--> ledger entry

The worker never accepts a client-provided reward quantity.

## Idempotency

The reward idempotency identity is sessionID + periodKey. Persistence must enforce uniqueness for the identity.
A retry of the same period therefore resolves to the existing financial operation instead of creating a second reward.
SettleReward is the application boundary for the financial operation. Its infrastructure implementation must make reward creation and its ledger effect atomic. This avoids the failure gap where a reward record could exist without its ledger posting.

## Exact quantities

Reward calculation continues to use rational arithmetic and decimal strings. PostgreSQL remains the durable source of truth with numeric(78,36).

## Asset independence

The worker receives the resolved AssetID from its input provider. No BTC, LTC, or other asset is hardcoded into the worker.

## Scope

This step establishes period calculation and settlement contracts. It does not implement blockchain settlement, payment processing, production worker scheduling, or final reward economics.