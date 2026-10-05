# Mining Runtime Persistence

Phase 1 Step 3 persists mining session intervals without moving reward authority into the database.

## Stored runtime state

`mining_session_intervals` records each active interval of a mining session:
- `session_id` binds the interval to one mining session;
- `started_at` is the server-authoritative start;
- `ended_at` is set when the session pauses or reaches a terminal state;
- the interval `id` is an opaque UUID persistence identity.

The database enforces:
- `ended_at` cannot precede `started_at`;
- at most one open interval exists for a session;
- intervals are ordered by session and start time.

## Repository boundary

`domain.MiningIntervalRepository` exposes only the operations required by the application:
- open an interval;
- close the currently open interval;
- list intervals for reward calculation.

The PostgreSQL adapter uses `database/sql` and keeps SQL outside the domain package. No PostgreSQL driver is selected in this step; driver wiring remains an infrastructure/runtime concern.

## Runtime transaction rule

Pause, resume, and terminal transitions must not be implemented as independent writes in production. The session status change and corresponding interval mutation must execute in one database transaction.

Expected lifecycle:
1. Start: create the session and its first open interval in one transaction.
2. Pause: close the open interval and set session status to `paused` in one transaction.
3. Resume: set session status to `active` and create a new open interval in one transaction.
4. Complete/cancel: close the open interval and set the terminal session state in one transaction.
5. Reward calculation reads persisted intervals and uses the existing deterministic policy and period idempotency key.

## Non-goals

This step does not add:
- a PostgreSQL driver;
- worker scheduling;
- blockchain settlement;
- wallet balance mutation;
- production payment processing.

Those concerns remain separate from runtime interval persistence.