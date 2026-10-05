# API Contract Baseline

Phase 0 establishes a versioned HTTP contract before transport implementation.

## Boundary

Clients:
- Flutter mobile/web
- Telegram Mini App

consume the same backend API and application use cases.

## Versioning

Current contract namespace: /api/v1

Breaking changes require a new API version. Additive compatible changes may remain within the current version.

## Error model

Every API error uses:
- stable machine-readable error.code
- human-readable error.message
- optional error.requestId

Clients should branch on error codes, not message text.

## Idempotency

State-changing requests that may be retried require Idempotency-Key.

The backend owns idempotency persistence. A client-generated key does not make an operation successful; it only identifies an intended retry-safe operation.

## Quantity representation

Crypto quantities are transmitted as strings, never JSON floating-point numbers.

Example: 0.00001234

## Authentication

The API exposes a bearer authentication boundary in Phase 0. The concrete identity providers and token issuance flow are Step 8 concerns.
