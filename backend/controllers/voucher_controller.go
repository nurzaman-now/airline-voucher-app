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

func (c *VoucherController) CheckExistsVoucher(ctx *gin.Context, req voucher.CheckVoucherRequest) (error, *models.Voucers, string) {
	var voucher models.Voucers
	query := `
		SELECT id, crew_name, crew_id, flight_number, flight_date, aircraft_type, seat1, seat2, seat3, created_at 
		FROM vouchers 
		WHERE flight_number = ? AND flight_date = ? 
		LIMIT 1
	`

	err := database.DB.QueryRow(query, req.FlightNumber, req.FlightDate).Scan(
		&voucher.ID,
		&voucher.CrewName,
		&voucher.CrewID,
		&voucher.FlightNumber,
		&voucher.FlightDate,
		&voucher.AircraftType,
		&voucher.Seat1,
		&voucher.Seat2,
		&voucher.Seat3,
		&voucher.CreatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil, "Voucher belum tersedia"
		}
		return err, nil, "Terjadi kesalahan saat memeriksa database"
	}

	return nil, &voucher, "Voucher sudah ada"
}

func (c *VoucherController) CheckVoucher(ctx *gin.Context) {
	// Menginisialisasi struct untuk menyimpan data dari request
	var req voucher.CheckVoucherRequest

	// Validasi Input JSON
	// ShouldBindJSON akan mencocokkan request body dengan struct CheckRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		reqErr := handlers.GetErrorMap(err)
		reqMessages := handlers.MapToString(reqErr)
		handlers.ResponseError(ctx, nil, "Data yang di inputkan tidak sesuai: "+reqMessages)
		return
	}

	// Cek apakah voucher sudah ada
	errResult, voucherResult, msg := c.CheckExistsVoucher(ctx, req)
	if errResult != nil || voucherResult != nil {
		handlers.ResponseError(ctx, voucherResult, msg)
		return
	}

	// Kembalikan berhasil jika Voucher tersedia
	handlers.ResponseSuccess(ctx, voucherResult, "Voucher bisa di generate")
}

func (c *VoucherController) GenerateVoucher(ctx *gin.Context) {
	// Menginisialisasi struct untuk menyimpan data dari request
	var req voucher.CreateVoucersRequest

	// Validasi Input JSON
	// ShouldBindJSON akan mencocokkan request body dengan struct CreateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		reqErr := handlers.GetErrorMap(err)
		reqMessages := handlers.MapToString(reqErr)
		handlers.ResponseError(ctx, nil, "Data yang di inputkan tidak sesuai: "+reqMessages)
		return
	}

	// inisiasi data untuk cek voucher
	var reqCheck = voucher.CheckVoucherRequest{
		FlightNumber: req.FlightNumber,
		FlightDate:   req.FlightDate,
	}

	// Cek apakah voucher sudah ada
	errResult, voucherResult, msg := c.CheckExistsVoucher(ctx, reqCheck)
	if errResult != nil || voucherResult != nil {
		handlers.ResponseError(ctx, voucherResult, msg)
		return
	}

	// Generate 3 kursi unik secara acak berdasarkan tipe pesawat
	seats, err := handlers.GenerateRandomSeats(req.AircraftType)
	if err != nil {
		handlers.ResponseError(ctx, nil, "Gagal menghasilkan kursi: "+err.Error())
		return
	}

	dataVoucher := models.Voucers{
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
	voucherCreated, err := database.DB.Exec(query, dataVoucher.CrewName, dataVoucher.CrewID, dataVoucher.FlightNumber, dataVoucher.FlightDate, dataVoucher.AircraftType, dataVoucher.Seat1, dataVoucher.Seat2, dataVoucher.Seat3, dataVoucher.CreatedAt)
	if err != nil {
		// Jika terjadi error pada sisi server/database
		handlers.ResponseError(ctx, nil, "Gagal menyimpan Voucher")
		return
	}

	// Mendapatkan ID Voucher yang baru dibuat
	id, err := voucherCreated.LastInsertId()
	if err != nil {
		handlers.ResponseError(ctx, nil, "Gagal mendapatkan ID Voucher")
		return
	}

	// Mengupdate ID Voucher
	dataVoucher.ID = int(id)

	// Mengembalikan data voucher yang berhasil disimpan
	handlers.ResponseSuccess(ctx, dataVoucher, "Generate Voucher Berhasil")
}
