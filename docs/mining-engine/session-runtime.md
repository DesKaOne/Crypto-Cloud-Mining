# Mining Session Runtime

## Active intervals

Mining time is represented by contiguous MiningInterval records.

Example:

    active  10:00 -> 10:10
    paused  10:10 -> 10:20
    active  10:20 -> 10:30

Effective mining time is 20 minutes, not 30 minutes.

This avoids deriving reward time from session wall-clock duration alone.

## Reward periods

A reward period identifies a deterministic calculation window with:

- session ID
- period key
- applicable mining plan/policy version
- authoritative mining intervals
- server-side hashrate/rate inputs

The reward idempotency key is sessionID + ":" + periodKey.

The persistence layer must enforce uniqueness for this key.

## Pause/resume

Pause closes the current active interval.

Resume starts a new active interval.

Terminal completion/cancellation closes the current interval before final reward processing.

## Policy version

A session captures its MiningPlan version at start. Reward calculation must use the corresponding policy version; mismatches are rejected.

## Phase 1 follow-ups

- persist interval records
- implement transactional pause/resume
- define fixed period boundaries for the worker
- prevent overlapping intervals
- finalize reward-period settlement transaction
