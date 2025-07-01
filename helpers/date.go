package helpers

import (
	"fmt"
	"time"
)

var indonesiaMonth = [...]string{
	"", "Januari", "Februari", "Maret", "April", "Mei", "Juni",
	"Juli", "Agustus", "September", "Oktober", "November", "Desember",
}

// FormatTanggalIndonesia mengubah time.Time menjadi string "09 Desember 2023"
func FormatTanggalIndonesia(t time.Time) string {
	hari := t.Day()
	bulan := indonesiaMonth[int(t.Month())]
	tahun := t.Year()
	return fmt.Sprintf("%02d %s %d", hari, bulan, tahun)
}
