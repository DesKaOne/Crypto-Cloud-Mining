# Phase 1 — Mining Engine Foundation

The mining engine is the authoritative calculation boundary for mining rewards.

## Design

The engine accepts versioned policy/configuration and authoritative inputs:

- hashrate
- active mining duration
- reward rate per hash-second
- policy version

The calculation uses exact rational arithmetic and returns a decimal string. It does not use floating-point arithmetic.

## Multi-asset boundary

The engine does not contain branches for BTC, LTC, or any other specific asset.

Asset-specific economics belong to a versioned policy/adapter that supplies the engine inputs.

## Authority

The server determines:

- mining start/end timestamps
- effective mining duration
- applicable plan/policy version
- reward rate
- resulting reward

The client may display estimates, but client-provided reward quantities are never authoritative.

## Precision

The current persistence baseline supports numeric(78,36). The engine emits at most 36 decimal places to remain compatible with that boundary.

## Phase 1 follow-ups

- connect policy selection to MiningPlan versions
- calculate active intervals across pause/resume transitions
- persist reward inputs for auditability
- implement reward posting with a real PostgreSQL transaction
- add asset-specific policy adapters
- add worker scheduling and idempotent period processing
- expose calculated/reconciled rewards through the API
