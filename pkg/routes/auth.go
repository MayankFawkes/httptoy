package routes

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func addAllAuth(r *gin.Engine) {
	r.GET("/basic/:username/:password", auth_basic_user_pass)
	r.GET("/basic", auth_basic)
	r.GET("/bearer", auth_bearer)
}

func auth_basic_user_pass(c *gin.Context) {
	username := c.Params.ByName("username")
	password := c.Params.ByName("password")

	user, pass, hasAuth := c.Request.BasicAuth()

	if hasAuth && user == username && pass == password {
		c.JSON(http.StatusOK, gin.H{
			"username": username,
			"password": password,
			"auth":     true,
		})
	} else {
		c.Writer.Header().Set("WWW-Authenticate", "Basic realm=Restricted")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}

}

func auth_basic(c *gin.Context) {
	user, pass, hasAuth := c.Request.BasicAuth()

	if hasAuth {
		c.JSON(http.StatusOK, gin.H{
			"username": user,
			"password": pass,
			"auth":     true,
		})
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}

func auth_bearer(c *gin.Context) {
	bearer_token := c.Request.Header["Authorization"]

	if len(bearer_token) >= 1 && strings.HasPrefix(bearer_token[0], "Bearer") {
		c.JSON(http.StatusOK, gin.H{
			"authenticated": true,
			"token":         bearer_token[0],
		})
	} else {
		c.Writer.Header().Set("WWW-Authenticate", "Bearer")
		c.AbortWithStatus(http.StatusUnauthorized)
		return
	}
}
