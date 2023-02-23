package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

func (c Config) String() string {
	return fmt.Sprintf("{reg: %v,grade: %v,log: %v,portal: %v}", c.Reg, c.Grade, c.Log, c.Portal)
}

type Config struct {
	Reg    ServiceConfig `json:"reg"`
	Grade  ServiceConfig `json:"grade"`
	Log    ServiceConfig `json:"log"`
	Portal ServiceConfig `json:"portal"`
}

func (c ServiceConfig) String() string {
	return fmt.Sprintf("{host: %s,port :%s}", c.Host, c.Port)
}

type ServiceConfig struct {
	Host string `json:"host" default:"localhost"`
	Port string `json:"port" default:"8080"`
}

func init() {
	//TODO how to use relative path
	s := "C:\\Users\\lqsgo\\Documents\\golandProjects\\distributed\\config\\distributed.yaml"
	bytes, err := os.ReadFile(s)
	if err != nil {
		fmt.Printf("read file %s error: %s", s, err.Error())
		return
	}
	config := Config{}
	err = yaml.Unmarshal(bytes, &config)
	if err != nil {
		fmt.Printf("resolve config error: %s", err.Error())
		return
	}
	fmt.Printf("config: %s", config)
}
