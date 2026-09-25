-- Migration: 000005_drop_fk_collection_commitments_account_no.up.sql
-- Description: Remove Foreign Key constraint on collection_commitments(account_no) pointing to debtor_accounts(account_no)

ALTER TABLE collection_commitments
DROP CONSTRAINT IF EXISTS collection_commitments_account_no_fkey;
