package logger

import (
	"os"

	"github.com/AbhishekSinghDev/scaleURL/internal/config"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Init() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	env := config.Get().Env
	log.Logger = zerolog.New(os.Stdout).With().Timestamp().Str("service", "scale-url").Str("env", env).Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}
