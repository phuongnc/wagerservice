package ginutil

import (
	"github.com/gin-gonic/gin"
)

func Response(c *gin.Context, httpCode int, success bool, data interface{}, err error) {
	c.JSON(httpCode, gin.H{
		"success": success,
		"data":    data,
		"error":   err.Error(),
	})
	return
}
