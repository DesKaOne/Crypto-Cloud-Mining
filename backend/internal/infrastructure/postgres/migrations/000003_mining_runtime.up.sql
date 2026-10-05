CREATE TABLE mining_intervals (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    mining_session_id uuid NOT NULL REFERENCES mining_sessions(id),
    started_at timestamptz NOT NULL,
    ended_at timestamptz,
    CONSTRAINT mining_intervals_time_order CHECK (ended_at IS NULL OR ended_at >= started_at)
);

CREATE UNIQUE INDEX mining_intervals_one_open_per_session_idx ON mining_intervals (mining_session_id) WHERE ended_at IS NULL;

CREATE INDEX mining_intervals_session_started_idx ON mining_intervals (mining_session_id, started_at);
