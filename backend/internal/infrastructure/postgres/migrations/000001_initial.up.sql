CREATE EXTENSION IF NOT EXISTS pgcrypto;

CREATE TABLE users (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    status text NOT NULL DEFAULT 'active',
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now()
);

CREATE TABLE assets (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    code text NOT NULL,
    name text NOT NULL,
    decimals smallint NOT NULL CHECK (decimals >= 0 AND decimals <= 36),
    enabled boolean NOT NULL DEFAULT true,
    created_at timestamptz NOT NULL DEFAULT now(),
    updated_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT assets_code_unique UNIQUE (code)
);

CREATE TABLE mining_plans (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id uuid NOT NULL REFERENCES assets(id),
    version integer NOT NULL CHECK (version > 0),
    status text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    effective_at timestamptz NOT NULL,
    CONSTRAINT mining_plans_version_unique UNIQUE (id, version)
);

CREATE TABLE mining_sessions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL REFERENCES users(id),
    mining_plan_id uuid NOT NULL,
    plan_version integer NOT NULL,
    status text NOT NULL,
    started_at timestamptz,
    ended_at timestamptz,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT mining_sessions_plan_fk
        FOREIGN KEY (mining_plan_id, plan_version)
        REFERENCES mining_plans(id, version)
);

CREATE TABLE rewards (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    mining_session_id uuid NOT NULL REFERENCES mining_sessions(id),
    asset_id uuid NOT NULL REFERENCES assets(id),
    quantity numeric(78, 36) NOT NULL CHECK (quantity >= 0),
    period_key text NOT NULL,
    idempotency_key text NOT NULL,
    policy_version integer NOT NULL CHECK (policy_version > 0),
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT rewards_idempotency_unique UNIQUE (idempotency_key)
);

CREATE TABLE ledger_entries (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    asset_id uuid NOT NULL REFERENCES assets(id),
    reference_id text NOT NULL,
    debit_account text NOT NULL,
    credit_account text NOT NULL,
    quantity numeric(78, 36) NOT NULL CHECK (quantity > 0),
    idempotency_key text NOT NULL,
    created_at timestamptz NOT NULL DEFAULT now(),
    CONSTRAINT ledger_idempotency_unique UNIQUE (idempotency_key),
    CONSTRAINT ledger_accounts_distinct CHECK (debit_account <> credit_account)
);

CREATE INDEX mining_sessions_user_status_idx
    ON mining_sessions (user_id, status);

CREATE INDEX rewards_session_period_idx
    ON rewards (mining_session_id, period_key);

CREATE INDEX ledger_asset_reference_idx
    ON ledger_entries (asset_id, reference_id);
