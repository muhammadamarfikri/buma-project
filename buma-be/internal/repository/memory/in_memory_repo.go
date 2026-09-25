package memory

import (
	"context"
	"fmt"
	"math"
	"sort"
	"strings"
	"time"

	"buma-be/internal/domain"
	"buma-be/internal/repository"
	"buma-be/internal/wording"
)

type inMemoryDebtorRepo struct {
	debtors []domain.DebtorAccount
}

func NewInMemoryDebtorRepository() repository.DebtorRepository {
	initial := []domain.DebtorAccount{
		{AccountNo: "1029384756", DebtorID: "DBT-001", DebtorName: "PT Nusantara Jaya Abadi", ProductType: "KMK", CreditLimit: 5000000000, OutstandingBalance: 4250000000, ExposureTier: "Tier 1", OfficerID: "OFF-001", OfficerName: "Budi Santoso", PairingTeam: "Desk 01", CommitmentDate: "2026-09-06", CommitmentStatus: "On Progress/Implementasi", Remarks: "Debitur berjanji melunasi cicilan Rp 250jt via transfer sore ini", NominalCommitment: 250000000, Phone: "081298765432"},
		{AccountNo: "2093847561", DebtorID: "DBT-002", DebtorName: "Hendra Wijaya", ProductType: "KPR", CreditLimit: 1200000000, OutstandingBalance: 980000000, ExposureTier: "Tier 2", OfficerID: "OFF-002", OfficerName: "Siti Rahma", PairingTeam: "Desk 02", CommitmentDate: "2026-09-07", CommitmentStatus: "On Progress/Implementasi", Remarks: "Debitur meminta menghubungi kembali besok jam 10:00 WIB", Phone: "081387654321"},
		{AccountNo: "3084756192", DebtorID: "DBT-003", DebtorName: "CV Karya Utama Mandiri", ProductType: "KI", CreditLimit: 8500000000, OutstandingBalance: 7100000000, ExposureTier: "Tier 1", OfficerID: "OFF-001", OfficerName: "Budi Santoso", PairingTeam: "Desk 01", CommitmentDate: "2026-09-06", CommitmentStatus: "On Progress/Implementasi", Remarks: "SP-2 telah dikirimkan. Debitur bersedia setor Rp 500jt", NominalCommitment: 500000000, Phone: "081123456789"},
		{AccountNo: "4075619283", DebtorID: "DBT-004", DebtorName: "Dewi Lestari", ProductType: "KKM", CreditLimit: 350000000, OutstandingBalance: 210000000, ExposureTier: "Tier 3", OfficerID: "OFF-003", OfficerName: "Ahmad Dahlan", PairingTeam: "Desk 03", CommitmentDate: "2026-09-10", CommitmentStatus: "On Progress/Implementasi", Remarks: "Pengajuan perpanjangan tenor kredit sedang dalam review komite", Phone: "085678901234"},
		{AccountNo: "5061928374", DebtorID: "DBT-005", DebtorName: "Bambang Sukoco", ProductType: "KPR", CreditLimit: 750000000, OutstandingBalance: 620000000, ExposureTier: "Tier 2", OfficerID: "OFF-002", OfficerName: "Siti Rahma", PairingTeam: "Desk 02", CommitmentDate: "2026-09-06", CommitmentStatus: "On Progress/Implementasi", Remarks: "Komitmen pembayaran angsuran Rp 15jt via autodebet malam ini", NominalCommitment: 15000000, Phone: "087812345678"},
		{AccountNo: "6019283745", DebtorID: "DBT-006", DebtorName: "PT Sinar Agro Makmur", ProductType: "KMK", CreditLimit: 15000000000, OutstandingBalance: 13800000000, ExposureTier: "Tier 1", OfficerID: "OFF-004", OfficerName: "Dian Sastro", PairingTeam: "Desk 04", CommitmentDate: "2026-09-05", CommitmentStatus: "Melewati Komitmen", Remarks: "Menolak pembayaran karena klaim dispute tagihan. Perlu mediasi legal", Phone: "081901234567"},
		{AccountNo: "7092837465", DebtorID: "DBT-007", DebtorName: "Agus Pratama", ProductType: "KKM", CreditLimit: 500000000, OutstandingBalance: 410000000, ExposureTier: "Tier 3", OfficerID: "OFF-003", OfficerName: "Ahmad Dahlan", PairingTeam: "Desk 03", CommitmentDate: "2026-09-12", CommitmentStatus: "Belum Dilaksanakan", Remarks: "Wa terkirim centang duabiru, belum ada balasan dari debitur", Phone: "082123456789"},
		{AccountNo: "8028374651", DebtorID: "DBT-008", DebtorName: "Rina Kusuma", ProductType: "KPR", CreditLimit: 1800000000, OutstandingBalance: 1450000000, ExposureTier: "Tier 2", OfficerID: "OFF-002", OfficerName: "Siti Rahma", PairingTeam: "Desk 02", CommitmentDate: "2026-09-08", CommitmentStatus: "On Progress/Implementasi", Remarks: "Suami debitur berjanji mengabarkan jadwal pelunasan lusa", Phone: "083890123456"},
		{AccountNo: "9037465182", DebtorID: "DBT-009", DebtorName: "PT Megah Konstruksi Indonesia", ProductType: "KI", CreditLimit: 12000000000, OutstandingBalance: 9900000000, ExposureTier: "Tier 1", OfficerID: "OFF-001", OfficerName: "Budi Santoso", PairingTeam: "Desk 01", CommitmentDate: "2026-09-06", CommitmentStatus: "Telah Dilaksanakan", Remarks: "Pencairan termin proyek hari ini. Komitmen setor Rp 750jt", NominalCommitment: 750000000, Phone: "081567890123"},
		{AccountNo: "1147561928", DebtorID: "DBT-010", DebtorName: "Eko Prasetyo", ProductType: "KMK", CreditLimit: 2200000000, OutstandingBalance: 1950000000, ExposureTier: "Tier 2", OfficerID: "OFF-004", OfficerName: "Dian Sastro", PairingTeam: "Desk 04", CommitmentDate: "2026-09-09", CommitmentStatus: "On Progress/Implementasi", Remarks: "Debitur menunggu pelunasan piutang usaha minggu depan", Phone: "081789012345"},
		{AccountNo: "2256192837", DebtorID: "DBT-011", DebtorName: "Maya Indah", ProductType: "KKM", CreditLimit: 250000000, OutstandingBalance: 180000000, ExposureTier: "Tier 3", OfficerID: "OFF-003", OfficerName: "Ahmad Dahlan", PairingTeam: "Desk 03", CommitmentDate: "2026-09-15", CommitmentStatus: "On Progress/Implementasi", Remarks: "Dokumen keringanan bunga telah diserahkan ke cabang pembantu", Phone: "082290123456"},
		{AccountNo: "3361928374", DebtorID: "DBT-012", DebtorName: "Fajri Ramadhan", ProductType: "KPR", CreditLimit: 950000000, OutstandingBalance: 820000000, ExposureTier: "Tier 3", OfficerID: "OFF-002", OfficerName: "Siti Rahma", PairingTeam: "Desk 02", CommitmentDate: "2026-09-06", CommitmentStatus: "On Progress/Implementasi", Remarks: "Setor denda dan cicilan via ATM sebelum pukul 21.00 WIB", Phone: "085701234567"},
	}
	return &inMemoryDebtorRepo{debtors: initial}
}

