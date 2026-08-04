package handlers

import (
	"fmt"
	"math/rand"
	"time"
)

// GenerateRandomSeats menghasilkan 3 kursi unik secara acak berdasarkan tipe pesawat
func GenerateRandomSeats(aircraftType string) ([]string, error) {
	// Mengambil layout kursi berdasarkan tipe pesawat
	layout, ok := SheatLayout[aircraftType]
	if !ok {
		return nil, fmt.Errorf("tipe pesawat tidak valid: %s", aircraftType)
	}

	// Menginisialisasi generator angka acak
	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Menggunakan map untuk menyimpan kursi yang sudah ada
	seen := make(map[string]bool, 3)

	// Membuat slice untuk menyimpan kursi yang akan dihasilkan
	seats := make([]string, 0, 3)

	// Melakukan perulangan hingga mendapatkan 3 kursi unik
	for len(seats) < 3 {
		row := rng.Intn(layout.MaxRow) + 1
		letter := layout.Seats[rng.Intn(len(layout.Seats))]
		seatCode := fmt.Sprintf("%d%s", row, letter)

		if !seen[seatCode] {
			seen[seatCode] = true
			seats = append(seats, seatCode)
		}
	}

	// Mengembalikan 3 kursi unik
	return seats, nil
}
