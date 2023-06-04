package usecase

import "gowire/02_project/internal/domain"

// PostUsecase PostUsecase
type PostUsecase struct {
	repo domain.IPostRepo
}

// PostUsecaseOption PostUsecaseOption
type PostUsecaseOption struct {
	Repo domain.IPostRepo
}

// NewPostUsecase NewPostUsecase
func NewPostUsecase(opt *PostUsecaseOption) (domain.IPostUsecase, error) {
	return &PostUsecase{repo: opt.Repo}, nil
}
