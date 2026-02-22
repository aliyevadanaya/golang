package router

import (
	"classwork/internal/controller/http/handler"
	"classwork/internal/usecase"

	"github.com/gin-gonic/gin"
)

func NewRouter(useCase usecase.UserInterface) *gin.Engine {
	router := gin.Default()

	router.GET("/healthcheck", func(c *gin.Context) {
		c.JSON(200, gin.H{"healthcheck": "ok"})
	})

	v1 := router.Group("/api/v1")
	{
		handler.NewHandler(v1, useCase)
	}

	return router
}
