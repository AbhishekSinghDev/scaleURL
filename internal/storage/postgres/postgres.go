package postgres

import (
	"context"
	"os"

	"github.com/AbhishekSinghDev/scaleURL/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

type Postgres struct {
	Db *pgxpool.Pool
}

func New(cfg *config.Env) *Postgres {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, cfg.Variables.DatabaseUrl)
	if err != nil {
		log.Error().AnErr("error", err).Msg("failed to connect database")
		os.Exit(1)
	}

	return &Postgres{
		Db: pool,
	}
}
