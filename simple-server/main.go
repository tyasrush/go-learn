package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"

	"gopkg.in/yaml.v2"
)

type (
	Config struct {
		Server Server `yaml:"server"`
	}

	Server struct {
		Port string `yaml:"port"`
	}

	LogFormat struct {
		Name  string
		Tag   string
		Msg   string
		ReqID string
	}
)

var configFilePath string

func parsePath() {
	flag.StringVar(&configFilePath, "cpath", "", "config file path")
	flag.Parse()
}

func main() {
	parsePath()
	configFile, err := ioutil.ReadFile(configFilePath)
	if err != nil {
		panic(err)
	}

	var cfg Config
	if err = yaml.Unmarshal(configFile, &cfg); err != nil {
		panic(err)
	}

	fmt.Printf("listen port %s\n", cfg.Server.Port)
	http.ListenAndServe(cfg.Server.Port, nil)
}
