// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import "github.com/google/wire"

// <-- 第二个参数设置成 error
func InitializeClient(config Config) (*Service, error) {
	wire.Build(NewService, NewAPIClient)
	return nil, nil
}
