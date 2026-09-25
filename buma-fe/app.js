/**
 * BUMA Debt Collection Monitoring & Reminder Dispatch System
 * Client-Side JavaScript Application Logic (ES6+)
 */

// ==========================================
// 1. MOCK DATA INITIALIZATION
// ==========================================
const INITIAL_DEBTORS = [
  {
    id: "DBT-001",
    noRekening: "1029384756",
    namaDebitur: "PT Nusantara Jaya Abadi",
    produk: "KMK",
    limit: 5000000000,
    bakiDebet: 4250000000,
    tierEksposur: "Tier 1",
    pengelola: "Budi Santoso",
    pairing: "Desk 01",
    tglKomitmen: "2026-09-06",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Debitur berjanji melunasi cicilan Rp 250jt via transfer sore ini",
    phone: "081298765432"
  },
  {
    id: "DBT-002",
    noRekening: "2093847561",
    namaDebitur: "Hendra Wijaya",
    produk: "KPR",
    limit: 1200000000,
    bakiDebet: 980000000,
    tierEksposur: "Tier 2",
    pengelola: "Siti Rahma",
    pairing: "Desk 02",
    tglKomitmen: "2026-09-07",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Debitur meminta menghubungi kembali besok jam 10:00 WIB",
    phone: "081387654321"
  },
  {
    id: "DBT-003",
    noRekening: "3084756192",
    namaDebitur: "CV Karya Utama Mandiri",
    produk: "KI",
    limit: 8500000000,
    bakiDebet: 7100000000,
    tierEksposur: "Tier 1",
    pengelola: "Budi Santoso",
    pairing: "Desk 01",
    tglKomitmen: "2026-09-06",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "SP-2 telah dikirimkan. Debitur bersedia setor Rp 500jt",
    phone: "081123456789"
  },
  {
    id: "DBT-004",
    noRekening: "4075619283",
    namaDebitur: "Dewi Lestari",
    produk: "KKM",
    limit: 350000000,
    bakiDebet: 210000000,
    tierEksposur: "Tier 3",
    pengelola: "Ahmad Dahlan",
    pairing: "Desk 03",
    tglKomitmen: "2026-09-10",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Pengajuan perpanjangan tenor kredit sedang dalam review komite",
    phone: "085678901234"
  },
  {
    id: "DBT-005",
    noRekening: "5061928374",
    namaDebitur: "Bambang Sukoco",
    produk: "KPR",
    limit: 750000000,
    bakiDebet: 620000000,
    tierEksposur: "Tier 2",
    pengelola: "Siti Rahma",
    pairing: "Desk 02",
    tglKomitmen: "2026-09-06",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Komitmen pembayaran angsuran Rp 15jt via autodebet malam ini",
    phone: "087812345678"
  },
  {
    id: "DBT-006",
    noRekening: "6019283745",
    namaDebitur: "PT Sinar Agro Makmur",
    produk: "KMK",
    limit: 15000000000,
    bakiDebet: 13800000000,
    tierEksposur: "Tier 1",
    pengelola: "Dian Sastro",
    pairing: "Desk 04",
    tglKomitmen: "2026-09-05",
    statusKomitmen: "Melewati Komitmen",
    keterangan: "Menolak pembayaran karena klaim dispute tagihan. Perlu mediasi legal",
    phone: "081901234567"
  },
  {
    id: "DBT-007",
    noRekening: "7092837465",
    namaDebitur: "Agus Pratama",
    produk: "KKM",
    limit: 500000000,
    bakiDebet: 410000000,
    tierEksposur: "Tier 3",
    pengelola: "Ahmad Dahlan",
    pairing: "Desk 03",
    tglKomitmen: "2026-09-12",
    statusKomitmen: "Belum Dilaksanakan",
    keterangan: "Wa terkirim centang duabiru, belum ada balasan dari debitur",
    phone: "082123456789"
  },
  {
    id: "DBT-008",
    noRekening: "8028374651",
    namaDebitur: "Rina Kusuma",
    produk: "KPR",
    limit: 1800000000,
    bakiDebet: 1450000000,
    tierEksposur: "Tier 2",
    pengelola: "Siti Rahma",
    pairing: "Desk 02",
    tglKomitmen: "2026-09-08",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Suami debitur berjanji mengabarkan jadwal pelunasan lusa",
    phone: "083890123456"
  },
  {
    id: "DBT-009",
    noRekening: "9037465182",
    namaDebitur: "PT Megah Konstruksi Indonesia",
    produk: "KI",
    limit: 12000000000,
    bakiDebet: 9900000000,
    tierEksposur: "Tier 1",
    pengelola: "Budi Santoso",
    pairing: "Desk 01",
    tglKomitmen: "2026-09-06",
    statusKomitmen: "Telah Dilaksanakan",
    keterangan: "Pencairan termin proyek hari ini. Komitmen setor Rp 750jt",
    phone: "081567890123"
  },
  {
    id: "DBT-010",
    noRekening: "1147561928",
    namaDebitur: "Eko Prasetyo",
    produk: "KMK",
    limit: 2200000000,
    bakiDebet: 1950000000,
    tierEksposur: "Tier 2",
    pengelola: "Dian Sastro",
    pairing: "Desk 04",
    tglKomitmen: "2026-09-09",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Debitur menunggu pelunasan piutang usaha minggu depan",
    phone: "081789012345"
  },
  {
    id: "DBT-011",
    noRekening: "2256192837",
    namaDebitur: "Maya Indah",
    produk: "KKM",
    limit: 250000000,
    bakiDebet: 180000000,
    tierEksposur: "Tier 3",
    pengelola: "Ahmad Dahlan",
    pairing: "Desk 03",
    tglKomitmen: "2026-09-15",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Dokumen keringanan bunga telah diserahkan ke cabang pembantu",
    phone: "082290123456"
  },
  {
    id: "DBT-012",
    noRekening: "3361928374",
    namaDebitur: "Fajri Ramadhan",
    produk: "KPR",
    limit: 950000000,
    bakiDebet: 820000000,
    tierEksposur: "Tier 3",
    pengelola: "Siti Rahma",
    pairing: "Desk 02",
    tglKomitmen: "2026-09-06",
    statusKomitmen: "On Progress/Implementasi",
    keterangan: "Setor denda dan cicilan via ATM sebelum pukul 21.00 WIB",
    phone: "085701234567"
  }
];

// ==========================================
// 2. STATE MANAGEMENT
// ==========================================
const state = {
  debtors: [...INITIAL_DEBTORS],
  filteredDebtors: [],
  officersList: [],
  searchQuery: "",
  selectedTier: "ALL",
  selectedPengelola: "ALL",
  selectedProduk: "ALL",
  sortColumn: "bakiDebet",
  sortDirection: "desc",
  currentPage: 1,
  itemsPerPage: 10,
  serverPagination: null,
  activeWordingDebtor: null,
  activeWordingChannel: "whatsapp",
  activeWordingTone: "firm",
  activeCommitmentDebtor: null
};

// ==========================================
// 3. UTILITY FUNCTIONS
// ==========================================

/** Parses any raw/formatted input or string into pure numeric float64 value */
function parseMoney(val) {
  if (val === null || val === undefined || val === "") return 0;
  if (typeof val === "number") return isNaN(val) ? 0 : val;

  let str = String(val).trim();
  if (!str) return 0;

  // Strip currency prefixes/suffixes (e.g. "Rp", "RP") and invalid characters
  str = str.replace(/[^0-9.,\-]/g, "");
  if (!str) return 0;

  // Handle multiple dots (e.g. 5.000.000.000 in IDR locale)
  if ((str.match(/\./g) || []).length > 1) {
    str = str.replace(/\./g, "");
  }

  // Handle multiple commas (e.g. 5,000,000,000 in US locale)
  if ((str.match(/,/g) || []).length > 1) {
    str = str.replace(/,/g, "");
  }

  // Handle combined dot and comma notation
  if (str.includes(".") && str.includes(",")) {
    if (str.indexOf(".") < str.indexOf(",")) {
      str = str.replace(/\./g, "").replace(",", ".");
    } else {
      str = str.replace(/,/g, "");
    }
  } else if (str.includes(",")) {
    const parts = str.split(",");
    if (parts.length === 2 && parts[1].length !== 3) {
      str = str.replace(",", ".");
    } else {
      str = str.replace(/,/g, "");
    }
  } else if (str.includes(".")) {
    const parts = str.split(".");
    if (parts.length === 2 && parts[1].length === 3 && parts[0].length >= 1 && parts[0].length <= 3) {
      str = str.replace(".", "");
    }
  }

  const result = parseFloat(str);
  return isNaN(result) ? 0 : result;
}

/** Formats number into Indonesian Rupiah currency format */
function formatIDR(amount) {
  const num = typeof amount === "number" ? amount : parseMoney(amount);
  if (isNaN(num)) return "Rp 0";
  return new Intl.NumberFormat("id-ID", {
    style: "currency",
    currency: "IDR",
    maximumFractionDigits: 0
  }).format(num);
}

/** Helper to update live money format preview labels under modal input fields */
function updateMoneyInputPreviews() {
  const nomEl = document.getElementById("inputNominal");
  const limEl = document.getElementById("inputLimit");
  const bdEl = document.getElementById("inputBakiDebetInput");

  const nomPrev = document.getElementById("inputNominalPreview");
  const limPrev = document.getElementById("inputLimitPreview");
  const bdPrev = document.getElementById("inputBakiDebetPreview");

  if (nomEl && nomPrev) {
    const val = parseMoney(nomEl.value);
    nomPrev.textContent = val > 0 ? formatIDR(val) : "";
  }
  if (limEl && limPrev) {
    const val = parseMoney(limEl.value);
    limPrev.textContent = val > 0 ? formatIDR(val) : "";
  }
  if (bdEl && bdPrev) {
    const val = parseMoney(bdEl.value);
    bdPrev.textContent = val > 0 ? formatIDR(val) : "";
  }
}

