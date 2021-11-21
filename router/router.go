package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Route is the information for every URI.
type Route struct {
	// Name is the name of this Route.
	Name string
	// Method is the string for the HTTP method. ex) GET, POST etc..
	Method string
	// Pattern is the pattern of the URI.
	Pattern string
	// HandlerFunc is the handler function of this route.
	HandlerFunc gin.HandlerFunc
}

var Routes = []Route{
	{
		"AdbGet",
		http.MethodGet,
		"/api/adb",
		AdbGet,
	},
	{
		"AdbPost",
		http.MethodPost,
		"api/adb",
		AdbPost,
	},
}