func (r *inMemoryDebtorRepo) GetMonitoringData(ctx context.Context, filter domain.DebtorFilter) (*domain.MonitoringResponse, error) {
	var totalBakiDebet float64
	var tier1Count int
	var commitmentsToday int
	todayStr := wording.FormatDate("")

	for _, d := range r.debtors {
		totalBakiDebet += d.OutstandingBalance
		if d.ExposureTier == "Tier 1" {
			tier1Count++
		}
		if d.CommitmentDate == todayStr {
			commitmentsToday++
		}
	}

	tier1Ratio := 0.0
	if len(r.debtors) > 0 {
		tier1Ratio = math.Round((float64(tier1Count)/float64(len(r.debtors))*100)*10) / 10
	}

	metrics := domain.MonitoringMetrics{
		TotalDebtors:     len(r.debtors),
		TotalBakiDebet:   totalBakiDebet,
		Tier1Count:       tier1Count,
		Tier1Ratio:       tier1Ratio,
		CommitmentsToday: commitmentsToday,
	}

	page := filter.Page
	if page < 1 {
		page = 1
	}
	limit := filter.Limit
	if limit < 1 {
		limit = 10
	}

	// Filter
	var filtered []domain.DebtorAccount
	for _, d := range r.debtors {
		if filter.Search != "" {
			if !containsInsensitive(d.DebtorName, filter.Search) && !containsInsensitive(d.AccountNo, filter.Search) {
				continue
			}
		}
		if filter.Tier != "" && filter.Tier != "ALL" && d.ExposureTier != filter.Tier {
			continue
		}
		if filter.OfficerName != "" && filter.OfficerName != "ALL" {
			if !containsInsensitive(d.OfficerName, filter.OfficerName) && !containsInsensitive(d.OfficerID, filter.OfficerName) {
				continue
			}
		}
		if filter.OfficerID != "" && filter.OfficerID != "ALL" && d.OfficerID != filter.OfficerID {
			continue
		}
		if filter.Product != "" && filter.Product != "ALL" && d.ProductType != filter.Product {
			continue
		}
		filtered = append(filtered, d)
	}

	if filter.SortBy != "" {
		sortDir := strings.ToLower(filter.SortDir)
		sort.Slice(filtered, func(i, j int) bool {
			var valI, valJ string
			switch filter.SortBy {
			case "noRekening", "account_no":
				valI, valJ = filtered[i].AccountNo, filtered[j].AccountNo
			case "namaDebitur", "debtor_name":
				valI, valJ = filtered[i].DebtorName, filtered[j].DebtorName
			case "tierEksposur", "exposure_tier":
				valI, valJ = filtered[i].ExposureTier, filtered[j].ExposureTier
			case "produk", "product":
				valI, valJ = filtered[i].ProductType, filtered[j].ProductType
			case "tglKomitmen", "commitment_date":
				valI, valJ = filtered[i].CommitmentDate, filtered[j].CommitmentDate
			default:
				valI, valJ = filtered[i].AccountNo, filtered[j].AccountNo
			}
			if sortDir == "asc" {
				return valI < valJ
			}
			return valI > valJ
		})
	}

	totalItems := len(filtered)
	totalPages := 1
	if totalItems > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(limit)))
	}

	start := (page - 1) * limit
	if start > totalItems {
		start = totalItems
	}
	end := start + limit
	if end > totalItems {
		end = totalItems
	}

	sliced := filtered[start:end]
	if sliced == nil {
		sliced = []domain.DebtorAccount{}
	}

	return &domain.MonitoringResponse{
		Metrics: metrics,
		Data:    sliced,
		Pagination: domain.PaginationMeta{
			CurrentPage:  page,
			ItemsPerPage: limit,
			TotalItems:   totalItems,
			TotalPages:   totalPages,
		},
	}, nil
}

