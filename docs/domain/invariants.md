# Domain Invariants

These invariants are the first guardrails for the core domain.

## Identity

1. Every User has a stable internal ID.
2. External identities (Telegram, email, OAuth provider, etc.) are identifiers, not primary domain identity.

## Asset

1. Asset has a stable internal ID and unique code.
2. Asset code is case-normalized.
3. Disabled assets cannot start new mining sessions or new deposits.
4. Asset-specific behavior belongs behind an adapter/policy boundary.

## MiningPlan

1. A plan belongs to exactly one Asset.
2. A plan has an explicit lifecycle status.
3. Plan configuration is versionable; historical sessions must retain the configuration/version used for their calculation.
4. A plan cannot be started when inactive.

## MiningSession

1. A session belongs to one User and one MiningPlan.
2. Session lifecycle transitions are explicit.
3. A session cannot accrue rewards outside its active mining interval.
4. Reward calculation must be deterministic from persisted inputs and the applicable policy/configuration version.
5. Client clocks are never authoritative for reward accrual.

## Reward

1. Reward records are attributable to a session and asset.
2. Reward creation must be idempotent for the same calculation period/input identity.
3. A reward record is not silently mutated after it has been posted to the financial ledger.

## Wallet and Ledger

1. Wallet balance is derived from or reconciled against ledger entries; it is not an independent source of truth.
2. Ledger entries are append-only.
3. Every financial movement has explicit debit/credit semantics.
4. Ledger writes must be idempotent.
5. Asset balances never mix across assets.

## Transaction

1. Deposits and withdrawals have explicit lifecycle states.
2. External transaction IDs/hashes are unique within the relevant asset/network scope.
3. A failed or rejected transaction cannot silently become completed.
