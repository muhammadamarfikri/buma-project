-- Migration: 000007_add_fields_to_commitments.up.sql
-- Description: Add credit_limit, outstanding_balance, exposure_tier, product, and debtor_name columns to collection_commitments table

ALTER TABLE collection_commitments
    ADD COLUMN IF NOT EXISTS credit_limit NUMERIC(18, 2) DEFAULT 0,
    ADD COLUMN IF NOT EXISTS outstanding_balance NUMERIC(18, 2) DEFAULT 0,
    ADD COLUMN IF NOT EXISTS exposure_tier VARCHAR(50),
    ADD COLUMN IF NOT EXISTS product VARCHAR(100),
    ADD COLUMN IF NOT EXISTS debtor_name VARCHAR(255);

CREATE INDEX IF NOT EXISTS idx_commitments_exposure_tier ON collection_commitments(exposure_tier);
CREATE INDEX IF NOT EXISTS idx_commitments_product ON collection_commitments(product);
