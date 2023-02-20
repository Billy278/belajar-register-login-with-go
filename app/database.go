package app

import (
	"belajar-register-login-with-go/helper"
	"database/sql"
	"time"
)

func NewDB() *sql.DB {
	DB, err := sql.Open("mysql", "root:@tcp(localhost:3306)/registration_login_api_latihan1")
	helper.PanicIfError(err)
	DB.SetMaxIdleConns(10)
	DB.SetMaxOpenConns(20)
	DB.SetConnMaxIdleTime(10 * time.Minute)
	DB.SetConnMaxLifetime(60 * time.Minute)
	return DB
}
