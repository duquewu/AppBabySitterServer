package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"strings"
	"unsafe"
)

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func unpair(str string) Pair {
	if str == "" {
		return Pair{}
	}
	fmt.Println("prepare unpair: " + str)
	str = strings.Trim(str, "\n")
	fmt.Println("trimmed: " + str)
	result := strings.Split(str, "]: [")
	fmt.Println(result)
	if len(result) <= 1 {
		return Pair{}
	}
	result[0] = strings.Replace(result[0], "[", "", -1)
	fmt.Println(result[0])
	result[1] = strings.Replace(result[1], "]", "", -1)
	fmt.Println(result[1])
	var key = strings.TrimSpace(result[0])
	var value = strings.TrimSpace(result[1])
	return Pair{
		Key:   key,
		Value: value,
	}
}

type Data struct {
	Info      Pair `json:"info,omitempty"`
	IsEnabled bool `json:"isEnabled"`
}

func AdbGet(c *gin.Context) {
	var cmd = exec.Command("/data/data/com.termux/files/usr/bin/sh", "-c", "getprop|grep service.adb.tcp.port")
	result, err := cmd.Output()
	// ExitCode 1 represent for grep not match, ignore this directly
	if err != nil && cmd.ProcessState.ExitCode() != 1 {
		panic(err)
		return
	}
	var keyValue = string(result)
	fmt.Printf("%s,%d\n", keyValue, unsafe.Sizeof(keyValue))
	var pair = unpair(keyValue)
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data: Data{
			IsEnabled: pair.Value == "5555",
			Info:      pair,
		},
	})
}

type AddPostRequestDto struct {
	IsAdbOpen bool `json:"isAdbOpen"`
}

func AdbPost(c *gin.Context) {
	json := AddPostRequestDto{}
	err := c.BindJSON(&json)
	if err != nil {
		panic(err)
		return
	}
	log.Printf("%v", &json)
	var command string
	if json.IsAdbOpen == true {
		command = `setprop service.adb.tcp.port 5555`
	} else {
		command = `setprop service.adb.tcp.port ""`
	}
	command += " && stop adbd && start adbd"
	var cmd = exec.Command("/data/data/com.termux/files/usr/bin/su", "-c", command)
	_, err = cmd.Output()
	if err != nil {
		panic(err)
		return
	}
	AdbGet(c)
}
