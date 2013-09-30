package isucon2

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	dir string    `json:"-"`
	Db  *DbConfig `json:"database"`
}

func LoadConfig(dir string) *Config {
	var c Config
	log.Println("Loading configuration")

	var env string
	if env = os.Getenv("ISUCON_ENV"); env == "" {
		env = "local"
	}

	if f, err := os.Open(dir + "common." + env + ".json"); err == nil {
		defer f.Close()
		json.NewDecoder(f).Decode(&c)
	} else {
		panic(err)
	}

	c.dir = dir
	return &c
}