func (r *inMemoryDebtorRepo) GetDebtorByAccountNo(ctx context.Context, accountNo string) (*domain.DebtorAccount, error) {
	for _, d := range r.debtors {
		if d.AccountNo == accountNo {
			return &d, nil
		}
	}
	return nil, nil
}

func (r *inMemoryDebtorRepo) GetHighExposureTier1Accounts(ctx context.Context) ([]domain.DebtorAccount, error) {
	var res []domain.DebtorAccount
	for _, d := range r.debtors {
		if d.ExposureTier == "Tier 1" {
			res = append(res, d)
		}
	}
	return res, nil
}

func containsInsensitive(str, sub string) bool {
	if sub == "" {
		return true
	}
	return strings.Contains(strings.ToLower(str), strings.ToLower(sub))
}

type inMemoryOfficerRepo struct {
	officers []domain.Officer
}

func NewInMemoryOfficerRepository() repository.OfficerRepository {
	initial := []domain.Officer{
		{OfficerID: "OFF-001", FullName: "Budi Santoso", Email: "budi.santoso@buma.co.id", PairingTeam: "Desk 01", Phone: "081211112222"},
		{OfficerID: "OFF-002", FullName: "Siti Rahma", Email: "siti.rahma@buma.co.id", PairingTeam: "Desk 02", Phone: "081233334444"},
		{OfficerID: "OFF-003", FullName: "Ahmad Dahlan", Email: "ahmad.dahlan@buma.co.id", PairingTeam: "Desk 03", Phone: "081255556666"},
		{OfficerID: "OFF-004", FullName: "Dian Sastro", Email: "dian.sastro@buma.co.id", PairingTeam: "Desk 04", Phone: "081277778888"},
	}
	return &inMemoryOfficerRepo{officers: initial}
}

