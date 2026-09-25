-- Migration: 000001_init_schema.down.sql
-- Description: Drop tables and indexes for rollback

DROP INDEX IF EXISTS idx_alert_logs_processed_severity;
DROP INDEX IF EXISTS idx_assignments_account_no;
DROP INDEX IF EXISTS idx_assignments_officer_id;
DROP INDEX IF EXISTS idx_commitments_account_no;
DROP INDEX IF EXISTS idx_commitments_status;
DROP INDEX IF EXISTS idx_commitments_date;
DROP INDEX IF EXISTS idx_debtor_accounts_debtor_name;
DROP INDEX IF EXISTS idx_debtor_accounts_product_type;
DROP INDEX IF EXISTS idx_debtor_accounts_exposure_tier;

DROP TABLE IF EXISTS alert_logs;
DROP TABLE IF EXISTS collection_commitments;
DROP TABLE IF EXISTS account_assignments;
DROP TABLE IF EXISTS debtor_accounts;
DROP TABLE IF EXISTS officers;
