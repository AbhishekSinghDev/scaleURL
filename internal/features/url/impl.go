package url

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type repoImpl struct {
	db *pgxpool.Pool
}

func NewRepository(db *pgxpool.Pool) Repository {
	return &repoImpl{db: db}
}

func (r *repoImpl) Create(ctx context.Context, params CreateURLParams) (*Url, error) {
	return nil, nil
}

func (r *repoImpl) GetByShortCode(ctx context.Context, shortCode string) (*Url, error) {
	// TODO: SELECT * FROM urls WHERE short_code = $1
	return nil, nil
}

func (r *repoImpl) Delete(ctx context.Context, shortCode string) error {
	// TODO: DELETE FROM urls WHERE short_code = $1
	return nil
}
