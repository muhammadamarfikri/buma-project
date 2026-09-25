-- Migration: 000007_add_fields_to_commitments.down.sql
-- Description: Drop added columns and indexes from collection_commitments table

DROP INDEX IF EXISTS idx_commitments_product;
DROP INDEX IF EXISTS idx_commitments_exposure_tier;

ALTER TABLE collection_commitments
    DROP COLUMN IF EXISTS debtor_name,
    DROP COLUMN IF EXISTS product,
    DROP COLUMN IF EXISTS exposure_tier,
    DROP COLUMN IF EXISTS outstanding_balance,
    DROP COLUMN IF EXISTS credit_limit;
