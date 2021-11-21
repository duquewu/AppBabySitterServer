package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os/exec"
	"strings"
)

type Pair struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func unpair(str string) Pair {
	fmt.Println("prepare unpair: " + str)
	str = strings.Trim(str, "\n")
	fmt.Println("trimmed: " + str)
	result := strings.Split(str, "]: [")
	fmt.Println(result)
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

func AdbGet(c *gin.Context) {
	var cmd = exec.Command("/data/data/com.termux/files/usr/bin/su", "-c", "getprop|grep service.adb.tcp.port")
	// 多返回值
	result, err := cmd.Output()
	if err != nil {
		fmt.Println(err)
		panic(err)
		return
	}
	c.JSON(http.StatusOK, Response{
		Code:    200,
		Message: "success",
		Data:    unpair(string(result)),
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
	// 多返回值
	_, err = cmd.Output()
	if err != nil {
		panic(err)
		return
	}
	AdbGet(c)

	//c.JSON(http.StatusOK, Response{
	//	Code:    http.StatusOK,
	//	Message: "Success",
	//	Data:    json,
	//})
	//var whoami []byte
	//var err error
	//var cmd *exec.Cmd
	//
	//// 执行单个shell命令时, 直接运行即可
	//cmd = exec.Command("/data/data/com.termux/files/usr/bin/su", "-c", "/data/data/com.termux/files/usr/bin/getprop")
	//if whoami, err = cmd.Output(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	//// 默认输出有一个换行
	//fmt.Println(string(whoami))
	//// 指定参数后过滤换行符
	//fmt.Println(strings.Trim(string(whoami), "\n"))
	//
	//fmt.Println("====")
}
