-- Migration: 000004_add_officer_pair_id_to_commitments.down.sql
-- Description: Rollback officer_pair_id column from collection_commitments table

DROP INDEX IF EXISTS idx_commitments_officer_pair_id;

ALTER TABLE collection_commitments 
    DROP COLUMN IF EXISTS officer_pair_id;