/** Formats YYYY-MM-DD into readable Indonesian Date */
function formatDate(dateStr) {
  if (!dateStr) return "-";
  const date = new Date(dateStr);
  if (isNaN(date.getTime())) return dateStr;
  return new Intl.DateTimeFormat("id-ID", {
    day: "numeric",
    month: "short",
    year: "numeric"
  }).format(date);
}

/** Returns today's date in YYYY-MM-DD format */
function getTodayISO() {
  const today = new Date();
  return today.toISOString().split("T")[0];
}

/** Generates HTML badge string based on Exposure Tier */
function getTierBadgeHTML(tier) {
  switch (tier) {
    case "Tier 1":
      return `<span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-[11px] font-bold bg-rose-50 text-rose-700 border border-rose-200">
                <span class="w-1.5 h-1.5 rounded-full bg-rose-500 pulse-glow"></span>
                Tier 1 (Tinggi)
              </span>`;
    case "Tier 2":
      return `<span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-[11px] font-bold bg-amber-50 text-amber-700 border border-amber-200">
                <span class="w-1.5 h-1.5 rounded-full bg-amber-500"></span>
                Tier 2 (Sedang)
              </span>`;
    case "Tier 3":
      return `<span class="inline-flex items-center gap-1.5 px-2.5 py-1 rounded-full text-[11px] font-bold bg-emerald-50 text-emerald-700 border border-emerald-200">
                <span class="w-1.5 h-1.5 rounded-full bg-emerald-500"></span>
                Tier 3 (Rendah)
              </span>`;
    default:
      return `<span class="px-2 py-1 rounded text-xs bg-slate-100 text-slate-700">${tier}</span>`;
  }
}

/** Normalizes any commitment status to one of four standard statuses */
function normalizeStatus(status) {
  if (!status) return "Belum Dilaksanakan";
  const s = String(status).trim().toLowerCase();
  if (s.includes("progress") || s.includes("implementasi") || s.includes("ptp") || s.includes("restrukturisasi") || s.includes("call back") || s === "promise_to_pay") {
    return "On Progress/Implementasi";
  }
  if (s.includes("belum") || s.includes("pending") || s.includes("tidak ada respon")) {
    return "Belum Dilaksanakan";
  }
  if (s.includes("telah") || s.includes("dilaksanakan") || s.includes("kept") || s.includes("setor")) {
    return "Telah Dilaksanakan";
  }
  if (s.includes("melewati") || s.includes("refusal") || s.includes("penolakan") || s.includes("dispute")) {
    return "Melewati Komitmen";
  }
  return status;
}

/** Generates HTML status badge for commitment status */
function getCommitmentStatusBadgeHTML(status) {
  if (!status) return "";
  const norm = normalizeStatus(status);
  if (norm === "On Progress/Implementasi") {
    return `<span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-amber-50 text-amber-700 border border-amber-200 inline-flex items-center gap-1"><i data-lucide="clock" class="w-3 h-3"></i>${norm}</span>`;
  } else if (norm === "Belum Dilaksanakan") {
    return `<span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-slate-100 text-slate-700 border border-slate-300 inline-flex items-center gap-1"><i data-lucide="circle-dashed" class="w-3 h-3"></i>${norm}</span>`;
  } else if (norm === "Telah Dilaksanakan") {
    return `<span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-emerald-50 text-emerald-700 border border-emerald-200 inline-flex items-center gap-1"><i data-lucide="check-circle-2" class="w-3 h-3"></i>${norm}</span>`;
  } else if (norm === "Melewati Komitmen") {
    return `<span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-rose-50 text-rose-700 border border-rose-200 inline-flex items-center gap-1"><i data-lucide="alert-circle" class="w-3 h-3"></i>${norm}</span>`;
  }
  return `<span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-slate-100 text-slate-700 border border-slate-200">${status}</span>`;
}

// ==========================================
// 4. TOAST NOTIFICATIONS
// ==========================================
function showToast(title, message, type = "info") {
  const container = document.getElementById("toastContainer");
  if (!container) return;

  const toast = document.createElement("div");
  toast.className = "pointer-events-auto bg-white border border-slate-200 text-slate-900 p-4 rounded-xl shadow-xl flex items-start gap-3 animate-slide-up transition-all";

  let iconHTML = `<i data-lucide="info" class="w-5 h-5 text-indigo-600"></i>`;
  if (type === "success") {
    iconHTML = `<i data-lucide="check-circle" class="w-5 h-5 text-emerald-600"></i>`;
  } else if (type === "warning") {
    iconHTML = `<i data-lucide="alert-triangle" class="w-5 h-5 text-amber-600"></i>`;
  } else if (type === "error") {
    iconHTML = `<i data-lucide="alert-circle" class="w-5 h-5 text-rose-600"></i>`;
  }

  toast.innerHTML = `
    <div class="mt-0.5">${iconHTML}</div>
    <div class="flex-1 text-xs">
      <h4 class="font-bold text-slate-900 text-sm">${title}</h4>
      <p class="text-slate-600 mt-0.5 leading-snug">${message}</p>
    </div>
    <button class="text-slate-400 hover:text-slate-700 transition-colors" onclick="this.parentElement.remove()">
      <i data-lucide="x" class="w-4 h-4"></i>
    </button>
  `;

  container.appendChild(toast);
  if (window.lucide) lucide.createIcons();

  setTimeout(() => {
    toast.style.opacity = "0";
    toast.style.transform = "translateY(10px)";
    setTimeout(() => toast.remove(), 300);
  }, 4000);
}

// ==========================================
// 5. DATA FILTERING & SORTING LOGIC
// ==========================================

function populatePengelolaFilterOptions() {
  const select = document.getElementById("pengelolaFilter");
  if (!select) return;

  const currentSelection = state.selectedPengelola || "ALL";
  const officers = (state.officersList && state.officersList.length > 0) ? state.officersList : DEFAULT_OFFICERS;

  const optionsMap = new Map();

  officers.forEach(o => {
    if (o.full_name) {
      const displayLabel = o.officer_id ? `${o.full_name} (${o.officer_id})` : o.full_name;
      optionsMap.set(o.full_name, displayLabel);
    }
  });

  if (state.debtors && state.debtors.length > 0) {
    state.debtors.forEach(d => {
      if (d.officerName && !optionsMap.has(d.officerName)) {
        optionsMap.set(d.officerName, d.pengelola || d.officerName);
      } else if (d.pengelola && !Array.from(optionsMap.values()).includes(d.pengelola)) {
        const nameOnly = d.pengelola.split(" (")[0];
        if (!optionsMap.has(nameOnly)) {
          optionsMap.set(nameOnly, d.pengelola);
        }
      }
    });
  }

  select.innerHTML = `<option value="ALL">Semua Account Officer</option>`;

  const sortedKeys = Array.from(optionsMap.keys()).sort();
  sortedKeys.forEach(name => {
    const label = optionsMap.get(name);
    const opt = document.createElement("option");
    opt.value = name;
    opt.textContent = label;
    if (name === currentSelection || label === currentSelection) {
      opt.selected = true;
    }
    select.appendChild(opt);
  });

  if (currentSelection && currentSelection !== "ALL") {
    select.value = currentSelection;
    if (!select.value && optionsMap.has(currentSelection)) {
      select.value = currentSelection;
    }
  } else {
    select.value = "ALL";
  }
}

function applyFiltersAndSort() {
  const query = state.searchQuery.trim().toLowerCase();

  // Filter dataset
  state.filteredDebtors = state.debtors.filter(debtor => {
    // Search query match
    const matchSearch = !query ||
      debtor.namaDebitur.toLowerCase().includes(query) ||
      debtor.noRekening.toLowerCase().includes(query);

    // Tier match
    const matchTier = state.selectedTier === "ALL" || debtor.tierEksposur === state.selectedTier;

    // Pengelola match
    const matchPengelola = state.selectedPengelola === "ALL" ||
      debtor.pengelola === state.selectedPengelola ||
      debtor.officerName === state.selectedPengelola;

    // Produk match
    const matchProduk = state.selectedProduk === "ALL" || debtor.produk === state.selectedProduk;

    return matchSearch && matchTier && matchPengelola && matchProduk;
  });

  // Sort dataset
  state.filteredDebtors.sort((a, b) => {
    let valA = a[state.sortColumn];
    let valB = b[state.sortColumn];

    if (typeof valA === "string") valA = valA.toLowerCase();
    if (typeof valB === "string") valB = valB.toLowerCase();

    if (valA < valB) return state.sortDirection === "asc" ? -1 : 1;
    if (valA > valB) return state.sortDirection === "asc" ? 1 : -1;
    return 0;
  });

  // Reset to first page if current page exceeds max pages
  const totalPages = state.serverPagination
    ? state.serverPagination.totalPages
    : (Math.ceil(state.filteredDebtors.length / state.itemsPerPage) || 1);
  if (state.currentPage > totalPages) {
    state.currentPage = 1;
  }
}

const DEFAULT_OFFICERS = [
  { officer_id: "OFF-001", full_name: "Budi Santoso", pairing_team: "Desk 01" },
  { officer_id: "OFF-002", full_name: "Siti Rahma", pairing_team: "Desk 02" },
  { officer_id: "OFF-003", full_name: "Ahmad Dahlan", pairing_team: "Desk 03" },
  { officer_id: "OFF-004", full_name: "Dian Sastro", pairing_team: "Desk 04" }
];

