# Core Domain Model

Phase 0 defines the domain vocabulary and invariants before persistence or production mining logic.

## Core aggregates

- User: platform identity and account state.
- Asset: supported crypto asset configuration.
- MiningPlan: product configuration defining capacity, pricing, duration, and reward rules.
- MiningSession: a user's active or historical participation in a mining plan.
- Reward: immutable record of reward accrual attributable to a mining session.
- Wallet: asset-specific user balance container.
- LedgerEntry: append-only financial/domain movement between accounts.
- Transaction: externally observable deposit, withdrawal, adjustment, or settlement lifecycle.

## Supporting concepts

- Hashrate: normalized mining capacity with unit and source.
- Referral: relationship and attribution between users.
- AssetAdapter: boundary for asset-specific chain/wallet operations.
- RewardPolicy: deterministic policy used to calculate rewards.

## Domain rule

Crypto assets are data/configuration plus optional adapters. Business flows must not branch on hardcoded coin names.

## Important separation

Mining reward calculation and financial ledger accounting are separate concerns. A calculated reward becomes financially meaningful only through an explicit ledger operation.
