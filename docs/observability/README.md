# Observability Baseline

## Signals

Phase 0 defines three observability categories:

- Logs: structured application events.
- Metrics: request and worker operational counters.
- Health/readiness: process and dependency status.

## Request ID

Every API request should have a server-generated request ID when the client does not provide one.

The request ID is returned in the response header and included in structured logs.

A client-supplied request ID may be accepted only after validation and should never be treated as an authentication credential.

## Logging rules

Logs must not contain:

- access tokens
- refresh/session secrets
- private signing keys
- full license secrets
- passwords
- raw Telegram authentication material

Business identifiers may be logged only when useful for operations and privacy requirements permit it.

## Health endpoints

The health endpoint indicates that the process is alive.

The readiness endpoint indicates whether the process is ready to serve normal traffic. Dependency checks can be added when database/cache adapters are wired.

Health endpoints must not expose credentials, environment secrets, or internal stack traces.

## Metrics

Future production metrics should cover:

- HTTP request count/latency/errors
- authentication failures
- authorization denials
- mining session transitions
- reward calculation/posting failures
- ledger conflicts
- worker job success/failure/latency
- license validation and activation outcomes

No commercial or financial secret should be used as a metric label.
