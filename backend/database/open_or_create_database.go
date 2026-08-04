package database

import (
	"database/sql"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func OpenOrCreateDatabase() {
	var err error
	// Membuka koneksi ke file vouchers.db
	DB, err = sql.Open("sqlite", "./database/vouchers.db")
	if err != nil {
		log.Fatalf("Gagal terhubung ke database: %v", err)
	}

	// Membuat tabel vouchers jika belum ada
	createTableQuery := `
	CREATE TABLE IF NOT EXISTS vouchers (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		crew_name TEXT NOT NULL,
		crew_id TEXT NOT NULL,
		flight_number TEXT NOT NULL,
		flight_date TEXT NOT NULL,
		aircraft_type TEXT NOT NULL,
		seat1 TEXT NOT NULL,
		seat2 TEXT NOT NULL,
		seat3 TEXT NOT NULL,
		created_at TEXT NOT NULL,
		UNIQUE(flight_number, flight_date)
	);`

	_, err = DB.Exec(createTableQuery)
	if err != nil {
		log.Fatalf("Gagal membuat tabel: %v", err)
	}

	log.Println("Database berhasil diinisialisasi!")
}
