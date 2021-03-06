// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"home/app/captcha/service/internal/biz"
	"home/app/captcha/service/internal/conf"
	"home/app/captcha/service/internal/data"
	"home/app/captcha/service/internal/server"
	"home/app/captcha/service/internal/service"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	captchaRepo := data.NewCaptchaRepo(dataData, logger)
	greeterUsecase := biz.NewCaptchaUseCase(captchaRepo, logger)
	greeterService := service.NewCaptchaService(greeterUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, greeterService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
