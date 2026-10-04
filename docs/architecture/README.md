# Crypto Cloud Mining — Architecture Baseline

## Purpose

Dokumen ini menetapkan baseline arsitektur Phase 0 untuk platform Crypto Cloud Mining.

## Architectural principles

1. Multi-asset first: crypto adalah konfigurasi/domain data, bukan hardcoded feature.
2. Backend owns truth: balance, mining state, rewards, plans, transactions, and authentication are server-authoritative.
3. Client is presentation: Flutter clients consume APIs and realtime events; business-critical calculations must not live only in clients.
4. Modular domain boundaries: mining, wallet, rewards, plans, users, payments, and asset configuration remain independently evolvable.
5. Infrastructure is reproducible: local development and deployment configuration must be versioned.
6. Security by default: secrets never live in source control; privileged operations are server-side.
7. Phase-driven delivery: foundation first, features only after the relevant architecture baseline is stable.

## Target topology

```text
                         +----------------------+
                         |   Flutter Client     |
                         | Mobile + Web         |
                         +----------+-----------+
                                    |
                                    | HTTPS / WS
                                    v
+------------------+      +----------------------+
| Telegram MiniApp | ---> |   API / App Server   |
+------------------+      |       Go             |
                          +----------+-----------+
                                     |
                    +----------------+----------------+
                    |                |                |
                    v                v                v
              PostgreSQL          Redis        Background Workers
                    |                                 |
                    +----------------+----------------+
                                     |
                                     v
                              Mining / Reward Core
```

## Repository boundaries

- `apps/`: user-facing applications.
- `backend/`: server-side application and workers.
- `packages/`: shared contracts and non-runtime shared assets.
- `infrastructure/`: local/deployment infrastructure.
- `docs/`: architecture and project documentation.
- `scripts/`: developer and CI helper scripts.

## Phase 0 scope

Phase 0 establishes repository boundaries, architecture decisions, configuration conventions, development infrastructure, and CI foundations.

It does **not** implement real mining, financial settlement, withdrawal processing, or production payment flows.
