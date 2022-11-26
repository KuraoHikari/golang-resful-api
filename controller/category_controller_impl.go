package controller

import (
	"golang-restful-api/helper"
	"golang-restful-api/model/web"
	"golang-restful-api/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CatgeoryControllerImpl struct {
	CategoryService service.CategoryService
}

func NewCategoryController(categoryService service.CategoryService) CatgeoryController{
	return &CatgeoryControllerImpl{
		CategoryService: categoryService,
	}
}

//*http.Request karena struct maka perlu pointer, Tapi httprouter.Params karena interface tidak perlu pointer
func (controller *CatgeoryControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryCreateRequest := web.CategoryCreateRequest{}
	helper.ReadFromRequestBody(request, &categoryCreateRequest)


	categoryResponse := controller.CategoryService.Create(request.Context(), categoryCreateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CatgeoryControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryUpdateRequest := web.CategoryUpdateRequest{}
	helper.ReadFromRequestBody(request, &categoryUpdateRequest)

	categoryId := params.ByName("categoryId")
	id,err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)
	categoryUpdateRequest.Id = id

	categoryResponse := controller.CategoryService.Update(request.Context(), categoryUpdateRequest)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CatgeoryControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryId := params.ByName("categoryId")
	id,err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	controller.CategoryService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CatgeoryControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryId := params.ByName("categoryId")
	id,err := strconv.Atoi(categoryId)
	helper.PanicIfError(err)

	categoryResponse := controller.CategoryService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
func (controller *CatgeoryControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params){
	categoryResponses := controller.CategoryService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code: 200,
		Status: "OK",
		Data: categoryResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}