package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func JSON(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responseData := gin.H{
		"message":   message,
		"data":      data,
		"errors":    errs,
		"status":    http.StatusText(status),
		"timestamp": time.Now().Format("2006-01-02 15:04:05"),
	}

	c.JSON(status, responseData)
}
