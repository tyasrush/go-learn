package config

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type TestingStructLagi struct {
	TestingName  string
	TestingName1 string
	TestingName2 string
}

type DbConfig struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	User   string `yaml:"user"`
	Pwd    string `yaml:"pwd"`
	DBName string `yaml:"db"`
}

type Config struct {
	Db DbParentConfig `yaml:"db"`
}

type DbParentConfig struct {
	PgConfig DbProp `yaml:"postgres"`
}

type DbProp struct {
	Master DbConfig `yaml:"master"`
	Slave  DbConfig `yaml:"slave"`
}

func GetPgConfig(path string) Config {
	var cfg Config
	cleanPath := filepath.Clean(path)
	yamlFile, err := ioutil.ReadFile(cleanPath)
	if err != nil {
		log.Printf("yamlFile.Get err #%v ", err)
	}

	if err = yaml.Unmarshal(yamlFile, &cfg); err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	return cfg
}
