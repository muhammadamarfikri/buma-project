-- Migration: 000010_add_cif_to_commitments_and_debtors.down.sql
-- Description: Drop cif column from collection_commitments and debtor_accounts tables

DROP INDEX IF EXISTS idx_commitments_cif;
DROP INDEX IF EXISTS idx_debtor_accounts_cif;

ALTER TABLE collection_commitments DROP COLUMN IF EXISTS cif;
ALTER TABLE debtor_accounts DROP COLUMN IF EXISTS cif;
