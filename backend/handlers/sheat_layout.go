package handlers

// Menyimpan konfigurasi tata letak kursi berdasarkan dokumen tes
var SheatLayout = map[string]struct {
	MaxRow int
	Seats  []string
}{
	ATR: {MaxRow: 18, Seats: []string{"A", "C", "D", "F"}},
	Airbus320: {MaxRow: 32, Seats: []string{"A", "B", "C", "D", "E", "F"}},
	Boeing737Max: {MaxRow: 32, Seats: []string{"A", "B", "C", "D", "E", "F"}},
}
