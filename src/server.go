package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func globalInit() {
	loadConf("./conf/config.yaml", &globalConf)

	if globalConf.mysql.host != "" {
		_, err := connectDB(globalConf.mysql)
		if err == nil {
			useMysql = true
		}
	}

	if err := initUrlMapping(); err != nil {
		fmt.Printf("Error in initUrlMapping: %s", err)
	}
}

func main() {
	router := gin.Default()

	globalInit()

	setRoute(router)
	fmt.Printf("Use MySQL: %t\n", useMysql)
	if globalConf.web.release {
		gin.SetMode(gin.ReleaseMode)
	}

	router.Run(fmt.Sprintf(":%d", globalConf.web.port))
}
