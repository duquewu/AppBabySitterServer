package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"os/exec"
)

type CommandDTO struct {
	Command string `json:"command"`
}

func ExeCommandPost(c *gin.Context) {
	commandDTO := CommandDTO{}
	err := c.BindJSON(&commandDTO)
	if err != nil {
		panic(err)
		return
	}
	log.Printf("%v", &commandDTO)
	var cmd = exec.Command("/data/data/com.termux/files/usr/bin/sh", "-c", commandDTO.Command)
	result, err := cmd.Output()
	if err != nil {
		panic(err)
		return
	}
	c.JSON(200, Response{Code: 200, Message: "success", Data: string(result)})
}