const OFFICERS_DICTIONARY = {
  "OFF-001": { name: "Budi Santoso", pair: "Desk 01" },
  "OFF-002": { name: "Siti Rahma", pair: "Desk 02" },
  "OFF-003": { name: "Ahmad Dahlan", pair: "Desk 03" },
  "OFF-004": { name: "Dian Sastro", pair: "Desk 04" }
};

/** Fetches officers list from GET http://localhost:8080/api/v1/officers */
async function fetchOfficersFromApi() {
  try {
    const res = await fetch("http://localhost:8080/api/v1/officers");
    if (!res.ok) throw new Error(`HTTP Error ${res.status}`);
    const json = await res.json();

    if (json.success && Array.isArray(json.data) && json.data.length > 0) {
      state.officersList = json.data;
      json.data.forEach(o => {
        if (o.officer_id) {
          OFFICERS_DICTIONARY[o.officer_id] = {
            name: o.full_name,
            pair: o.pairing_team || "Desk"
          };
        }
      });
    }
  } catch (err) {
    console.warn("[FE API Sync] GET /api/v1/officers offline or failed, using local dictionary:", err.message);
  }
  populateOfficerDropdowns();
  populatePengelolaFilterOptions();
}


/** Populates select dropdown options for Account Officer and Pairing Officer */
function populateOfficerDropdowns() {
  const officers = (state.officersList && state.officersList.length > 0) ? state.officersList : DEFAULT_OFFICERS;

  // Update Commitment Modal Selects
  const officerSelect = document.getElementById("inputOfficerID");
  const pairSelect = document.getElementById("inputOfficerPairID");

  // Create Commitment Modal Selects
  const createOfficerSelect = document.getElementById("createOfficerID");
  const createPairSelect = document.getElementById("createOfficerPairID");

  if (officerSelect) {
    const currentVal = officerSelect.value;
    officerSelect.innerHTML = officers.map(o =>
      `<option value="${o.officer_id}">${o.officer_id} - ${o.full_name}</option>`
    ).join("");
    if (currentVal && officers.some(o => o.officer_id === currentVal)) {
      officerSelect.value = currentVal;
    }
  }

  if (pairSelect) {
    const currentVal = pairSelect.value;
    pairSelect.innerHTML = officers.map(o =>
      `<option value="${o.officer_id}">${o.officer_id} - ${o.full_name} (${o.pairing_team || 'Desk'})</option>`
    ).join("");
    if (currentVal && officers.some(o => o.officer_id === currentVal)) {
      pairSelect.value = currentVal;
    }
  }

  if (createOfficerSelect) {
    const currentVal = createOfficerSelect.value;
    createOfficerSelect.innerHTML = officers.map(o =>
      `<option value="${o.officer_id}">${o.officer_id} - ${o.full_name}</option>`
    ).join("");
    if (currentVal && officers.some(o => o.officer_id === currentVal)) {
      createOfficerSelect.value = currentVal;
    }
  }

  if (createPairSelect) {
    const currentVal = createPairSelect.value;
    createPairSelect.innerHTML = officers.map(o =>
      `<option value="${o.officer_id}">${o.officer_id} - ${o.full_name} (${o.pairing_team || 'Desk'})</option>`
    ).join("");
    if (currentVal && officers.some(o => o.officer_id === currentVal)) {
      createPairSelect.value = currentVal;
    }
  }
}

function resolveOfficerName(val) {
  if (!val) return "Budi Santoso";
  if (OFFICERS_DICTIONARY[val]) return OFFICERS_DICTIONARY[val].name;
  return val;
}

function resolvePairingTeam(val, officerId) {
  if (OFFICERS_DICTIONARY[val]) return OFFICERS_DICTIONARY[val].pair;
  if (!val || val === "" || val === officerId) {
    if (officerId && OFFICERS_DICTIONARY[officerId]) {
      return OFFICERS_DICTIONARY[officerId].pair;
    }
    return "Desk 01";
  }
  return val;
}

/** Fetches paginated commitment listing data from backend GET /api/v1/commitments */
async function fetchCommitmentsFromApi() {
  const page = state.currentPage || 1;
  const limit = state.itemsPerPage || 10;
  const search = encodeURIComponent(state.searchQuery || "");
  const tier = encodeURIComponent(state.selectedTier || "ALL");
  const pengelola = encodeURIComponent(state.selectedPengelola || "ALL");
  const produk = encodeURIComponent(state.selectedProduk || "ALL");
  const sortBy = encodeURIComponent(state.sortColumn || state.sortField || "");
  const sortDir = encodeURIComponent(state.sortDirection || "asc");

  let url = `http://localhost:8080/api/v1/commitments?page=${page}&limit=${limit}`;
  if (search) url += `&search=${search}`;
  if (tier && tier !== "ALL") url += `&tier=${tier}`;
  if (pengelola && pengelola !== "ALL") url += `&pengelola=${pengelola}`;
  if (produk && produk !== "ALL") url += `&produk=${produk}`;
  if (sortBy) url += `&sort_by=${sortBy}&sort_dir=${sortDir}`;

  try {
    const res = await fetch(url);
    if (!res.ok) throw new Error(`HTTP Error ${res.status}`);
    const json = await res.json();

    if (json.success && json.data) {
      const commitments = json.data.data || [];
      const pagination = json.data.pagination || {};

      state.serverPagination = {
        currentPage: pagination.currentPage || page,
        itemsPerPage: pagination.itemsPerPage || limit,
        totalItems: pagination.totalItems || commitments.length,
        totalPages: pagination.totalPages || 1
      };

      if (commitments.length > 0) {
        state.debtors = commitments.map(c => {
          const existing = INITIAL_DEBTORS.find(d => d.noRekening === c.account_no);

          const officerId = c.officer_id || (existing ? existing.officerId : "");
          let officerName = c.officer_name || resolveOfficerName(officerId);
          if ((!officerName || officerName.startsWith("OFF-")) && existing) {
            officerName = existing.pengelola;
          }
          const officerDisplay = officerId && officerName && !officerName.includes(officerId)
            ? `${officerName} (${officerId})`
            : (officerName || officerId || "Budi Santoso");

          const pairId = c.officer_pair_id || (existing ? existing.pairing : "");
          let pairName = c.officer_pair_name || resolvePairingTeam(pairId, officerId);
          if ((!pairName || pairName.startsWith("OFF-")) && existing) {
            pairName = existing.pairing;
          }
          const pairDisplay = pairId && pairName && !pairName.includes(pairId)
            ? `${pairName} (${pairId})`
            : (pairName || pairId || "Desk 01");

          return {
            id: c.id || (existing ? existing.id : "CMT-" + c.account_no),
            noRekening: c.account_no,
            namaDebitur: existing ? existing.namaDebitur : `Debitur (${c.account_no})`,
            produk: c.product || (existing ? existing.produk : "KMK"),
            limit: parseMoney(c.credit_limit) > 0 ? parseMoney(c.credit_limit) : (existing ? existing.limit : (parseMoney(c.nominal) > 0 ? parseMoney(c.nominal) * 1.2 : 1000000000)),
            bakiDebet: parseMoney(c.outstanding_balance) > 0 ? parseMoney(c.outstanding_balance) : (parseMoney(c.nominal) > 0 ? parseMoney(c.nominal) : (existing ? existing.bakiDebet : 500000000)),
            nominalKomitmen: parseMoney(c.nominal) > 0 ? parseMoney(c.nominal) : (parseMoney(c.outstanding_balance) > 0 ? parseMoney(c.outstanding_balance) : (existing ? existing.nominalKomitmen : 500000000)),
            tierEksposur: c.exposure_tier || (existing ? existing.tierEksposur : "Tier 2"),
            officerId: officerId,
            officerName: officerName,
            officerPairId: pairId,
            officerPairName: pairName,
            pengelola: officerDisplay,
            pairing: pairDisplay,
            tglKomitmen: c.commitment_date || "",
            statusKomitmen: normalizeStatus(c.status || c.reason || "Belum Dilaksanakan"),
            keterangan: c.remarks || c.reason || "",
            phone: existing ? existing.phone : "081298765432"
          };
        });
      } else {
        state.debtors = [];
      }
      return true;
    }

  } catch (err) {
    console.warn("[FE API Sync] GET /api/v1/commitments offline or failed, using local dataset:", err.message);
    state.serverPagination = null;
    return false;
  }
}

async function loadAndRenderData() {
  await fetchOfficersFromApi();
  await fetchCommitmentsFromApi();
  populatePengelolaFilterOptions();
  applyFiltersAndSort();
  renderMetrics();
  renderTable();
}


// ==========================================
// 6. UI RENDER FUNCTIONS
// ==========================================

/** Render Metric Summary Cards at the Top */
function renderMetrics() {
  const totalDebtors = state.debtors.length;
  const totalBakiDebet = state.debtors.reduce((acc, curr) => acc + curr.bakiDebet, 0);

  const tier1Count = state.debtors.filter(d => d.tierEksposur === "Tier 1").length;
  const tier1Ratio = totalDebtors > 0 ? ((tier1Count / totalDebtors) * 100).toFixed(1) : "0";

  const todayStr = getTodayISO();
  const commitmentsToday = state.debtors.filter(d => d.tglKomitmen === todayStr).length;

  document.getElementById("metricTotalDebtors").textContent = totalDebtors;
  document.getElementById("metricTotalBakiDebet").textContent = formatIDR(totalBakiDebet);
  document.getElementById("metricTier1Count").textContent = tier1Count;
  document.getElementById("metricTier1Ratio").textContent = `(${tier1Ratio}%)`;
  document.getElementById("metricCommitmentsToday").textContent = commitmentsToday;
}

