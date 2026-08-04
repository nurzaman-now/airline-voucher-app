package controllers

import (
	"backend/database"
	"backend/handlers"
	"backend/models"
	"backend/requests/voucher"
	"database/sql"
	"time"

	"github.com/gin-gonic/gin"
)

type VoucherController struct{}

// InitVoucherController membuat controller baru dengan dependency injection
func InitVoucherController() *VoucherController {
	return &VoucherController{}
}

func (c *VoucherController) CheckVoucher(ctx *gin.Context) {
	// Menginisialisasi struct untuk menyimpan data dari request
	var req voucher.CheckVoucherRequest

	// Validasi Input JSON
	// ShouldBindJSON akan mencocokkan request body dengan struct CheckRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		reqErr := handlers.GetErrorMap(err)
		reqMessages := handlers.MapToString(reqErr)
		handlers.ResponseError(ctx, "Data yang di inputkan tidak sesuai: "+reqMessages)
		return
	}

	// Query ke Database dengan Parameterized Query
	var exists bool
	query := `
		SELECT EXISTS(
			SELECT 1 FROM vouchers 
			WHERE flight_number = ? AND flight_date = ? 
			LIMIT 1
		)
	`

	// Check Voucher di Database
	err := database.DB.QueryRow(query, req.FlightNumber, req.FlightDate).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		// Jika terjadi error pada sisi server/database
		handlers.ResponseError(ctx, "Terjadi kesalahan saat memeriksa database")
		return
	}

	if exists {
		// Jika Voucher sudah ada
		handlers.ResponseError(ctx, "Voucher sudah ada")
		return
	}

	// Kembalikan berhasil jika Voucher tersedia
	handlers.ResponseSuccess(ctx, nil, "Voucher tersedia")
}

func (c *VoucherController) GenerateVoucher(ctx *gin.Context) {
	// Menginisialisasi struct untuk menyimpan data dari request
	var req voucher.CreateVoucersRequest

	// Validasi Input JSON
	// ShouldBindJSON akan mencocokkan request body dengan struct CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		reqErr := handlers.GetErrorMap(err)
		reqMessages := handlers.MapToString(reqErr)
		handlers.ResponseError(ctx, "Data yang di inputkan tidak sesuai: "+reqMessages)
		return
	}

	// Generate 3 kursi unik secara acak berdasarkan tipe pesawat
	seats, err := handlers.GenerateRandomSeats(req.AircraftType)
	if err != nil {
		handlers.ResponseError(ctx, "Gagal menghasilkan kursi: "+err.Error())
		return
	}

	voucher := models.Voucers{
		CrewName:     req.CrewName,
		CrewID:       req.CrewId,
		FlightNumber: req.FlightNumber,
		FlightDate:   req.FlightDate,
		AircraftType: req.AircraftType,
		Seat1:        seats[0],
		Seat2:        seats[1],
		Seat3:        seats[2],
		CreatedAt:    time.Now().Format(time.RFC3339),
	}

	query := `
		INSERT INTO vouchers (crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at) 
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`
	result, err := database.DB.Exec(query, voucher.CrewName, voucher.CrewID, voucher.FlightNumber, voucher.FlightDate, voucher.AircraftType, voucher.Seat1, voucher.Seat2, voucher.Seat3, voucher.CreatedAt)
	if err != nil {
		// Jika terjadi error pada sisi server/database
		handlers.ResponseError(ctx, "Gagal menyimpan Voucher")
		return
	}

	// Mendapatkan ID Voucher yang baru dibuat
	id, err := result.LastInsertId()
	if err != nil {
		handlers.ResponseError(ctx, "Gagal mendapatkan ID Voucher")
		return
	}

	// Mengupdate ID Voucher
	voucher.ID = int(id)

	// Mengembalikan data voucher yang berhasil disimpan
	handlers.ResponseSuccess(ctx, voucher, "Generate Voucher Berhasil")
}
