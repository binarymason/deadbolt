package config

import (
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Port        string   `yaml:"port"`
	Secret      string   `yaml:"deadbolt_secret"`
	Whitelisted []string `yaml:"whitelisted_clients"`
}

func Load(p string) Config {
	c := Config{}
	yamlFile, err := ioutil.ReadFile(p)
	if err != nil {
		panic(fmt.Sprintf("yamlFile.Get err   #%v ", err))
	}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(fmt.Sprintf("Unmarshal: %v", err))
	}

	// DEADBOLT_SECRET takes precedence over config file
	s := os.Getenv("DEADBOLT_SECRET")
	if s != "" {
		c.Secret = s
	}
	if c.Secret == "" {
		panic("Fatal: deadbolt secret not in environment or config file.")
	}

	return c
}
