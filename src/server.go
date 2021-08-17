package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, REDIRECT_UNRECOGNIZED_REQUEST_TO)
	})

	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})

	router.GET("/:name", func(c *gin.Context) {
		name := strings.ToLower(c.Param("name"))

		if redirectUrl, ok := urlRedirectMap[name]; ok {
			c.Redirect(http.StatusMovedPermanently, redirectUrl)
		} else if responseTxt, ok := urlTextMap[name]; ok {
			c.String(http.StatusOK, responseTxt)
		} else {
			c.Redirect(http.StatusFound, REDIRECT_UNRECOGNIZED_REQUEST_TO)
		}
	})
	router.Run(fmt.Sprintf(":%d", PORT_TO_USE))
}
