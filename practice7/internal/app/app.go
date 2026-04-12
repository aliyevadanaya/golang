package app

import (
	"log"
	"practice7/internal/controller/http/v1"
	"practice7/internal/entity"
	"practice7/internal/usecase"
	"practice7/internal/usecase/repo"
	"practice7/pkg"

	"github.com/gin-gonic/gin"
)

func Run() {
	pg, err := pkg.New("host=localhost user=danaya_user password=DanayaKrasotka dbname=practice7 port=5432 sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	pg.Conn.AutoMigrate(&entity.User{})

	userRepo := repo.NewUserRepo(pg)
	userUseCase := usecase.NewUserUseCase(userRepo)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		v1.NewUserRoutes(api, userUseCase, nil)
	}

	r.Run(":8080")
}