function getCommitmentExpiryStatus(tglKomitmen) {
  if (!tglKomitmen) return { status: "normal", diffDays: null };

  const todayStr = getTodayISO();
  const today = new Date(todayStr);
  today.setHours(0, 0, 0, 0);

  const commitDate = new Date(tglKomitmen);
  commitDate.setHours(0, 0, 0, 0);

  if (isNaN(commitDate.getTime())) return { status: "normal", diffDays: null };

  const diffTime = commitDate.getTime() - today.getTime();
  const diffDays = Math.ceil(diffTime / (1000 * 3600 * 24));

  if (diffDays <= 0) {
    return { status: "expired", diffDays };
  } else if (diffDays <= 7) {
    return { status: "warning", diffDays };
  }
  return { status: "normal", diffDays };
}

/** Render Main Data Table Rows & Pagination */
function renderTable() {
  const tbody = document.getElementById("debtorTableBody");
  const emptyState = document.getElementById("emptyState");
  const visibleCountEl = document.getElementById("visibleRowCount");
  const totalCountEl = document.getElementById("totalRowCount");
  const activeBadge = document.getElementById("activeFilterBadge");

  if (!tbody) return;

  // Update counts
  const totalCount = state.serverPagination ? state.serverPagination.totalItems : state.debtors.length;
  visibleCountEl.textContent = state.filteredDebtors.length;
  totalCountEl.textContent = totalCount;

  const isFilterActive = state.searchQuery || state.selectedTier !== "ALL" || state.selectedPengelola !== "ALL" || state.selectedProduk !== "ALL";
  if (isFilterActive) {
    activeBadge.classList.remove("hidden");
  } else {
    activeBadge.classList.add("hidden");
  }

  // Handle empty state
  if (state.filteredDebtors.length === 0) {
    tbody.innerHTML = "";
    emptyState.classList.remove("hidden");
    renderPagination(0);
    return;
  }

  emptyState.classList.add("hidden");

  // Calculate Pagination Slices
  let paginatedItems = state.filteredDebtors;
  if (!state.serverPagination) {
    const startIndex = (state.currentPage - 1) * state.itemsPerPage;
    const endIndex = startIndex + state.itemsPerPage;
    paginatedItems = state.filteredDebtors.slice(startIndex, endIndex);
  }

  // Build HTML string
  let html = "";
  paginatedItems.forEach(debtor => {
    const expiry = getCommitmentExpiryStatus(debtor.tglKomitmen);

    let rowClass = "hover:bg-slate-50/80 transition-colors group border-l-4 border-l-transparent";
    let stickyBgClass = "bg-white group-hover:bg-slate-50/80";
    let tglDisplay = "";

    if (expiry.status === "expired") {
      rowClass = "bg-rose-50/80 hover:bg-rose-100/90 transition-colors group border-l-4 border-l-rose-500 font-medium";
      stickyBgClass = "bg-rose-50/90 group-hover:bg-rose-100/90";
      const isToday = debtor.tglKomitmen === getTodayISO();
      tglDisplay = `
        <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg text-[11px] font-extrabold bg-rose-100 text-rose-900 border border-rose-300 shadow-2xs">
          <i data-lucide="alert-triangle" class="w-3.5 h-3.5 text-rose-600"></i>
          <span>${isToday ? 'Hari Ini (Expire)' : `Expired (${formatDate(debtor.tglKomitmen)})`}</span>
        </span>
      `;
    } else if (expiry.status === "warning") {
      rowClass = "bg-amber-50/80 hover:bg-amber-100/90 transition-colors group border-l-4 border-l-amber-500 font-medium";
      stickyBgClass = "bg-amber-50/90 group-hover:bg-amber-100/90";
      tglDisplay = `
        <span class="inline-flex items-center gap-1 px-2.5 py-1 rounded-lg text-[11px] font-extrabold bg-amber-100 text-amber-900 border border-amber-300 shadow-2xs">
          <i data-lucide="clock" class="w-3.5 h-3.5 text-amber-600"></i>
          <span>${expiry.diffDays} Hari Lagi (${formatDate(debtor.tglKomitmen)})</span>
        </span>
      `;
    } else {
      tglDisplay = `<div class="text-xs font-semibold text-slate-800">${formatDate(debtor.tglKomitmen)}</div>`;
    }

    html += `
      <tr class="${rowClass}">
        <!-- No Rekening -->
        <td class="py-3.5 px-4 font-mono font-bold text-slate-900 whitespace-nowrap">
          <div class="flex items-center gap-1.5">
            <span>${debtor.noRekening}</span>
            <button onclick="copyText('${debtor.noRekening}', 'No Rekening ${debtor.noRekening}')" class="text-amber-600 hover:text-amber-700 opacity-0 group-hover:opacity-100 transition-opacity" title="Salin Rekening">
              <i data-lucide="copy" class="w-3.5 h-3.5"></i>
            </button>
          </div>
        </td>

        <!-- Nama Debitur -->
        <td class="py-3.5 px-4">
          <div class="font-bold text-slate-900 leading-tight">${debtor.namaDebitur}</div>
          <div class="text-[11px] text-slate-500 mt-0.5">ID: ${debtor.id} | Telp: ${debtor.phone}</div>
        </td>

        <!-- Produk -->
        <td class="py-3.5 px-4">
          <span class="px-2.5 py-0.5 rounded-lg font-extrabold text-[11px] bg-slate-100 text-slate-700 border border-slate-200">${debtor.produk}</span>
        </td>

        <!-- Limit -->
        <td class="py-3.5 px-4 text-right font-medium text-slate-600 whitespace-nowrap">
          ${formatIDR(debtor.limit)}
        </td>

        <!-- Baki Debet -->
        <td class="py-3.5 px-4 text-right font-extrabold text-emerald-600 whitespace-nowrap">
          ${formatIDR(debtor.bakiDebet)}
        </td>

        <!-- Tier Eksposur -->
        <td class="py-3.5 px-4 text-center whitespace-nowrap">
          ${getTierBadgeHTML(debtor.tierEksposur)}
        </td>

        <!-- Pengelola -->
        <td class="py-3.5 px-4 whitespace-nowrap">
          <div class="font-semibold text-slate-800">${debtor.pengelola}</div>
        </td>

        <!-- Pairing -->
        <td class="py-3.5 px-4 text-center whitespace-nowrap">
          <span class="px-2.5 py-0.5 rounded-full text-[10px] font-bold bg-slate-100 text-slate-600 border border-slate-200">${debtor.pairing}</span>
        </td>

        <!-- Tanggal Komitmen -->
        <td class="py-3.5 px-4 text-center whitespace-nowrap">
          <div>${tglDisplay}</div>
        </td>

        <!-- Status Komitmen -->
        <td class="py-3.5 px-4 text-center whitespace-nowrap">
          ${getCommitmentStatusBadgeHTML(debtor.statusKomitmen)}
        </td>

        <!-- Keterangan -->
        <td class="py-3.5 px-4 max-w-xs">
          <p class="text-slate-600 text-xs line-clamp-2 leading-tight" title="${debtor.keterangan}">${debtor.keterangan || "-"}</p>
        </td>

        <!-- Actions -->
        <td class="py-3.5 px-4 text-center whitespace-nowrap sticky right-0 ${stickyBgClass} border-l border-slate-100">

          <div class="flex items-center justify-center gap-1.5">
            <button 
              onclick="openWordingModal('${debtor.id}')"
              class="px-2.5 py-1.5 bg-white hover:bg-slate-50 text-slate-800 border border-slate-200 rounded-lg text-[11px] font-bold transition-all flex items-center gap-1 shadow-xs"
              title="Generate Collection Wording"
            >
              <i data-lucide="sparkles" class="w-3.5 h-3.5 text-amber-600"></i>
              <span>Wording</span>
            </button>
            <button 
              onclick="openCommitmentModal('${debtor.id}')"
              class="px-2.5 py-1.5 bg-amber-500 hover:bg-amber-400 text-slate-950 rounded-lg text-[11px] font-extrabold transition-all flex items-center gap-1 shadow-xs"
              title="Update Komitmen Debitur"
            >
              <i data-lucide="edit-3" class="w-3.5 h-3.5"></i>
              <span>Update</span>
            </button>
          </div>
        </td>
      </tr>
    `;
  });

  tbody.innerHTML = html;
  renderPagination(state.filteredDebtors.length);
  if (window.lucide) lucide.createIcons();
}

