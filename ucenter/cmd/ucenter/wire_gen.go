// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"ucenter/internal/biz"
	"ucenter/internal/conf"
	"ucenter/internal/data"
	"ucenter/internal/server"
	"ucenter/internal/service"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, error) {
	dataData, err := data.NewData(confData)
	if err != nil {
		return nil, err
	}
	memberRepo := data.NewMemberRepo(dataData, logger)
	memberMgr := biz.NewMemberMgr(memberRepo, logger)
	authService := service.NewAuthService(memberMgr, logger)
	httpServer := server.NewHTTPServer(confServer, authService)
	grpcServer := server.NewGRPCServer(confServer, authService)
	app := newApp(logger, httpServer, grpcServer)
	return app, nil
}
