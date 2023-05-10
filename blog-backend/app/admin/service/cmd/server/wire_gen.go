// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/registry"
	"kratos-blog/app/admin/service/internal/biz"
	"kratos-blog/app/admin/service/internal/data"
	"kratos-blog/app/admin/service/internal/server"
	"kratos-blog/app/admin/service/internal/service"
	"kratos-blog/gen/api/go/common/conf"
)

// Injectors from wire.go:

// initApp init kratos application.
func initApp(logger log.Logger, registrar registry.Registrar, bootstrap *conf.Bootstrap) (*kratos.App, func(), error) {
	authenticator := data.NewAuthenticator(bootstrap)
	engine := data.NewAuthorizer()
	discovery := data.NewDiscovery(bootstrap)
	userServiceClient := data.NewUserServiceClient(discovery, bootstrap)
	client := data.NewRedisClient(bootstrap, logger)
	attachmentServiceClient := data.NewAttachmentServiceClient(discovery, bootstrap)
	commentServiceClient := data.NewCommentServiceClient(discovery, bootstrap)
	categoryServiceClient := data.NewCategoryServiceClient(discovery, bootstrap)
	linkServiceClient := data.NewLinkServiceClient(discovery, bootstrap)
	postServiceClient := data.NewPostServiceClient(discovery, bootstrap)
	tagServiceClient := data.NewTagServiceClient(discovery, bootstrap)
	dataData, cleanup, err := data.NewData(client, authenticator, engine, userServiceClient, attachmentServiceClient, commentServiceClient, categoryServiceClient, linkServiceClient, postServiceClient, tagServiceClient, logger)
	if err != nil {
		return nil, nil, err
	}
	userTokenRepo := data.NewUserTokenRepo(dataData, authenticator, logger)
	userTokenUseCase := biz.NewUserAuthUseCase(userTokenRepo)
	authenticationService := service.NewAuthenticationService(logger, userServiceClient, userTokenUseCase)
	userService := service.NewUserService(logger, userServiceClient)
	postService := service.NewPostService(logger, postServiceClient)
	linkService := service.NewLinkService(logger, linkServiceClient)
	categoryService := service.NewCategoryService(logger, categoryServiceClient)
	commentService := service.NewCommentService(logger, commentServiceClient)
	tagService := service.NewTagService(logger, tagServiceClient)
	attachmentService := service.NewAttachmentService(logger, attachmentServiceClient)
	httpServer := server.NewHTTPServer(bootstrap, logger, authenticator, engine, authenticationService, userService, postService, linkService, categoryService, commentService, tagService, attachmentService)
	app := newApp(logger, registrar, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