/** Render Pagination Buttons and Summary */
function renderPagination(totalItems) {
  const summaryEl = document.getElementById("paginationSummary");
  const buttonsEl = document.getElementById("paginationButtons");

  if (!summaryEl || !buttonsEl) return;

  const totalItemsCount = state.serverPagination ? state.serverPagination.totalItems : totalItems;
  const totalPages = state.serverPagination ? state.serverPagination.totalPages : (Math.ceil(totalItemsCount / state.itemsPerPage) || 1);
  summaryEl.textContent = `Halaman ${state.currentPage} dari ${totalPages}`;

  let btnHtml = "";

  // Prev Button
  btnHtml += `
    <button 
      onclick="changePage(${state.currentPage - 1})"
      ${state.currentPage === 1 ? "disabled" : ""}
      class="p-1.5 rounded-lg border border-slate-200 bg-white text-slate-700 hover:bg-slate-100 disabled:opacity-40 disabled:cursor-not-allowed transition-colors"
      title="Halaman Sebelumnya"
    >
      <i data-lucide="chevron-left" class="w-4 h-4 text-slate-700"></i>
    </button>
  `;

  // Page Numbers
  for (let i = 1; i <= totalPages; i++) {
    if (
      i === 1 ||
      i === totalPages ||
      (i >= state.currentPage - 1 && i <= state.currentPage + 1)
    ) {
      const isActive = i === state.currentPage;
      btnHtml += `
        <button 
          onclick="changePage(${i})"
          class="px-3 py-1 rounded-lg text-xs font-bold transition-all ${isActive
          ? "bg-amber-500 text-slate-950 font-extrabold shadow-xs"
          : "bg-white border border-slate-200 text-slate-700 hover:bg-slate-100"
        }"
        >
          ${i}
        </button>
      `;
    } else if (
      (i === 2 && state.currentPage > 3) ||
      (i === totalPages - 1 && state.currentPage < totalPages - 2)
    ) {
      btnHtml += `<span class="px-1 text-slate-400">...</span>`;
    }
  }

  // Next Button
  btnHtml += `
    <button 
      onclick="changePage(${state.currentPage + 1})"
      ${state.currentPage === totalPages ? "disabled" : ""}
      class="p-1.5 rounded-lg border border-slate-200 bg-white text-slate-700 hover:bg-slate-100 disabled:opacity-40 disabled:cursor-not-allowed transition-colors"
      title="Halaman Berikutnya"
    >
      <i data-lucide="chevron-right" class="w-4 h-4 text-slate-700"></i>
    </button>
  `;

  buttonsEl.innerHTML = btnHtml;
  if (window.lucide) lucide.createIcons();
}

async function changePage(newPage) {
  const totalPages = state.serverPagination
    ? state.serverPagination.totalPages
    : (Math.ceil(state.filteredDebtors.length / state.itemsPerPage) || 1);

  if (newPage < 1 || newPage > totalPages) return;
  state.currentPage = newPage;
  await loadAndRenderData();
}
window.changePage = changePage;

// ==========================================
// 7. MODAL 1: WORDING GENERATOR LOGIC
// ==========================================

function openWordingModal(debtorId) {
  const debtor = state.debtors.find(d => d.id === debtorId);
  if (!debtor) return;

  state.activeWordingDebtor = debtor;
  state.activeWordingChannel = "whatsapp";
  state.activeWordingTone = debtor.tierEksposur === "Tier 1" ? "urgent" : "firm";

  // Update modal header & snapshot
  document.getElementById("wordingDebtorName").textContent = debtor.namaDebitur;
  document.getElementById("wordingAccNo").textContent = debtor.noRekening;
  document.getElementById("wordingBakiDebet").textContent = formatIDR(debtor.bakiDebet);
  document.getElementById("wordingTierBadge").innerHTML = getTierBadgeHTML(debtor.tierEksposur);
  document.getElementById("wordingTglKomitmen").textContent = formatDate(debtor.tglKomitmen);

  document.getElementById("wordingToneSelect").value = state.activeWordingTone;

  generateAndSetWordingText();

  // Reset Copy button state
  resetCopyBtnState();

  // Show modal
  const modal = document.getElementById("wordingModal");
  modal.classList.remove("hidden");
  if (window.lucide) lucide.createIcons();
}

function closeWordingModal() {
  const modal = document.getElementById("wordingModal");
  modal.classList.add("hidden");
  state.activeWordingDebtor = null;
}

function generateAndSetWordingText() {
  const debtor = state.activeWordingDebtor;
  if (!debtor) return;

  const tone = state.activeWordingTone;
  const channel = state.activeWordingChannel;
  const textarea = document.getElementById("wordingTextarea");

  let text = "";

  if (channel === "whatsapp") {
    if (tone === "soft") {
      text = `Yth. Bapak/Ibu ${debtor.namaDebitur},

Selamat pagi/siang. Kami dari Tim Collection PT Bank Utama Mandiri (BUMA) menginfokan bahwa kewajiban angsuran kredit Anda (${debtor.produk}) No. Rekening: *${debtor.noRekening}* dengan sisa baki debet *${formatIDR(debtor.bakiDebet)}* akan memasuki tanggal komitmen pada *${formatDate(debtor.tglKomitmen)}*.

Mohon dapat melakukan penyetoran sebelum pukul 17:00 WIB untuk menjaga kualitas kredit Anda tetap lancar.

Jika telah melakukan pembayaran, abaikan pesan ini. Terima kasih.
Petugas Pengelola: ${debtor.pengelola} (BUMA Collection Unit)`;
    } else if (tone === "firm") {
      text = `PEMBERITAHUAN PENAGIHAN KREDIT - BUMA

Kepada Yth.
*${debtor.namaDebitur}*
No. Rekening: *${debtor.noRekening}*

Diberitahukan bahwa tagihan kewajiban fasilitas kredit *${debtor.produk}* Anda sebesar *${formatIDR(debtor.bakiDebet)}* telah melewati jatuh tempo. Sesuai catatan komitmen, jadwal pembayaran jatuh pada hari *${formatDate(debtor.tglKomitmen)}*.

Mohon SEGERA melakukan pembayaran melalui rekening penampungan BUMA atau konfirmasi bukti transfer hari ini kepada Account Officer Anda:
*${debtor.pengelola}* (${debtor.pairing})

Hubungi Call Center BUMA jika membutuhkan bantuan kendala transaksi.`;
    } else { // urgent / legal
      text = `*SURAT PERINGATAN / SOMASI PRA-HUKUM*
PT BANK UTAMA MANDIRI (BUMA)

Kepada Yth. Debitur: *${debtor.namaDebitur}*
Fasilitas Kredit: *${debtor.produk}*
No. Rekening: *${debtor.noRekening}*
Total Baki Debet: *${formatIDR(debtor.bakiDebet)}*

Berdasarkan evaluasi risiko (${debtor.tierEksposur}), fasilitas kredit Anda telah dikategorikan DALAM PENAWASAN KHUSUS. Peringatan penagihan resmi telah diproses.

Anda diwajibkan melakukan penyelesaian kewajiban / penyetoran komitmen pada *${formatDate(debtor.tglKomitmen)}*. Kelalaian pembayaran akan mengakibatkan pelaporan kolektibilitas pada SLIK OJK dan tindakan hukum penanganan agunan.

Segera konfirmasi ke Pengelola: *${debtor.pengelola}*.`;
    }
  } else if (channel === "email") {
    text = `Subjek: [BUMA COLLECTION] Pemberitahuan Kewajiban Kredit No. Rek ${debtor.noRekening} - Yth. ${debtor.namaDebitur}

Kepada Yth.
Management / Bp/Ibu ${debtor.namaDebitur}

Dengan hormat,

Sehubungan dengan fasilitas kredit ${debtor.produk} atas nama ${debtor.namaDebitur} dengan nomor rekening ${debtor.noRekening}, melalui surat ini kami menyampaikan rincian posisi pinjaman Anda per tanggal ${formatDate(getTodayISO())}:

- Nama Debitur: ${debtor.namaDebitur}
- No. Rekening: ${debtor.noRekening}
- Jenis Fasilitas: ${debtor.produk}
- Baki Debet Pinjaman: ${formatIDR(debtor.bakiDebet)}
- Tanggal Komitmen Bayar: ${formatDate(debtor.tglKomitmen)}
- Tim Pengelola: ${debtor.pengelola} (${debtor.pairing})

Dimohon untuk dapat melakukan penyetoran dana ke rekening efektif tepat pada tanggal komitmen yang disepakati.

Demikian pemberitahuan ini kami sampaikan. Atas perhatian dan kerja samanya, kami ucapkan terima kasih.

Hormat kami,
PT Bank Utama Mandiri (BUMA)
Divisi Special Asset Management & Collection`;
  } else { // SMS / Call script
    text = `[SCRIPT TELEPON / SMS BUMA]
Halo Bp/Ibu ${debtor.namaDebitur}, saya ${debtor.pengelola} dari BUMA. Mengonfirmasi janji bayar fasilitas ${debtor.produk} No. Rek ${debtor.noRekening} (Baki debet ${formatIDR(debtor.bakiDebet)}) pada tanggal ${formatDate(debtor.tglKomitmen)}. Mohon dipastikan dana tersedia sebelum pukul 15.00 WIB. Terima kasih.`;
  }

  textarea.value = text;
}

function resetCopyBtnState() {
  const btnText = document.getElementById("copyBtnText");
  const icon = document.getElementById("copyIcon");
  btnText.textContent = "Salin Wording";
  icon.setAttribute("data-lucide", "copy");
  if (window.lucide) lucide.createIcons();
}

/** Copy wording text to clipboard with feedback */
function copyWordingToClipboard() {
  const textarea = document.getElementById("wordingTextarea");
  if (!textarea) return;

  copyText(textarea.value, "Wording penagihan berhasil disalin ke clipboard!");

  const btnText = document.getElementById("copyBtnText");
  const icon = document.getElementById("copyIcon");
  btnText.textContent = "Tersalin!";
  icon.setAttribute("data-lucide", "check");
  if (window.lucide) lucide.createIcons();

  setTimeout(resetCopyBtnState, 3000);
}

/** Open WhatsApp Web directly with pre-filled message */
function openWhatsAppWeb() {
  const debtor = state.activeWordingDebtor;
  const textarea = document.getElementById("wordingTextarea");
  if (!debtor || !textarea) return;

  let rawPhone = debtor.phone || "08123456789";
  if (rawPhone.startsWith("0")) {
    rawPhone = "62" + rawPhone.slice(1);
  }

  const encodedText = encodeURIComponent(textarea.value);
  const waUrl = `https://wa.me/${rawPhone}?text=${encodedText}`;

  window.open(waUrl, "_blank");
  showToast("Membuka WhatsApp", `Menghubungkan ke ${debtor.namaDebitur} (${rawPhone})`, "info");
}

