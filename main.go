package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	r "termux-tablet-dashboard/router"
)

func handleCommonError(code int, c *gin.Context) {
	res := r.Response{
		Code:    code,
		Message: "Unknown",
		Data:    nil,
	}
	switch code {
	case http.StatusNotFound:
		res.Message = "No Found"
	case http.StatusMethodNotAllowed:
		res.Message = "Method Not Allowed"
	default:
		res.Message = "UnKnown"
	}
	c.JSON(200, res)
}

func main() {

	var router = gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			//return origin == "https://github.com"
			return true
		},
	}))
	router.NoRoute(func(context *gin.Context) {
		handleCommonError(http.StatusNotFound, context)
	})
	router.NoMethod(func(context *gin.Context) {
		handleCommonError(http.StatusMethodNotAllowed, context)
	})
	router.Use(r.Recover)
	// ignore sensitive
	router.RedirectFixedPath = true
	// mapping GET, POST
	for _, route := range r.Routes {
		switch route.Method {
		case http.MethodGet:
			router.GET(route.Pattern, route.HandlerFunc)
		case http.MethodPost:
			router.POST(route.Pattern, route.HandlerFunc)
		}
	}
	// run in 0.0.0.0:8080
	log.Fatal(router.Run(":8081"))
}
