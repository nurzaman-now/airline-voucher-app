package models

// inisialisasi struct Voucers
type Voucers struct {
	ID           int    `json:"id"`
	CrewName     string `json:"crew_name" binding:"required"`
	CrewID       string `json:"crew_id" binding:"required"`
	FlightNumber string `json:"flight_number" binding:"required"`
	FlightDate   string `json:"flight_date" binding:"required"`
	AircraftType string `json:"aircraft_type" binding:"required"`
	Seat1        string `json:"seat1" binding:"required"`
	Seat2        string `json:"seat2" binding:"required"`
	Seat3        string `json:"seat3" binding:"required"`
	CreatedAt    string `json:"created_at" binding:"required"`
}
