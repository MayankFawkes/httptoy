package routes

import (
	"github.com/gin-gonic/gin"
)

const (
	DEFAULT_STRING  = "i love httptoy"
	DEFAULT_UNICODE = "i ❤️ httptoy 🎉"
	INVALID_DATA    = "Invalid Data"
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
