DROP INDEX IF EXISTS rewards_asset_period_idx;

ALTER TABLE rewards
    DROP CONSTRAINT IF EXISTS rewards_session_period_unique;
