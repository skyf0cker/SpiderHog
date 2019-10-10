package v1

import (
	"SpiderHog/Spider"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Querychan(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"urlChan":     len(Spider.UrlChan),
		"contentchan": len(Spider.ContentChan),
		"savechan":    len(Spider.SaveChan),
	})
}
