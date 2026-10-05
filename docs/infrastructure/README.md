# Infrastructure / Development Environment

Phase 0 provides a reproducible local foundation.

## Services

- PostgreSQL: durable application data.
- Redis: auxiliary cache/coordination/realtime support.
- API: HTTP API process.
- Worker: background processing process.

## Local startup

From the repository root:

    make infra-up

Then verify:

    curl http://localhost:8080/health

Stop the environment:

    make infra-down

## Configuration

.env.example documents local configuration names. Development credentials in Compose are intentionally non-production values.

Production secrets must come from the deployment secret mechanism and must never be committed.

## Persistence

Database schema changes are applied through ordered migrations. Migration execution is intentionally separate from application startup in Phase 0 so deployment orchestration can control it explicitly.

## Network boundary

Only the API is intended to be consumed by client applications. PostgreSQL and Redis ports are exposed for local development convenience and should remain private in deployed environments.
