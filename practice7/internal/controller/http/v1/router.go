package v1

import (
	"practice7/internal/usecase"

	"github.com/gin-gonic/gin"
	"practice7/utils"
)

func NewUserRoutes(handler *gin.RouterGroup, t usecase.UserInterface, l interface{}) {
	r := &userRoutes{t, l}

	h := handler.Group("/users")
	{
		h.POST("/", r.RegisterUser)
		h.POST("/login", r.LoginUser)

		protected := h.Group("/")
		protected.Use(utils.JWTAuthMiddleware(), utils.RateLimiter())
		{
			protected.GET("/me", r.GetMe)

			admin := protected.Group("/")
			admin.Use(utils.RoleMiddleware("admin"))
			{
				admin.PATCH("/promote/:id", r.PromoteUser)
				admin.GET("/users", r.GetAllUsers)
			}
		}
	}
}
