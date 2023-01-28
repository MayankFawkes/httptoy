package routes

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mayankfawkes/httptoy/pkg/parsers"
)

func addAllImages(r *gin.Engine) {
	r.GET("/image", images_default)
	r.GET("/image/png", images_png)
	r.GET("/image/jpeg", images_jpeg)
	r.GET("/image/svg", images_svg)
	r.GET("/image/webp", images_webp)

}

type imageDefaultSuportedType struct {
	Message string   `json:"message"`
	Support []string `json:"support"`
}

var supportedTypes = []string{"image/webp", "image/svg+xml", "image/jpeg", "image/png", "image/*"}
var imagesExtension = map[string]string{
	"image/webp":    "webp",
	"image/svg+xml": "svg",
	"image/jpeg":    "jpeg",
	"image/png":     "png",
	"image/*":       "png",
}

func images_default(c *gin.Context) {
	sup := parsers.ParseSupportedImage(c.Request, supportedTypes)
	fmt.Println("ehe", sup)
	if sup == "" {
		res := imageDefaultSuportedType{
			Message: "Accept header not found or supported type not found",
			Support: supportedTypes,
		}
		c.JSON(http.StatusOK, res)
		return
	}

	if sup == "image/*" {
		sup = "image/png"
	}

	dat, _ := os.ReadFile(fmt.Sprintf("sample_files/images/sample.%s", imagesExtension[sup]))
	c.Data(http.StatusOK, sup, dat)

}

func images_png(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/images/sample.png")
	c.Data(http.StatusOK, CONTENT_IMAGE_PNG, dat)
}
func images_jpeg(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/images/sample.jpeg")
	c.Data(http.StatusOK, CONTENT_IMAGE_JPEG, dat)
}
func images_svg(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/images/sample.svg")
	c.Data(http.StatusOK, CONTENT_IMAGE_SVG, dat)
}
func images_webp(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/images/sample.webp")
	c.Data(http.StatusOK, CONTENT_IMAGE_WEBP, dat)
}
