-- Migration: 000004_add_officer_pair_id_to_commitments.up.sql
-- Description: Add officer_pair_id column to collection_commitments table with foreign key constraint and index

ALTER TABLE collection_commitments 
    ADD COLUMN IF NOT EXISTS officer_pair_id VARCHAR(50) REFERENCES officers(officer_id) ON DELETE SET NULL;

CREATE INDEX IF NOT EXISTS idx_commitments_officer_pair_id ON collection_commitments(officer_pair_id);
