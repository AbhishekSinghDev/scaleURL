package config

import (
	"log"

	"github.com/AbhishekSinghDev/scaleURL/internal/types"
	"github.com/ilyakaznacheev/cleanenv"
)

type Env struct {
	types.HttpServer
	types.Variables
}

var cfg *Env

func MustLoad() {
	var env Env

	if err := cleanenv.ReadConfig(".env", &env); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	cfg = &env
}

func Get() *Env {
	if cfg == nil {
		log.Fatal("config not loaded, call MustLoad() first")
	}
	return cfg
}
