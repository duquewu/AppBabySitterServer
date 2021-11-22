package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os/exec"
)

func Reboot(c *gin.Context) {
	var cmd = exec.Command("/data/data/com.termux/files/usr/bin/su", "-c", "reboot")
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    result,
	})
}
