package server

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mayankfawkes/httptoy/pkg/routes"
)

func Server() *gin.Engine {
	r := gin.Default()
	routes.AddRoutes(r)

	r.LoadHTMLGlob("templates/*.html")

	r.GET("/", index_html)
	r.GET("/version", version)
	r.StaticFile("/favicon.ico", "sample_files/images/favicon.ico")

	return r
}

func index_html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func version(c *gin.Context) {
	version_text := fmt.Sprintf("VERSION=%s\nGIT_SHA=%s", os.Getenv("APP_VERSION"), os.Getenv("GIT_SHA"))

	c.Data(http.StatusOK, routes.CONTENT_TEXT, []byte(version_text))
}
