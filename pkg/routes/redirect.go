package routes

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func addAllRequestRedirect(r *gin.Engine) {
	r.GET("/redirect/:times", redirect_redirect_times)
	r.GET("/redirect-delay/:delay/:times", redirect_redirect_delay_times)
}

func redirect_redirect_times(c *gin.Context) {
	timeString := c.Params.ByName("times")
	times, err := strconv.Atoi(timeString)

	if err != nil {
		times = 0
	}

	if times <= 1 {
		c.Redirect(http.StatusFound, "/get")
	} else {
		c.Redirect(http.StatusFound, fmt.Sprintf("/redirect/%d", times-1))
	}

}

func redirect_redirect_delay_times(c *gin.Context) {
	delayString := c.Params.ByName("delay")
	timeString := c.Params.ByName("times")

	delay, err := strconv.Atoi(delayString)
	if err != nil {
		delay = 0
	}

	times, err := strconv.Atoi(timeString)
	if err != nil {
		times = 0
	}

	time.Sleep(time.Duration(delay) * time.Second)

	if times <= 1 {
		c.Redirect(http.StatusFound, "/get")
	} else {
		c.Redirect(http.StatusFound, fmt.Sprintf("/redirect-delay/%d/%d", delay, times-1))
	}

}
