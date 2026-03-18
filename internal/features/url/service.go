package url

import "context"

type Service interface {
	Create(ctx context.Context, params CreateURLParams) (*Url, error)
	GetByShortCode(ctx context.Context, shortCode string) (*Url, error)
	Delete(ctx context.Context, shortCode string) error
}

type serviceImpl struct {
	repo Repository
}

func NewService(repo Repository) Service {
	return &serviceImpl{repo: repo}
}

func (s *serviceImpl) Create(ctx context.Context, params CreateURLParams) (*Url, error) {
	// TODO: validate params, generate short code, call s.repo.Create()
	return nil, nil
}

func (s *serviceImpl) GetByShortCode(ctx context.Context, shortCode string) (*Url, error) {
	// TODO: call s.repo.GetByShortCode()
	return nil, nil
}

func (s *serviceImpl) Delete(ctx context.Context, shortCode string) error {
	// TODO: call s.repo.Delete()
	return nil
}
