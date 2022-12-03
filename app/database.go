package app

import (
	"database/sql"
	"golang-restful-api/helper"
	"time"
)

func NewDB() *sql.DB{
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/belajar_golang_restful_api_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
	//migrate create -ext sql -dir db/migration create_table_first
		//migrate create -ext sql -dir db/migration create_table_second
				//migrate create -ext sql -dir db/migration create_table_third
	//migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_restful_api_migration" -path db/migration up
		//migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_restful_api_migration" -path db/migration down
		//migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_restful_api_migration" -path db/migration version
		//migrate -database "mysql://root:root@tcp(localhost:3306)/belajar_golang_restful_api_migration" -path db/migration force [version when database not dirty]
}