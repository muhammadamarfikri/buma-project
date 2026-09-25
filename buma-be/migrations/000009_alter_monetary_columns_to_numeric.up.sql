-- Migration: 000009_alter_monetary_columns_to_numeric.up.sql
-- Description: Convert monetary columns (nominal, credit_limit, outstanding_balance) in collection_commitments and debtor_accounts to NUMERIC(18,2) to prevent 32-bit INTEGER overflow (error 22003) and support decimal monetary figures (error 22P02).

ALTER TABLE collection_commitments 
    ALTER COLUMN nominal TYPE NUMERIC(18, 2) USING nominal::NUMERIC(18, 2),
    ALTER COLUMN credit_limit TYPE NUMERIC(18, 2) USING credit_limit::NUMERIC(18, 2),
    ALTER COLUMN outstanding_balance TYPE NUMERIC(18, 2) USING outstanding_balance::NUMERIC(18, 2);

ALTER TABLE debtor_accounts 
    ALTER COLUMN credit_limit TYPE NUMERIC(18, 2) USING credit_limit::NUMERIC(18, 2),
    ALTER COLUMN outstanding_balance TYPE NUMERIC(18, 2) USING outstanding_balance::NUMERIC(18, 2);
