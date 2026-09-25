package wording

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"buma-be/internal/domain"
)

// GenerateCollectionWording generates parameterized wording text based on channel, tone, and debtor data
func GenerateCollectionWording(debtor domain.DebtorAccount, channel string, tone string) string {
	channel = strings.ToLower(strings.TrimSpace(channel))
	tone = strings.ToLower(strings.TrimSpace(tone))

	if channel == "" {
		channel = "whatsapp"
	}
	if tone == "" {
		if debtor.ExposureTier == "Tier 1" {
			tone = "urgent"
		} else {
			tone = "firm"
		}
	}

	formattedBalance := FormatIDR(debtor.OutstandingBalance)
	formattedDate := FormatDate(debtor.CommitmentDate)
	todayStr := FormatDate(time.Now().Format("2006-01-02"))

	officerName := debtor.OfficerName
	if officerName == "" {
		officerName = "Tim Collection"
	}

	pairingTeam := debtor.PairingTeam
	if pairingTeam == "" {
		pairingTeam = "Desk Collection"
	}

	switch channel {
	case "whatsapp":
		switch tone {
		case "soft":
			return fmt.Sprintf(
				"Yth. Bapak/Ibu %s,\n\n"+
					"Selamat pagi/siang. Kami dari Tim Collection PT Bank Utama Mandiri (BUMA) menginfokan bahwa kewajiban angsuran kredit Anda (%s) No. Rekening: *%s* dengan sisa baki debet *%s* akan memasuki tanggal komitmen pada *%s*.\n\n"+
					"Mohon dapat melakukan penyetoran sebelum pukul 17:00 WIB untuk menjaga kualitas kredit Anda tetap lancar.\n\n"+
					"Jika telah melakukan pembayaran, abaikan pesan ini. Terima kasih.\n"+
					"Petugas Pengelola: %s (BUMA Collection Unit)",
				debtor.DebtorName, debtor.ProductType, debtor.AccountNo, formattedBalance, formattedDate, officerName,
			)

		case "urgent":
			return fmt.Sprintf(
				"*SURAT PERINGATAN / SOMASI PRA-HUKUM*\n"+
					"PT BANK UTAMA MANDIRI (BUMA)\n\n"+
					"Kepada Yth. Debitur: *%s*\n"+
					"Fasilitas Kredit: *%s*\n"+
					"No. Rekening: *%s*\n"+
					"Total Baki Debet: *%s*\n\n"+
					"Berdasarkan evaluasi risiko (%s), fasilitas kredit Anda telah dikategorikan DALAM PENAWASAN KHUSUS. Peringatan penagihan resmi telah diproses.\n\n"+
					"Anda diwajibkan melakukan penyelesaian kewajiban / penyetoran komitmen pada *%s*. Kelalaian pembayaran akan mengakibatkan pelaporan kolektibilitas pada SLIK OJK dan tindakan hukum penanganan agunan.\n\n"+
					"Segera konfirmasi ke Pengelola: *%s*.",
				debtor.DebtorName, debtor.ProductType, debtor.AccountNo, formattedBalance, debtor.ExposureTier, formattedDate, officerName,
			)

		default: // "firm"
			return fmt.Sprintf(
				"PEMBERITAHUAN PENAGIHAN KREDIT - BUMA\n\n"+
					"Kepada Yth.\n"+
					"*%s*\n"+
					"No. Rekening: *%s*\n\n"+
					"Diberitahukan bahwa tagihan kewajiban fasilitas kredit *%s* Anda sebesar *%s* telah melewati jatuh tempo. Sesuai catatan komitmen, jadwal pembayaran jatuh pada hari *%s*.\n\n"+
					"Mohon SEGERA melakukan pembayaran melalui rekening penampungan BUMA atau konfirmasi bukti transfer hari ini kepada Account Officer Anda:\n"+
					"*%s* (%s)\n\n"+
					"Hubungi Call Center BUMA jika membutuhkan bantuan kendala transaksi.",
				debtor.DebtorName, debtor.AccountNo, debtor.ProductType, formattedBalance, formattedDate, officerName, pairingTeam,
			)
		}

	case "email":
		return fmt.Sprintf(
			"Subjek: [BUMA COLLECTION] Pemberitahuan Kewajiban Kredit No. Rek %s - Yth. %s\n\n"+
				"Kepada Yth.\n"+
				"Management / Bp/Ibu %s\n\n"+
				"Dengan hormat,\n\n"+
				"Sehubungan dengan fasilitas kredit %s atas nama %s dengan nomor rekening %s, melalui surat ini kami menyampaikan rincian posisi pinjaman Anda per tanggal %s:\n\n"+
				"- Nama Debitur: %s\n"+
				"- No. Rekening: %s\n"+
				"- Jenis Fasilitas: %s\n"+
				"- Baki Debet Pinjaman: %s\n"+
				"- Tanggal Komitmen Bayar: %s\n"+
				"- Tim Pengelola: %s (%s)\n\n"+
				"Dimohon untuk dapat melakukan penyetoran dana ke rekening efektif tepat pada tanggal komitmen yang disepakati.\n\n"+
				"Demikian pemberitahuan ini kami sampaikan. Atas perhatian dan kerja samanya, kami ucapkan terima kasih.\n\n"+
				"Hormat kami,\n"+
				"PT Bank Utama Mandiri (BUMA)\n"+
				"Divisi Special Asset Management & Collection",
			debtor.AccountNo, debtor.DebtorName, debtor.DebtorName, debtor.ProductType, debtor.DebtorName, debtor.AccountNo, todayStr,
			debtor.DebtorName, debtor.AccountNo, debtor.ProductType, formattedBalance, formattedDate, officerName, pairingTeam,
		)

	default: // "call" / "sms"
		return fmt.Sprintf(
			"[SCRIPT TELEPON / SMS BUMA]\n"+
				"Halo Bp/Ibu %s, saya %s dari BUMA. Mengonfirmasi janji bayar fasilitas %s No. Rek %s (Baki debet %s) pada tanggal %s. Mohon dipastikan dana tersedia sebelum pukul 15.00 WIB. Terima kasih.",
			debtor.DebtorName, officerName, debtor.ProductType, debtor.AccountNo, formattedBalance, formattedDate,
		)
	}
}

// FormatIDR formats float amount to Indonesian Rupiah currency string (e.g. Rp 4.250.000.000)
func FormatIDR(amount float64) string {
	intPart := int64(amount)
	str := strconv.FormatInt(intPart, 10)
	n := len(str)
	if n <= 3 {
		return "Rp " + str
	}

	var result []string
	remainder := n % 3
	if remainder > 0 {
		result = append(result, str[:remainder])
	}
	for i := remainder; i < n; i += 3 {
		result = append(result, str[i:i+3])
	}

	return "Rp " + strings.Join(result, ".")
}

// FormatDate converts YYYY-MM-DD string into readable Indonesian Date format (e.g. 6 Sep 2026)
func FormatDate(dateStr string) string {
	if dateStr == "" {
		return "-"
	}

	parsed, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		// try full timestamptz format if applicable
		parsed, err = time.Parse(time.RFC3339, dateStr)
		if err != nil {
			return dateStr
		}
	}

	months := []string{"", "Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Okt", "Nov", "Des"}
	return fmt.Sprintf("%d %s %d", parsed.Day(), months[parsed.Month()], parsed.Year())
}
