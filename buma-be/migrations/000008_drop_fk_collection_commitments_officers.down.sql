-- Migration: 000008_drop_fk_collection_commitments_officers.down.sql
-- Description: Re-add foreign key constraints on officer_id and officer_pair_id in collection_commitments pointing to officers(officer_id)

ALTER TABLE collection_commitments ADD CONSTRAINT collection_commitments_officer_id_fkey FOREIGN KEY (officer_id) REFERENCES officers(officer_id) ON DELETE SET NULL;
ALTER TABLE collection_commitments ADD CONSTRAINT collection_commitments_officer_pair_id_fkey FOREIGN KEY (officer_pair_id) REFERENCES officers(officer_id) ON DELETE SET NULL;
