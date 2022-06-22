package utils

import "github.com/gin-gonic/gin"

func ErrorResponse(err error) *gin.H {
	return &gin.H{
		"message": "error",
		"data":    nil,
		"error":   err.Error(),
	}
}

func SuccessResponse(data any) *gin.H {
	return &gin.H{
		"message": "success",
		"data":    data,
		"error":   nil,
	}
}
