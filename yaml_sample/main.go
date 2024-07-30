package main

import (
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	A string `yaml:"a"`
	B struct {
		C int
		D []int
	}
}

func main() {
	f, _ := os.Open("config.yml")

	var cfg Config
	yaml.NewDecoder(f).Decode(&cfg)

	println(cfg.A)
	println(cfg.B.C)
	for _, v := range cfg.B.D {
		println(v)
	}
}
