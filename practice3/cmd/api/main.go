package main

import (
	_ "practice3/docs"
	"practice3/internal/app"
)

// @title User API
// @version 1.0
// @description User CRUD API with middleware and pagination
// @host localhost:8080
// @BasePath /
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name X-API-KEY

func main() {
	app.Run()
}
