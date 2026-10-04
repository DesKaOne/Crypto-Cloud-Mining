CREATE TABLE mining_session_intervals (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    session_id uuid NOT NULL REFERENCES mining_sessions(id) ON DELETE CASCADE,
    started_at timestamptz NOT NULL,
    ended_at timestamptz,
    CONSTRAINT mining_session_intervals_time_check
        CHECK (ended_at IS NULL OR ended_at >= started_at)
);

CREATE UNIQUE INDEX mining_session_intervals_one_open_idx
    ON mining_session_intervals (session_id)
    WHERE ended_at IS NULL;

CREATE INDEX mining_session_intervals_session_started_idx
    ON mining_session_intervals (session_id, started_at);