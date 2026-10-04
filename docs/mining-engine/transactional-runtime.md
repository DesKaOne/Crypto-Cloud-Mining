# Transactional Session Runtime

Phase 1 Step 4 makes mining lifecycle transitions atomic at the application persistence boundary.

## Atomic operations

The runtime store exposes four lifecycle operations:
- `Start`: persist an active session and its first open interval together;
- `Pause`: close the open interval and move the session to `paused` together;
- `Resume`: move the session to `active` and create the new open interval together;
- `Finish`: close the open interval and move the session to `completed` or `cancelled` together.

Each operation is one transaction. The application layer owns the contract; PostgreSQL owns the actual transaction and row locking.

## Concurrency rule

A runtime implementation must lock the target session row before validating its current status. The transition and interval mutation then happen under the same transaction.

Do not use Redis as the source of truth for lifecycle concurrency. Redis may later be used as an optimization, but correctness must survive a Redis outage.

## Failure behavior

If any session or interval mutation fails, the transaction rolls back both sides. The system must never expose a successful lifecycle command with only half of its persistence changes applied.

## Start identity

The application supplies the session and interval identities. The persistence layer may use UUID-backed storage, but the domain remains independent of the UUID implementation.

## Scope

This step does not implement reward posting, worker scheduling, blockchain settlement, payment processing, or production PostgreSQL driver wiring.