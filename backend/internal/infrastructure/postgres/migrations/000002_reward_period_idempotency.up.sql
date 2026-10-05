ALTER TABLE rewards
    ADD CONSTRAINT rewards_session_period_unique UNIQUE (mining_session_id, period_key);

CREATE INDEX rewards_asset_period_idx
    ON rewards (asset_id, period_key);