func (r *inMemoryOfficerRepo) CreateOfficer(ctx context.Context, officer domain.Officer) (*domain.Officer, error) {
	r.officers = append(r.officers, officer)
	return &officer, nil
}

func (r *inMemoryOfficerRepo) GetAllOfficers(ctx context.Context) ([]domain.Officer, error) {
	return r.officers, nil
}

func (r *inMemoryOfficerRepo) GetOfficerByID(ctx context.Context, officerID string) (*domain.Officer, error) {
	for _, o := range r.officers {
		if o.OfficerID == officerID {
			return &o, nil
		}
	}
	return nil, nil
}

type inMemoryCommitmentRepo struct {
	commitments []domain.CollectionCommitment
}

func NewInMemoryCommitmentRepository() repository.CommitmentRepository {
	now := time.Now()
	initial := []domain.CollectionCommitment{
		{ID: "c101", AccountNo: "1029384756", CommitmentDate: "2026-09-06", OfficerID: "OFF-001", OfficerName: "Budi Santoso", OfficerPairID: "OFF-002", OfficerPairName: "Desk 01", Status: "On Progress/Implementasi", Remarks: "Debitur berjanji melunasi cicilan Rp 250jt via transfer sore ini", Nominal: 250000000, Reason: "On Progress/Implementasi", CreditLimit: 5000000000, OutstandingBalance: 4250000000, ExposureTier: "Tier 1", Product: "KMK", CreatedAt: now, UpdatedAt: now},
		{ID: "c102", AccountNo: "2093847561", CommitmentDate: "2026-09-07", OfficerID: "OFF-002", OfficerName: "Siti Rahma", OfficerPairID: "OFF-001", OfficerPairName: "Desk 02", Status: "On Progress/Implementasi", Remarks: "Debitur meminta menghubungi kembali besok jam 10:00 WIB", Reason: "On Progress/Implementasi", CreditLimit: 1200000000, OutstandingBalance: 980000000, ExposureTier: "Tier 2", Product: "KPR", CreatedAt: now, UpdatedAt: now},
		{ID: "c103", AccountNo: "3084756192", CommitmentDate: "2026-09-06", OfficerID: "OFF-001", OfficerName: "Budi Santoso", OfficerPairID: "OFF-002", OfficerPairName: "Desk 01", Status: "On Progress/Implementasi", Remarks: "SP-2 telah dikirimkan. Debitur bersedia setor Rp 500jt", Nominal: 500000000, Reason: "On Progress/Implementasi", CreditLimit: 8500000000, OutstandingBalance: 7100000000, ExposureTier: "Tier 1", Product: "KI", CreatedAt: now, UpdatedAt: now},
		{ID: "c104", AccountNo: "4075619283", CommitmentDate: "2026-09-10", OfficerID: "OFF-003", OfficerName: "Ahmad Dahlan", OfficerPairID: "OFF-004", OfficerPairName: "Desk 03", Status: "On Progress/Implementasi", Remarks: "Pengajuan perpanjangan tenor kredit sedang dalam review komite", Reason: "On Progress/Implementasi", CreditLimit: 350000000, OutstandingBalance: 210000000, ExposureTier: "Tier 3", Product: "KKM", CreatedAt: now, UpdatedAt: now},
		{ID: "c105", AccountNo: "5061928374", CommitmentDate: "2026-09-06", OfficerID: "OFF-002", OfficerName: "Siti Rahma", OfficerPairID: "OFF-001", OfficerPairName: "Desk 02", Status: "On Progress/Implementasi", Remarks: "Komitmen pembayaran angsuran Rp 15jt via autodebet malam ini", Nominal: 15000000, Reason: "On Progress/Implementasi", CreditLimit: 750000000, OutstandingBalance: 620000000, ExposureTier: "Tier 2", Product: "KPR", CreatedAt: now, UpdatedAt: now},
	}
	return &inMemoryCommitmentRepo{commitments: initial}
}

