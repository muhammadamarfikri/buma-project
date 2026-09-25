-- Migration: 000003_add_officer_id_and_reason_to_commitments.down.sql
-- Description: Rollback additions of officer_id and reason columns from collection_commitments table

DROP INDEX IF EXISTS idx_commitments_officer_id;

ALTER TABLE collection_commitments 
    DROP COLUMN IF EXISTS officer_id,
    DROP COLUMN IF EXISTS reason;
