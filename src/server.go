package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	setRoute(router)

	router.Run(fmt.Sprintf(":%d", PORT_TO_USE))
}
