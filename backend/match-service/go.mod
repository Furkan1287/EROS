module eros/match-service

go 1.24.4

require (
	eros/shared v0.0.0
	github.com/gorilla/mux v1.8.1
	github.com/joho/godotenv v1.5.1
	github.com/mattn/go-sqlite3 v1.14.28
)

replace eros/shared => ../shared
