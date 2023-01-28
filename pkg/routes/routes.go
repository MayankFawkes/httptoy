package routes

import (
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_STRING  = "i love httptoy"
	DEFAULT_UNICODE = "i ❤️ httptoy 🎉"
	INVALID_DATA    = "Invalid Data"
)

const (
	CONTENT_TEXT  = "text/plain; charset=utf-8"
	CONTENT_JSON  = "application/json; charset=utf-8"
	CONTENT_XML   = "application/xml; charset=utf-8"
	CONTENT_YAML  = "text/yaml; charset=utf-8"
	CONTENT_HTML  = "text/html; charset=utf-8"
	CONTENT_BYTES = "application/octet-stream"

	CONTENT_IMAGE_PNG  = "image/png"
	CONTENT_IMAGE_JPEG = "image/jpeg"
	CONTENT_IMAGE_WEBP = "image/webp"
	CONTENT_IMAGE_SVG  = "image/svg+xml"
)

func AddRoutes(r *gin.Engine) {
	addAllMethods(r)
	addAllAuth(r)
	addAllStatus(r)
	addAllRequestInspection(r)
	addAllCompression(r)
	addAllResponse(r)
	addAllRequestRedirect(r)
	addAllImages(r)
	addAllCookies(r)
}
