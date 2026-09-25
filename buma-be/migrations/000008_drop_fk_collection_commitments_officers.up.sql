-- Migration: 000008_drop_fk_collection_commitments_officers.up.sql
-- Description: Drop foreign key constraints on officer_id and officer_pair_id in collection_commitments to allow storing officer names/IDs directly

ALTER TABLE collection_commitments DROP CONSTRAINT IF EXISTS collection_commitments_officer_id_fkey;
ALTER TABLE collection_commitments DROP CONSTRAINT IF EXISTS collection_commitments_officer_pair_id_fkey;
