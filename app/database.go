package app

import (
	"belajar-golang-restful-api/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("mysql", "root@tcp(localhost:3306)/golang_database_migration")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
	//	create table category migrate
	//	create -ext sql -dir db/migrations create_category_table

	//	run database migration
	//	migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migrations up

	//	run or drop migration some version
	//	migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migrations up/down {sum version}

	//	check version migration
	//	migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migrations version

	//	change version migration
	//	migrate -database "mysql://root@tcp(localhost:3306)/golang_database_migration" -path db/migrations force {sum version}
}
