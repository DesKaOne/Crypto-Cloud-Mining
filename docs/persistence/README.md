# Persistence Foundation

Phase 0 persistence establishes the database boundary without implementing production use cases.

## Database

PostgreSQL is the durable source of truth.

Core rules:
- opaque domain identifiers, persisted as UUID;
- timestamptz for persisted timestamps;
- exact numeric for crypto quantities;
- migrations are ordered and versioned;
- repository interfaces live at the domain boundary;
- SQL belongs in infrastructure, not domain entities.

## Financial integrity

Ledger records are append-only by application contract. Idempotency keys prevent duplicate financial operations.

A wallet balance is a projection/reconciliation result, not an authorization to mutate financial truth directly.

## Redis

Redis remains auxiliary. It must not become the source of truth for balances, rewards, deposits, or withdrawals.

## Migration policy

Every schema change adds a new numbered migration. Existing migrations are immutable after being applied to a shared environment.
