package voucher

// CheckVoucherRequest untuk endpoint POST /api/check
type CheckVoucherRequest struct {
	FlightNumber string `json:"flight_number" binding:"required"`
	FlightDate   string `json:"flight_date" binding:"required"`
}
