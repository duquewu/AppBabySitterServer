package router

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runtime/debug"
)

func Recover(c *gin.Context) {
	defer func() {
		if r := recover(); r != nil {
			log.Printf("panic: %v\n", r)
			debug.PrintStack()
			c.JSON(http.StatusOK, Response{
				Code:    500,
				Message: ErrorToString(r),
				Data:    nil,
			})
			c.Abort()
		}
	}()
	c.Next()
}

func ErrorToString(r interface{}) string {
	switch v := r.(type) {
	case error:
		return v.Error()
	default:
		return r.(string)
	}
}
