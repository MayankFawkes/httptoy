package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func addAllCookies(r *gin.Engine) {
	r.GET("/cookies", cookies_default)
	r.GET("/cookies/set/:key/:value", cookies_set_key_value)
	r.GET("/cookies/delete/:key", cookies_delete_key)
}

type cookies_payload struct {
	Cookies dict `json:"cookies"`
}
type dict map[string]string

func cookies_default(c *gin.Context) {
	cookiesHeader := c.Request.Cookies()
	cookieDict := make(dict)

	for _, cookie := range cookiesHeader {
		cookieDict[cookie.Name] = cookie.Value
	}

	return_payload := cookies_payload{
		Cookies: cookieDict,
	}

	c.JSON(http.StatusOK, return_payload)

}
func cookies_set_key_value(c *gin.Context) {
	cookie_key := c.Params.ByName("key")
	value_key := c.Params.ByName("value")

	c.SetCookie(cookie_key, value_key, 0, "/", "", false, false)

	c.Redirect(http.StatusFound, "/cookies")

}
func cookies_delete_key(c *gin.Context) {
	cookie_key := c.Params.ByName("key")
	c.SetCookie(cookie_key, "", -1, "/", "", false, false)
	c.Redirect(http.StatusFound, "/cookies")
}
