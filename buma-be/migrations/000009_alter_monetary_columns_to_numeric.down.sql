-- Migration: 000009_alter_monetary_columns_to_numeric.down.sql
-- Description: Rollback monetary columns type alter

ALTER TABLE collection_commitments 
    ALTER COLUMN nominal TYPE NUMERIC(18, 2),
    ALTER COLUMN credit_limit TYPE NUMERIC(18, 2),
    ALTER COLUMN outstanding_balance TYPE NUMERIC(18, 2);

ALTER TABLE debtor_accounts 
    ALTER COLUMN credit_limit TYPE NUMERIC(18, 2),
    ALTER COLUMN outstanding_balance TYPE NUMERIC(18, 2);
