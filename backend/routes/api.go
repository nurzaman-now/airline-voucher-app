package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

// SetupRouter - Setup semua routes untuk aplikasi
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// Inisialisasi controller
	voucherController := controllers.InitVoucherController()

	// Membuat grup endpoint API
	api := router.Group("/api")
	{
		// Endpoint untuk mengecek ketersediaan voucher
		api.POST("/check", voucherController.CheckVoucher)

		// Endpoint untuk membuat voucher baru
		api.POST("/generate", voucherController.GenerateVoucher)
	}

	return router
}
