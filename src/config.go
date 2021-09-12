package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type config struct {
	web   webserverConf `yaml:"Webserver"`
	mysql dbConf        `yaml:"Database"`
}

func loadConf(filepath string, conf interface{}) error {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return err
	}
	err = yaml.Unmarshal(content, conf)
	return err
}

type webserverConf struct {
	port    uint16 `yaml:"port"`
	release bool   `yaml:"GIN-release"`
}

// Only for MySQL
type dbConf struct {
	user           string `yaml:"user"`
	passwd         string `yaml:"passwd"`
	host           string `yaml:"host"`
	port           uint16 `yaml:"port"`
	database       string `yaml:"database"`
	serverCAPath   string `yaml:"server-CA"`
	clientKeyPath  string `yaml:"client-key"`
	clientCertPath string `yaml:"client-cert"`
}
