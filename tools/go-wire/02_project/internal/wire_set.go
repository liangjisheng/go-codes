package internal

import (
	"github.com/google/wire"
	"gowire/02_project/internal/repo"
	"gowire/02_project/internal/service"
	"gowire/02_project/internal/usecase"
)

//Set ProviderSet
var Set = wire.NewSet(
	wire.Struct(new(service.PostService), "*"),
	wire.Struct(new(usecase.PostUsecaseOption), "*"),
	usecase.NewPostUsecase,
	repo.NewPostRepo,
)
