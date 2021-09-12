package main

import "time"

const (
	// Gin Config
	DEFAULT_PORT uint16 = 7000

	// Reloading interval, used for: urlMapping auto reload from database
	RELOAD_INTERVAL time.Duration = 5 * time.Minute

	// URL Redirector
	REDIRECT_UNRECOGNIZED_REQUEST_TO string = "https://gaukas.wang"
)

var (
	globalConf config = config{
		Web: webserverConf{
			Port:    DEFAULT_PORT,
			Release: false,
		},
		Mysql: dbConf{},
	}

	useMysql bool = false
)
