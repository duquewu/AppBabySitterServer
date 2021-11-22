package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"os/exec"
	"strings"
)

type StatusBarHideInfo struct {
	Hide bool `json:"hide"`
}

func StatusBarGet(c *gin.Context) {
	command := "settings get system systembar_hide"
	var cmd = exec.Command("/data/data/com.termux/files/usr/bin/su", "-c", command)
	result, err := cmd.Output()
	resultStr := strings.TrimSpace(string(result))
	fmt.Println("result: " + resultStr)
	if err != nil {
		panic(err)
		return
	}
	c.JSON(200, Response{Code: 200, Message: "success", Data: StatusBarHideInfo{
		Hide: resultStr == "1",
	}})
}

func StatusBarPost(c *gin.Context) {
	showStatusBar := c.Query("hide")
	log.Printf("%v", &showStatusBar)
	var command string
	if showStatusBar == "true" {
		command = `settings put system systembar_hide 1 && am broadcast -a com.tchip.changeBarHideStatus`
	} else {
		command = `settings put system systembar_hide 0 && am broadcast -a com.tchip.changeBarHideStatus`
	}
	var cmd = exec.Command("/data/data/com.termux/files/usr/bin/su", "-c", command)
	_, err := cmd.Output()
	if err != nil {
		panic(err)
		return
	}
	StatusBarGet(c)
}
