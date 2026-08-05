package voucher

// CreateVoucersRequest untuk endpoint POST /api/generate
type CreateVoucersRequest struct {
	CrewName     string `json:"crew_name" binding:"required"`
	CrewId       string `json:"crew_id" binding:"required"`
	FlightNumber string `json:"flight_number" binding:"required"`
	FlightDate   string `json:"flight_date" binding:"required,datetime=2006-01-02 15:04:05"`
	AircraftType string `json:"aircraft_type" binding:"required"`
}