// Global copy helper
function copyText(text, successMessage) {
  if (navigator.clipboard && window.isSecureContext) {
    navigator.clipboard.writeText(text).then(() => {
      showToast("Berhasil Disalin", successMessage, "success");
    }).catch(err => {
      fallbackCopyText(text, successMessage);
    });
  } else {
    fallbackCopyText(text, successMessage);
  }
}

function fallbackCopyText(text, successMessage) {
  const input = document.createElement("textarea");
  input.value = text;
  document.body.appendChild(input);
  input.select();
  document.execCommand("copy");
  document.body.removeChild(input);
  showToast("Berhasil Disalin", successMessage, "success");
}

// ==========================================
// 8. MODAL 2: COMMITMENT UPDATE LOGIC
// ==========================================

async function openCommitmentModal(debtorId) {
  const debtor = state.debtors.find(d => d.id === debtorId);
  if (!debtor) return;

  state.activeCommitmentDebtor = debtor;

  // Re-fetch officers from API to ensure options are up-to-date
  await fetchOfficersFromApi();

  // Update header context
  document.getElementById("commitmentDebtorName").textContent = debtor.namaDebitur;
  document.getElementById("commitmentAccNo").textContent = debtor.noRekening;
  document.getElementById("commitmentBakiDebet").textContent = formatIDR(debtor.bakiDebet);
  document.getElementById("commitmentDebtorId").value = debtor.id;

  // Pre-fill inputs
  document.getElementById("inputTglKomitmen").value = debtor.tglKomitmen || getTodayISO();
  document.getElementById("inputStatusKomitmen").value = normalizeStatus(debtor.statusKomitmen) || "On Progress/Implementasi";
  // Pre-fill inputs with formatted display strings
  const initialNominal = debtor.nominalKomitmen || debtor.bakiDebet || 0;
  const initialLimit = debtor.limit || 0;
  const initialBakiDebet = debtor.bakiDebet || 0;

  document.getElementById("inputNominal").value = initialNominal ? initialNominal.toLocaleString("id-ID") : "";
  document.getElementById("inputKeterangan").value = debtor.keterangan || "";

  if (document.getElementById("inputOfficerID")) {
    document.getElementById("inputOfficerID").value = debtor.officerId || "OFF-001";
  }
  if (document.getElementById("inputOfficerPairID")) {
    document.getElementById("inputOfficerPairID").value = debtor.officerPairId || "OFF-002";
  }
  if (document.getElementById("inputLimit")) {
    document.getElementById("inputLimit").value = initialLimit ? initialLimit.toLocaleString("id-ID") : "";
  }
  if (document.getElementById("inputBakiDebetInput")) {
    document.getElementById("inputBakiDebetInput").value = initialBakiDebet ? initialBakiDebet.toLocaleString("id-ID") : "";
  }
  if (document.getElementById("inputExposureTier")) {
    document.getElementById("inputExposureTier").value = debtor.tierEksposur || "Tier 2";
  }

  updateMoneyInputPreviews();

  // Show modal
  const modal = document.getElementById("commitmentModal");
  modal.classList.remove("hidden");
  if (window.lucide) lucide.createIcons();
}

function closeCommitmentModal() {
  const modal = document.getElementById("commitmentModal");
  modal.classList.add("hidden");
  state.activeCommitmentDebtor = null;
}

async function handleCommitmentSubmit(e) {
  e.preventDefault();

  const debtorId = document.getElementById("commitmentDebtorId").value;
  const debtorIndex = state.debtors.findIndex(d => d.id === debtorId);

  if (debtorIndex === -1) return;

  const targetDebtor = state.debtors[debtorIndex];
  const newTgl = document.getElementById("inputTglKomitmen").value;
  const newStatus = document.getElementById("inputStatusKomitmen").value;
  const newKeterangan = document.getElementById("inputKeterangan") ? document.getElementById("inputKeterangan").value : "";

  const rawNominal = document.getElementById("inputNominal") ? document.getElementById("inputNominal").value : "";
  const newNominal = parseMoney(rawNominal);

  const rawLimit = document.getElementById("inputLimit") ? document.getElementById("inputLimit").value : "";
  const newLimit = parseMoney(rawLimit) || targetDebtor.limit;

  const rawBakiDebet = document.getElementById("inputBakiDebetInput") ? document.getElementById("inputBakiDebetInput").value : "";
  const newBakiDebet = parseMoney(rawBakiDebet) || (newNominal > 0 ? newNominal : targetDebtor.bakiDebet);

  const newOfficerId = document.getElementById("inputOfficerID") ? document.getElementById("inputOfficerID").value : (targetDebtor.officerId || "");
  const newOfficerPairId = document.getElementById("inputOfficerPairID") ? document.getElementById("inputOfficerPairID").value : (targetDebtor.officerPairId || "");
  const newTier = document.getElementById("inputExposureTier") ? document.getElementById("inputExposureTier").value : (targetDebtor.tierEksposur || "Tier 2");

  // Update local state in memory
  state.debtors[debtorIndex].tglKomitmen = newTgl;
  state.debtors[debtorIndex].statusKomitmen = newStatus;
  state.debtors[debtorIndex].keterangan = newKeterangan;
  state.debtors[debtorIndex].nominalKomitmen = newNominal;
  state.debtors[debtorIndex].limit = newLimit;
  state.debtors[debtorIndex].bakiDebet = newBakiDebet;
  state.debtors[debtorIndex].tierEksposur = newTier;
  state.debtors[debtorIndex].officerId = newOfficerId;
  state.debtors[debtorIndex].officerPairId = newOfficerPairId;
  state.debtors[debtorIndex].pengelola = resolveOfficerName(newOfficerId) + (newOfficerId ? ` (${newOfficerId})` : "");
  state.debtors[debtorIndex].pairing = resolvePairingTeam(newOfficerPairId, newOfficerId) + (newOfficerPairId ? ` (${newOfficerPairId})` : "");

  // Send PUT request to backend API with exact UpdateCommitmentRequest structure
  try {
    const updatePayload = {
      commitment_date: newTgl,
      status: newStatus,
      reason: newStatus,
      remarks: newKeterangan,
      nominal: newNominal,
      officer_id: newOfficerId,
      officer_pair_id: newOfficerPairId,
      credit_limit: newLimit,
      outstanding_balance: newBakiDebet,
      exposure_tier: newTier
    };

    const commitmentIdParam = targetDebtor.id || debtorId;
    const response = await fetch(`http://localhost:8080/api/v1/commitments/${commitmentIdParam}`, {
      method: "PUT",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(updatePayload)
    });

    const data = await response.json();
    if (!response.ok) {
      console.warn("[BUMA API Sync] PUT returned non-ok status:", response.status, data);
      showToast("Peringatan Sync Backend", data.error || data.message || "Gagal memperbarui ke server", "warning");
    } else {
      console.log("[BUMA API Sync] Commitment updated successfully in backend:", data);
    }
  } catch (err) {
    console.warn("[BUMA API Sync] Backend offline or unreachable, using local state:", err.message);
  }

  // Re-fetch from API & refresh UI
  await loadAndRenderData();

  closeCommitmentModal();
  showToast(
    "Komitmen Diperbarui",
    `Data komitmen debitur ${targetDebtor.namaDebitur} berhasil disimpan!`,
    "success"
  );
}

// ==========================================
// 8.1 MODAL 3: CREATE NEW COMMITMENT (POST /api/v1/commitments) LOGIC
// ==========================================

function getCreateCommitmentPayload() {
  return {
    account_no: document.getElementById("createAccountNo")?.value.trim() || "",
    debtor_name: document.getElementById("createDebtorName")?.value.trim() || "",
    officer_id: document.getElementById("createOfficerID")?.value.trim() || "",
    officer_pair_id: document.getElementById("createOfficerPairID")?.value.trim() || "",
    product: document.getElementById("createProduct")?.value.trim() || "",
    exposure_tier: document.getElementById("createExposureTier")?.value || "Tier 2",
    outstanding_balance: parseMoney(document.getElementById("createOutstandingBalance")?.value),
    credit_limit: parseMoney(document.getElementById("createCreditLimit")?.value),
    nominal: parseMoney(document.getElementById("createOutstandingBalance")?.value),
    commitment_date: document.getElementById("createCommitmentDate")?.value || "",
    status: document.getElementById("createStatus")?.value.trim() || "",
    reason: document.getElementById("createReason")?.value.trim() || "",
    remarks: document.getElementById("createRemarks")?.value.trim() || ""
  };
}

async function openCreateCommitmentModal() {
  const modal = document.getElementById("createCommitmentModal");
  if (!modal) return;

  await fetchOfficersFromApi();
  modal.classList.remove("hidden");
  if (window.lucide) lucide.createIcons();
}

function closeCreateCommitmentModal() {
  const modal = document.getElementById("createCommitmentModal");
  if (!modal) return;
  modal.classList.add("hidden");
}

