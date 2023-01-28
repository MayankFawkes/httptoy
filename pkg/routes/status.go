package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func addAllStatus(r *gin.Engine) {
	r.GET("/status/:code", status_universal)
	r.POST("/status/:code", status_universal)
	r.DELETE("/status/:code", status_universal)
	r.PUT("/status/:code", status_universal)
	r.PATCH("/status/:code", status_universal)
}

func status_universal(c *gin.Context) {
	statusCodeString := c.Params.ByName("code")
	statusCode, err := strconv.Atoi(statusCodeString)

	if http.StatusText(statusCode) == "" {
		c.AbortWithStatus(http.StatusBadGateway)
		return
	}

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.AbortWithStatus(statusCode)
}
