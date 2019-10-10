package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

func Getlog(c *gin.Context) {
	var logName string
	logs, _ := ioutil.ReadDir("./log")
	for _, log := range logs {
		if strings.Contains(log.Name(), "Spider.log.") {
			logName = log.Name()
		}
	}
	fmt.Println(logName)
	str , err := ioutil.ReadFile("./log/"+logName)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{
			"status":"ok",
			"error":err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status":"ok",
			"content":string(str),
		})
	}
}
