# Security Baseline

Phase 0 security establishes boundaries and threat-model decisions rather than a complete production security system.

## Trust boundaries

1. Client applications are untrusted.
2. API transport authenticates requests.
3. Application use cases authorize actions.
4. PostgreSQL is the durable source of truth.
5. Redis is auxiliary and must not become the sole source of financial truth.
6. Vendor licensing infrastructure is outside the customer application trust boundary.

## Required controls

- server-authoritative business state
- input validation at transport/application boundaries
- idempotency for retryable mutations
- append-oriented financial ledger
- secrets supplied by environment/deployment secret management
- no private signing keys in customer source code
- request correlation identifiers
- structured server-side logs without credentials or sensitive token material
- explicit authentication and authorization failures
- rate-limit boundary at transport/API infrastructure
- dependency and artifact verification in release processes

## License-specific threats

The commercial license model must consider:

- license tampering
- signature forgery
- replay
- activation abuse
- revocation
- expiration
- clock manipulation
- installation binding
- offline activation lifetime
- private signing-key compromise

Phase 0 records these threats and boundaries. Concrete cryptographic implementation and key operations are later work.

## Source-code reality

Customers who receive source code can technically modify it. The security goal is therefore authenticity and traceability of official releases, license payloads, entitlements, and activations rather than impossible source-code control.
