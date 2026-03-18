package url

import "context"

type Repository interface {
	Create(ctx context.Context, params CreateURLParams) (*Url, error)
	GetByShortCode(ctx context.Context, shortCode string) (*Url, error)
	Delete(ctx context.Context, shortCode string) error
}
