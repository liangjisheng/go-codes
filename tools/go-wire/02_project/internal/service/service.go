package service

import "gowire/02_project/internal/domain"

// PostService PostService
type PostService struct {
	Usecase domain.IPostUsecase
}
