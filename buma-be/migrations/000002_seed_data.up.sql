-- Migration: 000002_seed_data.up.sql
-- Description: Seed initial debtor records, officers, assignments, and commitments matching buma-fe frontend dataset

-- 1. Insert Officers
INSERT INTO officers (officer_id, full_name, email, pairing_team, phone) VALUES
('OFF-001', 'Budi Santoso', 'budi.santoso@buma.co.id', 'Desk 01', '081211112222'),
('OFF-002', 'Siti Rahma', 'siti.rahma@buma.co.id', 'Desk 02', '081233334444'),
('OFF-003', 'Ahmad Dahlan', 'ahmad.dahlan@buma.co.id', 'Desk 03', '081255556666'),
('OFF-004', 'Dian Sastro', 'dian.sastro@buma.co.id', 'Desk 04', '081277778888')
ON CONFLICT (officer_id) DO NOTHING;

-- 2. Insert Debtor Accounts
INSERT INTO debtor_accounts (account_no, debtor_id, debtor_name, product_type, credit_limit, outstanding_balance, exposure_tier, phone) VALUES
('1029384756', 'DBT-001', 'PT Nusantara Jaya Abadi', 'KMK', 5000000000.00, 4250000000.00, 'Tier 1', '081298765432'),
('2093847561', 'DBT-002', 'Hendra Wijaya', 'KPR', 1200000000.00, 980000000.00, 'Tier 2', '081387654321'),
('3084756192', 'DBT-003', 'CV Karya Utama Mandiri', 'KI', 8500000000.00, 7100000000.00, 'Tier 1', '081123456789'),
('4075619283', 'DBT-004', 'Dewi Lestari', 'KKM', 350000000.00, 210000000.00, 'Tier 3', '085678901234'),
('5061928374', 'DBT-005', 'Bambang Sukoco', 'KPR', 750000000.00, 620000000.00, 'Tier 2', '087812345678'),
('6019283745', 'DBT-006', 'PT Sinar Agro Makmur', 'KMK', 15000000000.00, 13800000000.00, 'Tier 1', '081901234567'),
('7092837465', 'DBT-007', 'Agus Pratama', 'KKM', 500000000.00, 410000000.00, 'Tier 3', '082123456789'),
('8028374651', 'DBT-008', 'Rina Kusuma', 'KPR', 1800000000.00, 1450000000.00, 'Tier 2', '083890123456'),
('9037465182', 'DBT-009', 'PT Megah Konstruksi Indonesia', 'KI', 12000000000.00, 9900000000.00, 'Tier 1', '081567890123'),
('1147561928', 'DBT-010', 'Eko Prasetyo', 'KMK', 2200000000.00, 1950000000.00, 'Tier 2', '081789012345'),
('2256192837', 'DBT-011', 'Maya Indah', 'KKM', 250000000.00, 180000000.00, 'Tier 3', '082290123456'),
('3361928374', 'DBT-012', 'Fajri Ramadhan', 'KPR', 950000000.00, 820000000.00, 'Tier 3', '085701234567')
ON CONFLICT (account_no) DO NOTHING;

-- 3. Insert Account Assignments
INSERT INTO account_assignments (account_no, officer_id) VALUES
('1029384756', 'OFF-001'),
('2093847561', 'OFF-002'),
('3084756192', 'OFF-001'),
('4075619283', 'OFF-003'),
('5061928374', 'OFF-002'),
('6019283745', 'OFF-004'),
('7092837465', 'OFF-003'),
('8028374651', 'OFF-002'),
('9037465182', 'OFF-001'),
('1147561928', 'OFF-004'),
('2256192837', 'OFF-003'),
('3361928374', 'OFF-002')
ON CONFLICT DO NOTHING;

-- 4. Insert Initial Commitments
INSERT INTO collection_commitments (account_no, commitment_date, status, reason, remarks, nominal) VALUES
('1029384756', '2026-09-06', 'Janji Bayar (PTP)', 'Janji pelunasan cicilan via transfer', 'Debitur berjanji melunasi cicilan Rp 250jt via transfer sore ini', 250000000.00),
('2093847561', '2026-09-07', 'Janji Call Back', 'Minta dihubungi lusa', 'Debitur meminta menghubungi kembali besok jam 10:00 WIB', 0.00),
('3084756192', '2026-09-06', 'Janji Bayar (PTP)', 'Penyetoran pasca SP-2', 'SP-2 telah dikirimkan. Debitur bersedia setor Rp 500jt', 500000000.00),
('4075619283', '2026-09-10', 'Restrukturisasi', 'Pengajuan keringanan tenor', 'Pengajuan perpanjangan tenor kredit sedang dalam review komite', 0.00),
('5061928374', '2026-09-06', 'Janji Bayar (PTP)', 'Autodebet angsuran', 'Komitmen pembayaran angsuran Rp 15jt via autodebet malam ini', 15000000.00),
('6019283745', '2026-09-05', 'Penolakan (Refusal)', 'Dispute tagihan', 'Menolak pembayaran karena klaim dispute tagihan. Perlu mediasi legal', 0.00),
('7092837465', '2026-09-12', 'Tidak Ada Respon', 'Belum membalas pesan', 'Wa terkirim centang duabiru, belum ada balasan dari debitur', 0.00),
('8028374651', '2026-09-08', 'Janji Call Back', 'Konfirmasi pihak keluarga', 'Suami debitur berjanji mengabarkan jadwal pelunasan lusa', 0.00),
('9037465182', '2026-09-06', 'Janji Bayar (PTP)', 'Pencairan termin proyek', 'Pencairan termin proyek hari ini. Komitmen setor Rp 750jt', 750000000.00),
('1147561928', '2026-09-09', 'Janji Bayar (PTP)', 'Pelunasan piutang usaha', 'Debitur menunggu pelunasan piutang usaha minggu depan', 0.00),
('2256192837', '2026-09-15', 'Restrukturisasi', 'Pengajuan keringanan bunga', 'Dokumen keringanan bunga telah diserahkan ke cabang pembantu', 0.00),
('3361928374', '2026-09-06', 'Janji Bayar (PTP)', 'Setor via ATM', 'Setor denda dan cicilan via ATM sebelum pukul 21.00 WIB', 0.00);
