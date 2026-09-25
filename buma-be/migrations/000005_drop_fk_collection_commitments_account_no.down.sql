-- Migration: 000005_drop_fk_collection_commitments_account_no.down.sql
-- Description: Re-add Foreign Key constraint on collection_commitments(account_no) pointing to debtor_accounts(account_no)

ALTER TABLE collection_commitments
ADD CONSTRAINT collection_commitments_account_no_fkey
FOREIGN KEY (account_no) REFERENCES debtor_accounts(account_no) ON DELETE CASCADE;
