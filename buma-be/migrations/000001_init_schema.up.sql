-- Migration: 000001_init_schema.up.sql
-- Description: Create tables, constraints, foreign keys, and indexes for BUMA Debt Collection System

CREATE TABLE IF NOT EXISTS officers (
    officer_id VARCHAR(50) PRIMARY KEY,
    full_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    pairing_team VARCHAR(100) NOT NULL,
    phone VARCHAR(50),
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS debtor_accounts (
    account_no VARCHAR(50) PRIMARY KEY,
    debtor_id VARCHAR(50) NOT NULL,
    debtor_name VARCHAR(255) NOT NULL,
    product_type VARCHAR(50) NOT NULL,
    credit_limit NUMERIC(18, 2) NOT NULL DEFAULT 0,
    outstanding_balance NUMERIC(18, 2) NOT NULL DEFAULT 0,
    exposure_tier VARCHAR(20) NOT NULL,
    phone VARCHAR(50) NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS account_assignments (
    id SERIAL PRIMARY KEY,
    account_no VARCHAR(50) NOT NULL REFERENCES debtor_accounts(account_no) ON DELETE CASCADE,
    officer_id VARCHAR(50) NOT NULL REFERENCES officers(officer_id) ON DELETE CASCADE,
    assigned_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT uq_account_officer UNIQUE (account_no, officer_id)
);

CREATE TABLE IF NOT EXISTS collection_commitments (
    id SERIAL PRIMARY KEY,
    account_no VARCHAR(50) NOT NULL REFERENCES debtor_accounts(account_no) ON DELETE CASCADE,
    commitment_date DATE NOT NULL,
    officer_id VARCHAR(50) REFERENCES officers(officer_id) ON DELETE SET NULL,
    officer_pair_id VARCHAR(50) REFERENCES officers(officer_id) ON DELETE SET NULL,
    status VARCHAR(50) NOT NULL,
    reason TEXT,
    remarks TEXT,
    nominal NUMERIC(18, 2) DEFAULT 0,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS alert_logs (
    id SERIAL PRIMARY KEY,
    account_no VARCHAR(50) NOT NULL REFERENCES debtor_accounts(account_no) ON DELETE CASCADE,
    alert_type VARCHAR(50) NOT NULL,
    severity VARCHAR(20) NOT NULL,
    message TEXT NOT NULL,
    is_processed BOOLEAN NOT NULL DEFAULT FALSE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- ============================================================================
-- B-TREE INDEXES FOR SEARCH AND DYNAMIC FILTERING PERFORMANCE
-- ============================================================================

CREATE INDEX IF NOT EXISTS idx_debtor_accounts_exposure_tier ON debtor_accounts(exposure_tier);
CREATE INDEX IF NOT EXISTS idx_debtor_accounts_product_type ON debtor_accounts(product_type);
CREATE INDEX IF NOT EXISTS idx_debtor_accounts_debtor_name ON debtor_accounts(debtor_name);

CREATE INDEX IF NOT EXISTS idx_commitments_date ON collection_commitments(commitment_date);
CREATE INDEX IF NOT EXISTS idx_commitments_status ON collection_commitments(status);
CREATE INDEX IF NOT EXISTS idx_commitments_account_no ON collection_commitments(account_no);
CREATE INDEX IF NOT EXISTS idx_commitments_officer_id ON collection_commitments(officer_id);
CREATE INDEX IF NOT EXISTS idx_commitments_officer_pair_id ON collection_commitments(officer_pair_id);

CREATE INDEX IF NOT EXISTS idx_assignments_officer_id ON account_assignments(officer_id);
CREATE INDEX IF NOT EXISTS idx_assignments_account_no ON account_assignments(account_no);

CREATE INDEX IF NOT EXISTS idx_alert_logs_processed_severity ON alert_logs(is_processed, severity);
