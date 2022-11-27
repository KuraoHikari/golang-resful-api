package test

import (
	"database/sql"
	"encoding/json"
	"golang-restful-api/app"
	"golang-restful-api/controller"
	"golang-restful-api/helper"
	"golang-restful-api/middleware"
	"golang-restful-api/repository"
	"golang-restful-api/service"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"

	"github.com/go-playground/validator/v10"
)

func setupTestDB() *sql.DB {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restful_api_test")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}

func setupRouter(db *sql.DB) http.Handler{
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService :=service.NewCategoryService(categoryRepository,db, validate)
	catgeoryController := controller.NewCategoryController(categoryService)

	router:= app.NewRouter(catgeoryController)

	return middleware.NewAuthMiddleware(router)
}
func truncateCategory(db *sql.DB) {
	db.Exec("TRUNCATE category")
}


func TestCreateCategorySuccess(t *testing.T){
	db := setupTestDB()
	truncateCategory(db)
	router := setupRouter(db)

	requestBody := strings.NewReader(`{"name" : "Gadget"}`)
	request := httptest.NewRequest(http.MethodPost, "http://localhost:3000/api/categories", requestBody)
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("X-API-Key", "RAHASIA")

	recorder := httptest.NewRecorder()

	router.ServeHTTP(recorder, request)

	response := recorder.Result()
	assert.Equal(t, 200, response.StatusCode)

	body, _ := io.ReadAll(response.Body)
	var responseBody map[string]interface{}
	json.Unmarshal(body, &responseBody)

	assert.Equal(t, 200, int(responseBody["code"].(float64)))
	assert.Equal(t, "OK", responseBody["status"])
	assert.Equal(t, "Gadget", responseBody["data"].(map[string]interface{})["name"])
}
// func TestCreateCategoryFailed(t *testing.T){
	
// }
// func TestUpdateCategorySuccess(t *testing.T){
	
// }
// func TestUpdateCategoryFailed(t *testing.T){
	
// }

// func TestGetCategorySuccess(t *testing.T){
	
// }

// func TestGetCategoryFailed(t *testing.T){
	
// }

// func TestDeleteCategorySuccess(t *testing.T){
	
// }

// func TestDeleteCategoryFailed(t *testing.T){
	
// }
// func TestListCategorySuccess(t *testing.T){
	
// }
// func TestUnauthorized(t *testing.T){
	
// }