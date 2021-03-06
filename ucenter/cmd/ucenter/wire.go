// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"ucenter/internal/biz"
	"ucenter/internal/conf"
	"ucenter/internal/data"
	"ucenter/internal/server"
	"ucenter/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// initApp init kratos application.
func initApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
