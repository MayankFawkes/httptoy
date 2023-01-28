package routes

import (
	"crypto/rand"
	"encoding/base64"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/mayankfawkes/httptoy/pkg/parsers"
)

func addAllResponse(r *gin.Engine) {
	r.GET("/bytes/:count", response_bytes_count)
	r.GET("/base64en/:data", response_base64en)
	r.GET("/base64de/:data", response_base64de)

	r.GET("/delay/:seconds", response_delay)
	r.POST("/delay/:seconds", response_delay)
	r.DELETE("/delay/:seconds", response_delay)
	r.PUT("/delay/:seconds", response_delay)
	r.PATCH("/delay/:seconds", response_delay)

	r.GET("/uuid", response_uuid)
	r.GET("/body/:data", response_body_data)

	r.GET("/text", response_text)
	r.GET("/json", response_json)
	r.GET("/xml", response_xml)
	r.GET("/yaml", response_yaml)
	r.GET("/html", response_html)

}

func response_bytes_count(c *gin.Context) {
	countString := c.Params.ByName("count")
	count, err := strconv.Atoi(countString)

	if err != nil {
		count = 1024
	}

	if count > 1024 {
		count = 1024
	}

	randomBytes := make([]byte, count)
	rand.Read(randomBytes)

	c.Data(http.StatusOK, CONTENT_BYTES, randomBytes)
}

func response_base64en(c *gin.Context) {
	dat := c.Params.ByName("data")
	rawDecodedText := base64.StdEncoding.EncodeToString([]byte(dat))
	c.Data(http.StatusOK, CONTENT_TEXT, []byte(rawDecodedText))
}

func response_base64de(c *gin.Context) {
	dat := c.Params.ByName("data")
	rawDecodedText, err := base64.StdEncoding.DecodeString(dat)

	if err != nil {
		c.Data(http.StatusOK, CONTENT_TEXT, []byte(INVALID_DATA))
		return
	}

	c.Data(http.StatusOK, CONTENT_TEXT, []byte(rawDecodedText))
}
func response_delay(c *gin.Context) {
	secString := c.Params.ByName("seconds")
	sec, err := strconv.Atoi(secString)

	if err != nil {
		sec = 0
	}

	if sec > 10 {
		sec = 10
	}

	time.Sleep(time.Duration(sec) * time.Second)

	rt := method_get_payload{
		Args:    parsers.ParseParams(c.Request),
		Headers: parsers.ParseHeaders(c.Request),
		Origin:  c.ClientIP(),
		Url:     parsers.ParseURL(c.Request),
	}

	c.JSON(http.StatusOK, rt)
}

func response_uuid(c *gin.Context) {
	uid := uuid.New()
	c.Data(http.StatusOK, CONTENT_TEXT, []byte(uid.String()))
}
func response_body_data(c *gin.Context) {
	dat := c.Params.ByName("data")
	c.Data(http.StatusOK, CONTENT_TEXT, []byte(dat))
}

func response_text(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/data/sample.text")
	c.Data(http.StatusOK, CONTENT_TEXT, dat)
}

func response_json(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/data/sample.json")
	c.Data(http.StatusOK, CONTENT_JSON, dat)
}

func response_xml(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/data/sample.xml")
	c.Data(http.StatusOK, CONTENT_XML, dat)
}

func response_yaml(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/data/sample.yaml")
	c.Data(http.StatusOK, CONTENT_YAML, dat)
}

func response_html(c *gin.Context) {
	dat, _ := os.ReadFile("sample_files/data/sample.html")
	c.Data(http.StatusOK, CONTENT_HTML, dat)
}
