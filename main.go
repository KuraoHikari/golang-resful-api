package main

import (
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService :=service.NewCategoryService(categoryRepository,db, validate)
	catgeoryController := controller.NewCategoryController(categoryService)

	router:= httprouter.New()

	router.GET("/api/categories", catgeoryController.FindAll)
	router.GET("/api/categories/:categoryId", catgeoryController.FindById)
	router.POST("/api/categories", catgeoryController.Create)
	router.PUT("/api/categories/:categoryId", catgeoryController.Update)
	router.DELETE("/api/categories/:categoryId", catgeoryController.Delete)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}