-- Migration: 000003_add_officer_id_and_reason_to_commitments.up.sql
-- Description: Alter table collection_commitments to add officer_id and reason columns with index and foreign key constraint

ALTER TABLE collection_commitments 
    ADD COLUMN IF NOT EXISTS reason TEXT,
    ADD COLUMN IF NOT EXISTS officer_id VARCHAR(50) REFERENCES officers(officer_id) ON DELETE SET NULL;

-- B-Tree Index for filtering commitments by officer_id
CREATE INDEX IF NOT EXISTS idx_commitments_officer_id ON collection_commitments(officer_id);
