package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mayankfawkes/httptoy/pkg/parsers"
)

func addAllMethods(r *gin.Engine) {
	r.GET("/get", method_get)

	r.POST("/post", method_with_body)
	r.PUT("/put", method_with_body)
	r.DELETE("/delete", method_with_body)
	r.PATCH("/delete", method_with_body)
}

type method_get_payload struct {
	Args    map[string]string `json:"args"`
	Headers map[string]string `json:"headers"`
	Origin  string            `json:"origin"`
	Url     string            `json:"url"`
}

func method_get(c *gin.Context) {
	rt := method_get_payload{
		Args:    parsers.ParseParams(c.Request),
		Headers: parsers.ParseHeaders(c.Request),
		Origin:  c.ClientIP(),
		Url:     parsers.ParseURL(c.Request),
	}

	c.JSON(http.StatusOK, rt)
}

type method_with_body_payload struct {
	Args       map[string]string `json:"args"`
	Headers    map[string]string `json:"headers"`
	Origin     string            `json:"origin"`
	Url        string            `json:"url"`
	Bodytype   string            `json:"body_type"`
	BodyLength int64             `json:"body_length"`
	Data       map[string]string `json:"data"`
	Text       string            `json:"text"`
}

func method_with_body(c *gin.Context) {

	data, text := parsers.ParseBodyData(c.Request, c.ContentType())

	rt := method_with_body_payload{
		Args:       parsers.ParseParams(c.Request),
		Headers:    parsers.ParseHeaders(c.Request),
		Origin:     c.ClientIP(),
		Url:        parsers.ParseURL(c.Request),
		Bodytype:   c.ContentType(),
		BodyLength: c.Request.ContentLength,
		Data:       data,
		Text:       text,
	}

	c.JSON(http.StatusOK, rt)
}
