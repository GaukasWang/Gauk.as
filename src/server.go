package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func globalInit() {
	loadConf("./conf/config.yaml", &globalConf)

	if globalConf.Mysql.Host != "" {
		_, err := connectDB(globalConf.Mysql)
		if err == nil {
			useMysql = true
		}
	}

	if err := initUrlMapping(); err != nil {
		fmt.Printf("Error in initUrlMapping: %s", err)
	}
}

func main() {
	globalInit()

	fmt.Printf("Use MySQL: %t\n", useMysql)
	if globalConf.Web.Release {
		gin.SetMode(gin.ReleaseMode)
		fmt.Printf("Set Production: %t\n", globalConf.Web.Release)
	}

	router := gin.Default()
	setRoute(router)

	router.Run(fmt.Sprintf(":%d", globalConf.Web.Port))
}
