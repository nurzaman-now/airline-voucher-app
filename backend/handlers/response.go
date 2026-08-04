package handlers

import "github.com/gin-gonic/gin"

func ResponseSuccess(ctx *gin.Context, data any, message string) {
	ctx.JSON(200, gin.H{
		"status":  "success",
		"data":    data,
		"message": message,
	})
}

func ResponseError(ctx *gin.Context, data any, message string) {
	ctx.JSON(400, gin.H{
		"status":  "error",
		"data":    data,
		"message": message,
	})
}
