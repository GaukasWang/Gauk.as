package main

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var testConf = `
web:
    port: 7001
    GIN-release: false
mysql:
    user: gaukas
    passwd: 8uS@mEpret?h
    host: localhost
`

type config struct {
	Web   webserverConf // `yaml:"webserver"`
	Mysql dbConf        // `yaml:"database"`
}

func loadConf(filepath string, conf interface{}) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, conf)
	if err != nil {
		fmt.Printf("loadConf err: %s\n", err)
	}

	str, err := yaml.Marshal(conf)
	if err == nil {
		fmt.Printf("loadConf loaded:\n%s\n", str)
	}

	return err
}

type webserverConf struct {
	Port    uint16 `yaml:"port"`
	Release bool   `yaml:"gin-release"`
}

// Only for MySQL
type dbConf struct {
	User           string `yaml:"user"`
	Passwd         string `yaml:"passwd"`
	Host           string `yaml:"host"`
	Port           uint16 `yaml:"port"`
	Database       string `yaml:"database"`
	ServerCAPath   string `yaml:"server-ca"`
	ClientKeyPath  string `yaml:"client-key"`
	ClientCertPath string `yaml:"client-cert"`
}
