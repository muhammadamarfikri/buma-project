-- Migration: 000010_add_cif_to_commitments_and_debtors.up.sql
-- Description: Add cif column to collection_commitments and debtor_accounts tables and populate existing seed records

ALTER TABLE collection_commitments
    ADD COLUMN IF NOT EXISTS cif VARCHAR(50);

ALTER TABLE debtor_accounts
    ADD COLUMN IF NOT EXISTS cif VARCHAR(50);

CREATE INDEX IF NOT EXISTS idx_commitments_cif ON collection_commitments(cif);
CREATE INDEX IF NOT EXISTS idx_debtor_accounts_cif ON debtor_accounts(cif);

-- Populate CIF for existing seed debtor accounts if null
UPDATE debtor_accounts SET cif = '11002022001' WHERE account_no = '1029384756' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022002' WHERE account_no = '2093847561' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022003' WHERE account_no = '3084756192' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022004' WHERE account_no = '4075619283' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022005' WHERE account_no = '5061928374' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022006' WHERE account_no = '6019283745' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022007' WHERE account_no = '7092837465' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022008' WHERE account_no = '8028374651' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022009' WHERE account_no = '9037465182' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022010' WHERE account_no = '1147561928' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022011' WHERE account_no = '2256192837' AND (cif IS NULL OR cif = '');
UPDATE debtor_accounts SET cif = '11002022012' WHERE account_no = '3361928374' AND (cif IS NULL OR cif = '');

-- Populate CIF for existing commitments from debtor_accounts
UPDATE collection_commitments c
SET cif = d.cif
FROM debtor_accounts d
WHERE c.account_no = d.account_no AND (c.cif IS NULL OR c.cif = '');
