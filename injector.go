//go:build wireinject
// +build wireinject

package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"golang-restful-api/app"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/controller"
	"golang-restful-api/service"
)

var categorySet = wire.NewSet(
	repository.NewCategoryRepository, 
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService),new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CatgeoryController), new(*controller.CatgeoryControllerImpl)),
)

func InitilizedServer() *http.Server {
	wire.Build(
		app.NewDB,
		validator.New,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}