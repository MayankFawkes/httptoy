package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayankfawkes/httptoy/pkg/routes"
)

func Server() *gin.Engine {
	r := gin.Default()
	routes.AddRoutes(r)

	r.LoadHTMLGlob("templates/*.html")

	// r.GET("/json", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"status":  "posted",
	// 		"message": "Hi...",
	// 		"nick":    "Nick here",
	// 	})
	// })

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	return r
}
