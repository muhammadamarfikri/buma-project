-- Migration: 000006_convert_ids_to_uuid.up.sql
-- Description: Convert auto-incrementing ID primary keys to UUID v4 in collection_commitments, account_assignments, and alert_logs

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Convert id in collection_commitments to UUID
ALTER TABLE collection_commitments
ALTER COLUMN id DROP DEFAULT,
ALTER COLUMN id SET DATA TYPE UUID USING (gen_random_uuid()),
ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- Convert id in account_assignments to UUID
ALTER TABLE account_assignments
ALTER COLUMN id DROP DEFAULT,
ALTER COLUMN id SET DATA TYPE UUID USING (gen_random_uuid()),
ALTER COLUMN id SET DEFAULT gen_random_uuid();

-- Convert id in alert_logs to UUID
ALTER TABLE alert_logs
ALTER COLUMN id DROP DEFAULT,
ALTER COLUMN id SET DATA TYPE UUID USING (gen_random_uuid()),
ALTER COLUMN id SET DEFAULT gen_random_uuid();
