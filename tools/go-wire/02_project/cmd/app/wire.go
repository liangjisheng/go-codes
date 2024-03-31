//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"gowire/02_project/internal"
	"gowire/02_project/internal/service"
)

func NewPostService() (*service.PostService, error) {
	panic(wire.Build(internal.Set))
}
