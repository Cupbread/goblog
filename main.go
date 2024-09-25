package main

import (
	"database/sql"
	"github.com/gorilla/mux"
	"goblog/app/http/middlewares"
	"goblog/bootstrap"
	"goblog/pkg/database"
	"goblog/pkg/logger"
	"net/http"
)

var router = mux.NewRouter()
var db *sql.DB

func main() {

	database.InitDB()
	db = database.DB

	bootstrap.SetupDB()
	router = bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}
