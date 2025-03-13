package utils

import "github.com/gin-gonic/gin"

func SuccessResponse(c *gin.Context, status int, data interface{}) {
	c.JSON(status, gin.H{
		"status": "success",
		"data":   data,
	})
}
