package config

import "fmt"

type Config struct{}

func ReadConfig(path string) *Config {
	fmt.Println("Constructing Config")
	return &Config{}
}
