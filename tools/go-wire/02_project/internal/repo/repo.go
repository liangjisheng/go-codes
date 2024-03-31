package repo

import "gowire/02_project/internal/domain"

// NewPostRepo NewPostRepo
func NewPostRepo() (domain.IPostRepo, error) {
	return new(domain.IPostRepo), nil
}