func (r *inMemoryCommitmentRepo) CreateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error) {
	c := domain.CollectionCommitment{
		ID:                 fmt.Sprintf("c%d", len(r.commitments)+100),
		AccountNo:          req.AccountNo,
		CommitmentDate:     req.CommitmentDate,
		OfficerID:          req.OfficerID,
		OfficerPairID:      req.OfficerPairID,
		Status:             req.Status,
		Remarks:            req.Remarks,
		Nominal:            req.CreditLimit,
		Reason:             req.Reason,
		CreditLimit:        req.CreditLimit,
		OutstandingBalance: req.OutstandingBalance,
		ExposureTier:       req.ExposureTier,
		Product:            req.Product,
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}
	r.commitments = append([]domain.CollectionCommitment{c}, r.commitments...)
	return &c, nil
}

func (r *inMemoryCommitmentRepo) UpdateCommitmentByReason(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error) {
	for i, c := range r.commitments {
		if c.AccountNo == req.AccountNo && strings.EqualFold(c.Reason, req.Reason) {
			r.commitments[i].CommitmentDate = req.CommitmentDate
			r.commitments[i].OfficerID = req.OfficerID
			r.commitments[i].OfficerPairID = req.OfficerPairID
			r.commitments[i].Status = req.Status
			r.commitments[i].Remarks = req.Remarks
			r.commitments[i].Nominal = req.CreditLimit
			if req.Product != "" {
				r.commitments[i].Product = req.Product
			}
			if req.ExposureTier != "" {
				r.commitments[i].ExposureTier = req.ExposureTier
			}
			if req.CreditLimit > 0 {
				r.commitments[i].CreditLimit = req.CreditLimit
			}
			if req.OutstandingBalance > 0 {
				r.commitments[i].OutstandingBalance = req.OutstandingBalance
			}
			r.commitments[i].UpdatedAt = time.Now()
			return &r.commitments[i], nil
		}
	}
	return nil, fmt.Errorf("commitment with reason '%s' not found for account %s", req.Reason, req.AccountNo)
}

func (r *inMemoryCommitmentRepo) CreateOrUpdateCommitment(ctx context.Context, req domain.CreateCommitmentRequest) (*domain.CollectionCommitment, error) {
	updated, err := r.UpdateCommitmentByReason(ctx, req)
	if err == nil && updated != nil {
		return updated, nil
	}
	return r.CreateCommitment(ctx, req)
}

func (r *inMemoryCommitmentRepo) GetCommitmentsByAccountNo(ctx context.Context, accountNo string) ([]domain.CollectionCommitment, error) {
	var list []domain.CollectionCommitment
	for _, c := range r.commitments {
		if c.AccountNo == accountNo {
			list = append(list, c)
		}
	}
	if list == nil {
		list = []domain.CollectionCommitment{}
	}
	return list, nil
}

func (r *inMemoryCommitmentRepo) GetCommitmentsByDate(ctx context.Context, dateStr string) ([]domain.CollectionCommitment, error) {
	var list []domain.CollectionCommitment
	for _, c := range r.commitments {
		if c.CommitmentDate == dateStr {
			list = append(list, c)
		}
	}
	if list == nil {
		list = []domain.CollectionCommitment{}
	}
	return list, nil
}