async function handleCreateCommitmentApiSubmit() {
  const payload = getCreateCommitmentPayload();

  if (!payload.account_no) {
    showToast("Validasi Gagal", "Field account_no wajib diisi!", "error");
    return;
  }
  if (!payload.reason) {
    showToast("Validasi Gagal", "Field reason wajib diisi!", "error");
    return;
  }

  const btn = document.getElementById("submitCreateCommitmentBtn");
  if (btn) {
    btn.disabled = true;
    btn.innerHTML = `<i data-lucide="loader-2" class="w-4 h-4 animate-spin"></i><span>Mengirim...</span>`;
    if (window.lucide) lucide.createIcons();
  }

  try {
    const response = await fetch("http://localhost:8080/api/v1/commitments", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(payload)
    });

    const data = await response.json();

    // If HTTP status is not 201 (Created), display error in FE
    if (response.status !== 201) {
      const errorMsg = data.error || data.message || `Request gagal (Status Code: ${response.status})`;
      throw new Error(errorMsg);
    }

    // HTTP status 201 Created Success
    const item = data.data || {};
    const createdID = item.id || "OK";
    const successMsg = data.message || "Commitment recorded successfully";

    showToast(
      "Berhasil Dibuat!",
      `${successMsg} (ID: ${createdID})`,
      "success"
    );

    // Sync to local FE state table
    const existingIndex = state.debtors.findIndex(d => d.noRekening === (item.account_no || payload.account_no));
    if (existingIndex !== -1) {
      state.debtors[existingIndex].tglKomitmen = item.commitment_date || payload.commitment_date;
      state.debtors[existingIndex].statusKomitmen = item.status || payload.status;
      state.debtors[existingIndex].keterangan = item.remarks || payload.remarks;
    } else {
      state.debtors.unshift({
        id: createdID,
        noRekening: item.account_no || payload.account_no,
        namaDebitur: payload.debtor_name || "Debitur Baru",
        produk: payload.product || "KMK",
        limit: payload.credit_limit,
        bakiDebet: item.nominal || payload.outstanding_balance,
        tierEksposur: payload.exposure_tier || "Tier 2",
        pengelola: item.officer_id || payload.officer_id || "OFF-003",
        pairing: item.officer_pair_id || payload.officer_pair_id || "OFF-001",
        tglKomitmen: item.commitment_date || payload.commitment_date,
        statusKomitmen: item.status || payload.status,
        keterangan: item.remarks || payload.remarks,
        phone: "081298765432"
      });
    }

    await loadAndRenderData();
    closeCreateCommitmentModal();
  } catch (err) {
    console.error("[Create Commitment Error]", err);
    showToast("Gagal Membuat Komitmen", err.message, "error");
  } finally {
    if (btn) {
      btn.disabled = false;
      btn.innerHTML = `<i data-lucide="send" class="w-4 h-4"></i><span>Kirim Ke API</span>`;
      if (window.lucide) lucide.createIcons();
    }
  }
}

// ==========================================
// 8.2 MODAL 4: CREATE NEW OFFICER (POST /api/v1/officers) LOGIC
// ==========================================

function getCreateOfficerPayload() {
  return {
    officer_id: document.getElementById("createOfficerIDInput")?.value.trim() || "",
    full_name: document.getElementById("createOfficerFullName")?.value.trim() || "",
    email: document.getElementById("createOfficerEmail")?.value.trim() || "",
    pairing_team: document.getElementById("createOfficerPairingTeam")?.value.trim() || "",
    phone: document.getElementById("createOfficerPhone")?.value.trim() || ""
  };
}

function openCreateOfficerModal() {
  const modal = document.getElementById("createOfficerModal");
  if (!modal) return;

  modal.classList.remove("hidden");
  if (window.lucide) lucide.createIcons();
}

function closeCreateOfficerModal() {
  const modal = document.getElementById("createOfficerModal");
  if (!modal) return;
  modal.classList.add("hidden");
}

async function handleCreateOfficerApiSubmit() {
  const payload = getCreateOfficerPayload();

  if (!payload.full_name) {
    showToast("Validasi Gagal", "Field full_name wajib diisi!", "error");
    return;
  }
  if (!payload.email) {
    showToast("Validasi Gagal", "Field email wajib diisi!", "error");
    return;
  }
  if (!payload.pairing_team) {
    showToast("Validasi Gagal", "Field pairing_team wajib diisi!", "error");
    return;
  }

  const btn = document.getElementById("submitCreateOfficerBtn");
  if (btn) {
    btn.disabled = true;
    btn.innerHTML = `<i data-lucide="loader-2" class="w-4 h-4 animate-spin"></i><span>Mengirim...</span>`;
    if (window.lucide) lucide.createIcons();
  }

  try {
    const response = await fetch("http://localhost:8080/api/v1/officers", {
      method: "POST",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(payload)
    });

    const data = await response.json();

    if (response.status !== 201) {
      const errorMsg = data.error || data.message || `Request gagal (Status Code: ${response.status})`;
      throw new Error(errorMsg);
    }

    const item = data.data || {};
    const createdID = item.officer_id || payload.officer_id || "OFF-NEW";
    const successMsg = data.message || "Officer added successfully";

    showToast(
      "Officer Berhasil Ditambahkan!",
      `${successMsg} (ID: ${createdID})`,
      "success"
    );

    if (createdID) {
      OFFICERS_DICTIONARY[createdID] = {
        name: payload.full_name,
        pair: payload.pairing_team
      };
    }

    await loadAndRenderData();
    closeCreateOfficerModal();
  } catch (err) {
    console.error("[Create Officer Error]", err);
    showToast("Gagal Membuat Officer", err.message, "error");
  } finally {
    if (btn) {
      btn.disabled = false;
      btn.innerHTML = `<i data-lucide="send" class="w-4 h-4"></i><span>Kirim Ke API</span>`;
      if (window.lucide) lucide.createIcons();
    }
  }
}

// ==========================================
// 8.3 MODAL 5: IMPORT EXCEL (POST /api/v1/commitments/import-excel) LOGIC
// ==========================================

let selectedExcelFile = null;

function openImportExcelModal() {
  const modal = document.getElementById("importExcelModal");
  if (!modal) return;

  selectedExcelFile = null;
  const fileInput = document.getElementById("excelFileInput");
  if (fileInput) fileInput.value = "";

  document.getElementById("excelFileName").textContent = "Klik atau Seret File Excel ke Sini";
  document.getElementById("excelFileSize").textContent = "Format didukung: .xlsx, .xls (Maks 10MB)";

  const submitBtn = document.getElementById("submitImportExcelBtn");
  if (submitBtn) submitBtn.disabled = true;

  const resultBox = document.getElementById("importResultBox");
  if (resultBox) resultBox.classList.add("hidden");

  modal.classList.remove("hidden");
  if (window.lucide) lucide.createIcons();
}

function closeImportExcelModal() {
  const modal = document.getElementById("importExcelModal");
  if (!modal) return;
  modal.classList.add("hidden");
}

function handleExcelFileSelect(file) {
  if (!file) return;

  if (!file.name.endsWith(".xlsx") && !file.name.endsWith(".xls")) {
    showToast("Format File Tidak Valid", "Hanya file .xlsx dan .xls yang didukung!", "error");
    return;
  }

  selectedExcelFile = file;
  document.getElementById("excelFileName").textContent = file.name;
  const sizeMB = (file.size / (1024 * 1024)).toFixed(2);
  document.getElementById("excelFileSize").textContent = `Ukuran file: ${sizeMB} MB`;

  const submitBtn = document.getElementById("submitImportExcelBtn");
  if (submitBtn) submitBtn.disabled = false;
}

async function handleImportExcelSubmit() {
  if (!selectedExcelFile) {
    showToast("File Belum Dipilih", "Silakan pilih file Excel terlebih dahulu!", "error");
    return;
  }

  const btn = document.getElementById("submitImportExcelBtn");
  if (btn) {
    btn.disabled = true;
    btn.innerHTML = `<i data-lucide="loader-2" class="w-4 h-4 animate-spin"></i><span>Memproses Import...</span>`;
    if (window.lucide) lucide.createIcons();
  }

  const formData = new FormData();
  formData.append("file", selectedExcelFile);

  try {
    const response = await fetch("http://localhost:8080/api/v1/commitments/import-excel", {
      method: "POST",
      body: formData
    });

    const data = await response.json();

    if (response.status !== 200 && response.status !== 201) {
      const errorMsg = data.error || data.message || `Request gagal (Status Code: ${response.status})`;
      throw new Error(errorMsg);
    }

    const result = data.data || {};
    const totalProc = result.total_processed || 0;
    const totalSucc = result.total_success || 0;
    const totalFail = result.total_failed || 0;
    const errors = result.errors || [];

    const resultBox = document.getElementById("importResultBox");
    const summaryEl = document.getElementById("importResultSummary");
    const detailsEl = document.getElementById("importResultDetails");

    if (resultBox && summaryEl) {
      resultBox.classList.remove("hidden");
      summaryEl.innerHTML = `
        <span class="text-emerald-700 font-bold">Total: ${totalProc} | Berhasil: ${totalSucc}</span>
        <span class="${totalFail > 0 ? 'text-rose-600 font-bold' : 'text-slate-500'}">Gagal: ${totalFail}</span>
      `;

      if (errors.length > 0 && detailsEl) {
        detailsEl.classList.remove("hidden");
        detailsEl.innerHTML = errors.map(e => `<div>Row ${e.row}: ${e.error}</div>`).join("");
      } else if (detailsEl) {
        detailsEl.classList.add("hidden");
      }
    }

    showToast(
      "Import Excel Selesai!",
      `Berhasil mengimpor ${totalSucc} dari ${totalProc} baris komitmen.`,
      totalFail === 0 ? "success" : "warning"
    );

    await loadAndRenderData();
  } catch (err) {
    console.error("[Import Excel Error]", err);
    showToast("Gagal Import Excel", err.message, "error");
  } finally {
    if (btn) {
      btn.disabled = false;
      btn.innerHTML = `<i data-lucide="upload" class="w-4 h-4"></i><span>Proses Import Excel</span>`;
      if (window.lucide) lucide.createIcons();
    }
  }
}

// ==========================================
// 9. EVENT LISTENERS & INITIALIZATION
// ==========================================

