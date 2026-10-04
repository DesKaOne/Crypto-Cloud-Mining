# API Contracts

The API contract is the compatibility boundary between backend, Flutter, Telegram Mini App, and future integrations.

## Rules

- OpenAPI is the source contract for HTTP APIs.
- API versioning starts at /api/v1.
- Response envelopes use data for successful resource responses.
- Errors use a stable error.code plus human-readable error.message.
- Financial quantities are strings containing exact decimal values.
- Mutating/retryable operations require Idempotency-Key.
- Server-authoritative fields such as timestamps, reward quantities, balances, and plan versions are not accepted from clients.
- Authentication is transport-level; authorization remains an application-layer decision.

The contract is intentionally small in Phase 0. Endpoint coverage expands only when the corresponding use case is implemented.
