# PostgreSQL Mining Runtime Persistence

Step 6 binds the mining runtime contracts to PostgreSQL transactions.

## Runtime transaction boundary

- Start inserts the session and its first interval in one transaction.
- Pause locks the session row with `FOR UPDATE`, then closes the open interval.
- Resume locks the session row, then inserts the next open interval.
- Complete/Cancel locks the session, closes any open interval, and updates session lifecycle state atomically.

The session row is the concurrency lock target. Redis is not the source of truth for lifecycle mutation.

## Interval invariants

PostgreSQL enforces:

- interval end cannot precede interval start;
- at most one open interval exists per session;
- intervals are indexed by session and start time.

The application/domain layer remains responsible for lifecycle semantics; the database enforces structural invariants.

## Reward settlement

`RewardSettlementStore` inserts the reward and its double-entry ledger effect in the same PostgreSQL transaction.

The reward idempotency key is unique. If a worker retries an already committed period, the reward insert conflicts and the transaction returns successfully without adding another ledger entry.

The ledger accounts are explicit configuration. No asset-specific account or coin is hardcoded.

## Scope

This step provides PostgreSQL persistence boundaries and SQL transaction behavior. It does not add a production PostgreSQL driver, worker scheduler, blockchain settlement, payment processing, or production account provisioning.
