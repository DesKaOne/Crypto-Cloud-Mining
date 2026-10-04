# Architecture Decisions

## ADR-0001 — Modular monorepo

Status: Accepted

The project uses a single repository containing clients, backend services, shared contracts, infrastructure, and documentation.

Reason: the platform is one product with multiple delivery surfaces. A monorepo keeps API contracts, architecture, and release changes coordinated.

## ADR-0002 — Server-authoritative domain state

Status: Accepted

Balances, mining sessions, reward calculations, plan state, transactions, and security-sensitive state are authoritative on the backend.

Reason: clients are untrusted environments and must not be able to determine financial state.

## ADR-0003 — Multi-asset domain model

Status: Accepted

The system models supported crypto assets through configuration/domain entities rather than creating one implementation branch per coin.

Reason: adding an asset should primarily require configuration and an asset adapter where chain-specific behavior is genuinely different.

## ADR-0004 — Flutter as primary client technology

Status: Accepted

Flutter is the primary UI technology for mobile and web client surfaces where practical.

Reason: shared presentation and domain-client code reduce duplicated client implementations.

Telegram Mini App remains a separate delivery surface when Telegram-specific integration requires web/JavaScript behavior.

## ADR-0005 — Go backend baseline

Status: Accepted

Go is the initial backend implementation baseline.

Reason: predictable deployment, strong concurrency primitives, low runtime overhead, and a clear path to worker-oriented services.

The architecture keeps backend domain boundaries independent from the transport layer so a future Dart service is not structurally blocked.

## ADR-0006 — PostgreSQL + Redis baseline

Status: Accepted

PostgreSQL is the durable system of record. Redis is an auxiliary store for caching, rate limiting, ephemeral state, and queue-like workloads where appropriate.

Neither replaces the other.