document.addEventListener("DOMContentLoaded", () => {
  // Initialize Header Date Display
  const currentDateDisplay = document.getElementById("currentDateDisplay");
  if (currentDateDisplay) {
    currentDateDisplay.textContent = `Real-time System | ${formatDate(getTodayISO())}`;
  }

  // Initial Data Load & Render from GET /api/v1/commitments
  loadAndRenderData();

  // Money Input Formatting & Live Preview Listeners
  ["inputNominal", "inputLimit", "inputBakiDebetInput"].forEach(id => {
    const el = document.getElementById(id);
    if (el) {
      el.addEventListener("input", updateMoneyInputPreviews);
      el.addEventListener("blur", () => {
        const val = parseMoney(el.value);
        if (val > 0) {
          el.value = val.toLocaleString("id-ID");
        }
        updateMoneyInputPreviews();
      });
    }
  });

  // Search Input Event
  const searchInput = document.getElementById("searchInput");
  const clearSearchBtn = document.getElementById("clearSearchBtn");

  searchInput.addEventListener("input", (e) => {
    state.searchQuery = e.target.value;
    if (state.searchQuery) {
      clearSearchBtn.classList.remove("hidden");
    } else {
      clearSearchBtn.classList.add("hidden");
    }
    state.currentPage = 1;
    loadAndRenderData();
  });

  clearSearchBtn.addEventListener("click", () => {
    searchInput.value = "";
    state.searchQuery = "";
    clearSearchBtn.classList.add("hidden");
    state.currentPage = 1;
    loadAndRenderData();
  });

  // Dropdown Filters
  document.getElementById("tierFilter").addEventListener("change", (e) => {
    state.selectedTier = e.target.value;
    state.currentPage = 1;
    loadAndRenderData();
  });

  document.getElementById("pengelolaFilter").addEventListener("change", (e) => {
    state.selectedPengelola = e.target.value;
    state.currentPage = 1;
    loadAndRenderData();
  });

  document.getElementById("produkFilter").addEventListener("change", (e) => {
    state.selectedProduk = e.target.value;
    state.currentPage = 1;
    loadAndRenderData();
  });


  // Reset Filters Button
  const resetFilters = () => {
    state.searchQuery = "";
    state.selectedTier = "ALL";
    state.selectedPengelola = "ALL";
    state.selectedProduk = "ALL";
    state.currentPage = 1;

    searchInput.value = "";
    clearSearchBtn.classList.add("hidden");
    document.getElementById("tierFilter").value = "ALL";
    document.getElementById("pengelolaFilter").value = "ALL";
    document.getElementById("produkFilter").value = "ALL";

    loadAndRenderData();
    showToast("Filter Direset", "Semua kriteria filter telah dikembalikan ke awal", "info");
  };

  document.getElementById("resetFiltersBtn").addEventListener("click", resetFilters);
  document.getElementById("emptyResetBtn").addEventListener("click", resetFilters);

  // Items per page dropdown
  document.getElementById("itemsPerPage").addEventListener("change", (e) => {
    state.itemsPerPage = parseInt(e.target.value, 10);
    state.currentPage = 1;
    loadAndRenderData();
  });

  // Column Headers Sorting
  document.querySelectorAll("th[data-sort]").forEach(th => {
    th.addEventListener("click", () => {
      const col = th.getAttribute("data-sort");
      if (state.sortColumn === col) {
        state.sortDirection = state.sortDirection === "asc" ? "desc" : "asc";
      } else {
        state.sortColumn = col;
        state.sortDirection = "asc";
      }
      loadAndRenderData();
    });
  });


  // Refresh & Export Buttons
  document.getElementById("refreshBtn").addEventListener("click", () => {
    loadAndRenderData();
    showToast("Data Commitment Diperbarui", "Data komitmen terbaru berhasil dimuat dari API GET /api/v1/commitments", "success");
  });

  document.getElementById("exportCsvBtn").addEventListener("click", () => {
    showToast("Export Data", "File CSV data debitur berhasil di-download!", "success");
  });

  document.getElementById("quickDispatchBtn").addEventListener("click", () => {
    showToast("Batch Dispatch Scheduled", "Wording penagihan telah di-antrekan untuk 12 debitur", "info");
  });

  // Modal 1 (Wording Generator) Events
  document.getElementById("closeWordingModal").addEventListener("click", closeWordingModal);
  document.getElementById("closeWordingBottomBtn").addEventListener("click", closeWordingModal);
  document.getElementById("copyWordingBtn").addEventListener("click", copyWordingToClipboard);
  document.getElementById("openWaBtn").addEventListener("click", openWhatsAppWeb);

  document.querySelectorAll(".channel-btn").forEach(btn => {
    btn.addEventListener("click", (e) => {
      document.querySelectorAll(".channel-btn").forEach(b => {
        b.classList.remove("active", "bg-emerald-50", "border-emerald-300", "text-emerald-800");
        b.classList.add("bg-slate-50", "border-slate-200", "text-slate-600");
      });
      const targetBtn = e.currentTarget;
      targetBtn.classList.add("active", "bg-emerald-50", "border-emerald-300", "text-emerald-800");
      targetBtn.classList.remove("bg-slate-50", "border-slate-200", "text-slate-600");

      state.activeWordingChannel = targetBtn.getAttribute("data-channel");
      generateAndSetWordingText();
      resetCopyBtnState();
    });
  });

  document.getElementById("wordingToneSelect").addEventListener("change", (e) => {
    state.activeWordingTone = e.target.value;
    generateAndSetWordingText();
    resetCopyBtnState();
  });

  // Modal 2 (Commitment Form) Events
  document.getElementById("closeCommitmentModal")?.addEventListener("click", closeCommitmentModal);
  document.getElementById("closeCommitmentBottomBtn")?.addEventListener("click", closeCommitmentModal);
  document.getElementById("commitmentForm")?.addEventListener("submit", handleCommitmentSubmit);
  document.getElementById("submitCommitmentBtn")?.addEventListener("click", (e) => {
    e.preventDefault();
    const form = document.getElementById("commitmentForm");
    if (form) {
      if (typeof form.requestSubmit === "function") {
        form.requestSubmit();
      } else {
        handleCommitmentSubmit(e);
      }
    }
  });

  // Modal 3 (Create Commitment JSON Modal) Events
  const createCommitmentBtn = document.getElementById("createCommitmentBtn");
  if (createCommitmentBtn) {
    createCommitmentBtn.addEventListener("click", openCreateCommitmentModal);
  }
  document.getElementById("closeCreateCommitmentModal")?.addEventListener("click", closeCreateCommitmentModal);
  document.getElementById("closeCreateCommitmentBottomBtn")?.addEventListener("click", closeCreateCommitmentModal);
  document.getElementById("submitCreateCommitmentBtn")?.addEventListener("click", handleCreateCommitmentApiSubmit);

  // Modal 4 (Create Officer Modal) Events
  const openCreateOfficerBtn = document.getElementById("openCreateOfficerBtn");
  if (openCreateOfficerBtn) {
    openCreateOfficerBtn.addEventListener("click", openCreateOfficerModal);
  }
  document.getElementById("closeCreateOfficerModal")?.addEventListener("click", closeCreateOfficerModal);
  document.getElementById("closeCreateOfficerBottomBtn")?.addEventListener("click", closeCreateOfficerModal);
  document.getElementById("submitCreateOfficerBtn")?.addEventListener("click", handleCreateOfficerApiSubmit);

  // Modal 5 (Import Excel Modal) Events
  const importExcelBtn = document.getElementById("importExcelBtn");
  if (importExcelBtn) {
    importExcelBtn.addEventListener("click", openImportExcelModal);
  }
  document.getElementById("closeImportExcelModal")?.addEventListener("click", closeImportExcelModal);
  document.getElementById("closeImportExcelBottomBtn")?.addEventListener("click", closeImportExcelModal);
  document.getElementById("submitImportExcelBtn")?.addEventListener("click", handleImportExcelSubmit);

  const excelDropzone = document.getElementById("excelDropzone");
  const excelFileInput = document.getElementById("excelFileInput");

  if (excelDropzone && excelFileInput) {
    excelDropzone.addEventListener("click", () => excelFileInput.click());
    excelFileInput.addEventListener("change", (e) => {
      if (e.target.files.length > 0) {
        handleExcelFileSelect(e.target.files[0]);
      }
    });

    excelDropzone.addEventListener("dragover", (e) => {
      e.preventDefault();
      excelDropzone.classList.add("border-emerald-500", "bg-emerald-50/50");
    });
    excelDropzone.addEventListener("dragleave", () => {
      excelDropzone.classList.remove("border-emerald-500", "bg-emerald-50/50");
    });
    excelDropzone.addEventListener("drop", (e) => {
      e.preventDefault();
      excelDropzone.classList.remove("border-emerald-500", "bg-emerald-50/50");
      if (e.dataTransfer.files.length > 0) {
        handleExcelFileSelect(e.dataTransfer.files[0]);
      }
    });
  }

  // Close modals on Backdrop Click
  window.addEventListener("click", (e) => {
    const wordingModal = document.getElementById("wordingModal");
    const commitmentModal = document.getElementById("commitmentModal");
    const createCommitmentModal = document.getElementById("createCommitmentModal");
    const createOfficerModal = document.getElementById("createOfficerModal");
    const importExcelModal = document.getElementById("importExcelModal");

    if (e.target === wordingModal) closeWordingModal();
    if (e.target === commitmentModal) closeCommitmentModal();
    if (e.target === createCommitmentModal) closeCreateCommitmentModal();
    if (e.target === createOfficerModal) closeCreateOfficerModal();
    if (e.target === importExcelModal) closeImportExcelModal();
  });
});
