package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayankfawkes/httptoy/pkg/parsers"
)

func addAllRequestInspection(r *gin.Engine) {
	r.GET("/headers", ri_headers)
	r.GET("/header/:header_name", ri_header_by_name)
	r.GET("/ip", ri_client_real_ip)
}

func ri_headers(c *gin.Context) {
	c.JSON(http.StatusOK, parsers.ParseHeaders(c.Request))
}

func ri_header_by_name(c *gin.Context) {
	header_name := c.Params.ByName("header_name")
	hdr := c.GetHeader(header_name)

	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(hdr))
}

func ri_client_real_ip(c *gin.Context) {
	c.Data(http.StatusOK, "text/plain; charset=utf-8", []byte(c.ClientIP()))
}
