package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayankfawkes/httptoy/pkg/encoding"
)

func addAllCompression(r *gin.Engine) {
	r.GET("/deflate", compression_flate)
	r.GET("/brotli", compression_brotli)
	r.GET("/zlib", compression_zlib)
	r.GET("/gzip", compression_gzip)
	r.GET("/utf8", compression_utf8)
}

func compression_flate(c *gin.Context) {
	encodedData := encoding.CompressFlate([]byte(DEFAULT_STRING))
	c.Writer.Header().Set("content-encoding", "deflate")
	c.Data(http.StatusOK, "text/plain; charset=utf-8", encodedData)
}

func compression_brotli(c *gin.Context) {
	encodedData := encoding.CompressBrotli([]byte(DEFAULT_STRING))
	c.Writer.Header().Set("content-encoding", "br")
	c.Data(http.StatusOK, "text/plain; charset=utf-8", encodedData)
}

func compression_gzip(c *gin.Context) {
	encodedData := encoding.CompressGzip([]byte(DEFAULT_STRING))
	c.Writer.Header().Set("content-encoding", "gzip")
	c.Data(http.StatusOK, "text/plain; charset=utf-8", encodedData)
}

func compression_zlib(c *gin.Context) {
	encodedData := encoding.CompressZlib([]byte(DEFAULT_STRING))
	c.Writer.Header().Set("content-encoding", "zlib")
	c.Data(http.StatusOK, "text/plain; charset=utf-8", encodedData)
}

func compression_utf8(c *gin.Context) {
	encodedData := []byte(DEFAULT_UNICODE)
	c.Writer.Header().Set("content-encoding", "utf-8")
	c.Data(http.StatusOK, "text/plain; charset=utf-8", encodedData)
}
