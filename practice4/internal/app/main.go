package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/swaggo/http-swagger"
	"log"
	"net/http"
	"os"
	"practice4/internal/handler"
	"practice4/internal/middleware"
	"practice4/internal/repository"
	"practice4/internal/repository/_postgres"
	"practice4/internal/usecase"
	"time"

	"practice4/pkg/modules"
)

func Run() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Println("No .env file found")
	}
	ctx := context.Background()
	//ctx, cancel := context.WithCancel(context.Background())
	//defer cancel()
	//
	dbConfig := initPostgreConfig()

	_postgre := _postgres.NewPGXDialect(ctx, dbConfig)

	repositories := repository.NewRepositories(_postgre)

	userUsecase := usecase.NewUserUsecase(repositories.UserRepository)

	userHandler := handler.NewHandler(userUsecase)

	mux := http.NewServeMux()

	mux.Handle("/swagger/", httpSwagger.WrapHandler)

	api := http.NewServeMux()
	api.HandleFunc("/users", userHandler.GetUsers)
	api.HandleFunc("/user", userHandler.GetUserByID)
	api.HandleFunc("/users/create", userHandler.CreateUser)
	api.HandleFunc("/users/update", userHandler.UpdateUser)
	api.HandleFunc("/users/delete", userHandler.DeleteUser)
	api.HandleFunc("/health", userHandler.Health)

	protected := middleware.Logging(
		middleware.APIKey(api),
	)

	mux.Handle("/", protected)
	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", mux))

	//users, err := repositories.GetUsers()
	//if err != nil {
	//	fmt.Printf("Error fetching users: %v\n", err)
	//	return
	//}
	//
	//fmt.Printf("Users: %+v\n", users)
}

func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        os.Getenv("DB_HOST"),
		Port:        os.Getenv("DB_PORT"),
		Username:    os.Getenv("DB_USER"),
		Password:    os.Getenv("DB_PASSWORD"),
		DBName:      os.Getenv("DB_NAME"),
		SSLMode:     os.Getenv("DB_SSLMODE"),
		ExecTimeout: 5 * time.Second,
	}
}