func (r *inMemoryCommitmentRepo) GetCommitments(ctx context.Context, filter domain.CommitmentFilter) (*domain.PaginatedCommitmentResponse, error) {
	page := filter.Page
	if page < 1 {
		page = 1
	}
	limit := filter.Limit
	if limit < 1 {
		limit = 10
	}

	var filtered []domain.CollectionCommitment
	for _, c := range r.commitments {
		if filter.AccountNo != "" && c.AccountNo != filter.AccountNo {
			continue
		}
		if filter.Status != "" && filter.Status != "ALL" && !strings.EqualFold(c.Status, filter.Status) {
			continue
		}
		if filter.Tier != "" && filter.Tier != "ALL" && !strings.EqualFold(c.ExposureTier, filter.Tier) {
			continue
		}
		if filter.Product != "" && filter.Product != "ALL" && !strings.EqualFold(c.Product, filter.Product) {
			continue
		}
		if filter.Pengelola != "" && filter.Pengelola != "ALL" {
			if !containsInsensitive(c.OfficerID, filter.Pengelola) &&
				!containsInsensitive(c.OfficerName, filter.Pengelola) &&
				!containsInsensitive(filter.Pengelola, c.OfficerID) &&
				!containsInsensitive(filter.Pengelola, c.OfficerName) {
				continue
			}
		}
		if filter.Search != "" {
			search := filter.Search
			match := containsInsensitive(c.AccountNo, search) ||
				containsInsensitive(c.Status, search) ||
				containsInsensitive(c.Reason, search) ||
				containsInsensitive(c.Remarks, search) ||
				containsInsensitive(c.OfficerName, search) ||
				containsInsensitive(c.OfficerID, search) ||
				containsInsensitive(c.Product, search) ||
				containsInsensitive(c.ExposureTier, search)
			if !match {
				continue
			}
		}
		filtered = append(filtered, c)
	}

	if filter.SortBy != "" {
		sortDir := strings.ToLower(filter.SortDir)
		sort.Slice(filtered, func(i, j int) bool {
			var valI, valJ string
			switch filter.SortBy {
			case "account_no", "noRekening":
				valI, valJ = filtered[i].AccountNo, filtered[j].AccountNo
			case "commitment_date", "tglKomitmen":
				valI, valJ = filtered[i].CommitmentDate, filtered[j].CommitmentDate
			case "status", "statusKomitmen":
				valI, valJ = filtered[i].Status, filtered[j].Status
			case "tier", "exposure_tier", "tierEksposur":
				valI, valJ = filtered[i].ExposureTier, filtered[j].ExposureTier
			case "product", "produk":
				valI, valJ = filtered[i].Product, filtered[j].Product
			default:
				valI, valJ = filtered[i].CreatedAt.Format(time.RFC3339), filtered[j].CreatedAt.Format(time.RFC3339)
			}
			if sortDir == "asc" {
				return valI < valJ
			}
			return valI > valJ
		})
	}

	totalItems := len(filtered)
	totalPages := 1
	if totalItems > 0 {
		totalPages = int(math.Ceil(float64(totalItems) / float64(limit)))
	}

	start := (page - 1) * limit
	if start > totalItems {
		start = totalItems
	}
	end := start + limit
	if end > totalItems {
		end = totalItems
	}

	sliced := filtered[start:end]
	if sliced == nil {
		sliced = []domain.CollectionCommitment{}
	}

	return &domain.PaginatedCommitmentResponse{
		Data: sliced,
		Pagination: domain.PaginationMeta{
			CurrentPage:  page,
			ItemsPerPage: limit,
			TotalItems:   totalItems,
			TotalPages:   totalPages,
		},
	}, nil
}

func (r *inMemoryCommitmentRepo) GetCommitmentByID(ctx context.Context, id string) (*domain.CollectionCommitment, error) {
	for _, c := range r.commitments {
		if c.ID == id {
			return &c, nil
		}
	}
	return nil, fmt.Errorf("commitment with id %s not found", id)
}

func (r *inMemoryCommitmentRepo) UpdateCommitmentByID(ctx context.Context, id string, req domain.UpdateCommitmentRequest) (*domain.CollectionCommitment, error) {
	for i, c := range r.commitments {
		if c.ID == id || c.AccountNo == id || (req.AccountNo != "" && c.AccountNo == req.AccountNo) {
			if req.CommitmentDate != "" {
				r.commitments[i].CommitmentDate = req.CommitmentDate
			}
			if req.Status != "" {
				r.commitments[i].Status = req.Status
			}
			if req.Reason != "" {
				r.commitments[i].Reason = req.Reason
			}
			if req.Remarks != "" {
				r.commitments[i].Remarks = req.Remarks
			}
			r.commitments[i].Nominal = req.Nominal
			if req.OfficerID != "" {
				r.commitments[i].OfficerID = req.OfficerID
			}
			if req.OfficerPairID != "" {
				r.commitments[i].OfficerPairID = req.OfficerPairID
			}
			if req.CreditLimit > 0 {
				r.commitments[i].CreditLimit = req.CreditLimit
			}
			if req.OutstandingBalance > 0 {
				r.commitments[i].OutstandingBalance = req.OutstandingBalance
			}
			if req.ExposureTier != "" {
				r.commitments[i].ExposureTier = req.ExposureTier
			}
			if req.Product != "" {
				r.commitments[i].Product = req.Product
			}
			r.commitments[i].UpdatedAt = time.Now()
			return &r.commitments[i], nil
		}
	}
	return nil, fmt.Errorf("commitment with id %s not found", id)
}
