package app

import (
	"golang-restful-api/controller"
	"golang-restful-api/exception"

	"github.com/julienschmidt/httprouter"
)

func NewRouter(catgeoryController controller.CatgeoryController) *httprouter.Router {
	router:= httprouter.New()

	router.GET("/api/categories", catgeoryController.FindAll)
	router.GET("/api/categories/:categoryId", catgeoryController.FindById)
	router.POST("/api/categories", catgeoryController.Create)
	router.PUT("/api/categories/:categoryId", catgeoryController.Update)
	router.DELETE("/api/categories/:categoryId", catgeoryController.Delete)

	router.PanicHandler = exception.ErrorHandler

	return router
}