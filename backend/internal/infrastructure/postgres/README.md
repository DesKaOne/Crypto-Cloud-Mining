# PostgreSQL Persistence

PostgreSQL is the durable system of record.

Rules:
- schema changes use ordered migrations;
- financial tables are append-oriented;
- quantities are stored as exact numeric values, never floating-point;
- all timestamps are stored as timestamptz;
- asset scope is explicit on financial records;
- migrations must be reversible where practical;
- application code must not bypass repository boundaries for domain persistence.
