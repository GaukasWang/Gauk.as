package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func badRequest(c *gin.Context) {
	// Alpha test: hardcode behaviour
	c.Status(http.StatusBadRequest)
}

func setRoute(router *gin.Engine) {
	// Serve Static Files
	router.Static("/assets", "./res/assets")
	router.StaticFile("/favicon.ico", "./res/assets/favicon.ico") // Emoji for cocktail (GIN)

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/", func(c *gin.Context) {
		c.Redirect(301, REDIRECT_UNRECOGNIZED_REQUEST_TO)
	})

	router.GET("/test", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "This is a test message: Hello %s %s", firstname, lastname)
	})

	router.GET("/hash", func(c *gin.Context) {
		alg := c.DefaultQuery("alg", "SHA-256")    // SHA-256, SHA-224, MD5(don't!)
		format := c.DefaultQuery("format", "json") // json, text
		msg := c.Query("msg")                      // the actual message to compute hash on

		hashResult := computeHash(msg, alg)

		switch format {
		case "text":
			c.String(http.StatusOK, hashResult.String())
		case "json":
			c.JSON(http.StatusOK, hashResult.GinJSON())
		default: // JSON
			c.String(http.StatusBadRequest, "UNSUPPORTED FORMAT")
		}
	})

	router.GET("/:name", func(c *gin.Context) {
		name := strings.ToLower(c.Param("name"))

		if redirectUrl, ok := urlRedirect(name); ok {
			c.Redirect(http.StatusMovedPermanently, redirectUrl)
			// } else if responseFilepath, ok := LoadTXTFile(name); ok {
			// 	c.File(responseFilepath)
		} else if responseTxt, ok := urlLoadFile(name); ok {
			c.String(http.StatusOK, responseTxt)
		} else {
			badRequest(c)
		}
	})
}
