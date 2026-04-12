package v1

import (
	"practice7/internal/entity"
	"practice7/internal/usecase"
	"practice7/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type userRoutes struct {
	t usecase.UserInterface
	l interface{}
}

func (r *userRoutes) RegisterUser(c *gin.Context) {
	var dto entity.CreateUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	hash, _ := utils.HashPassword(dto.Password)

	user := entity.User{
		ID:       uuid.New(),
		Username: dto.Username,
		Email:    dto.Email,
		Password: hash,
		Role:     "user",
	}

	created, _, err := r.t.RegisterUser(&user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(201, created)
}

func (r *userRoutes) LoginUser(c *gin.Context) {
	var dto entity.LoginUserDTO

	if err := c.ShouldBindJSON(&dto); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	token, err := r.t.LoginUser(&dto)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})
}

func (r *userRoutes) GetMe(c *gin.Context) {
	userID, _ := c.Get("userID")

	user, err := r.t.GetMe(userID.(string))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, user)
}

func (r *userRoutes) PromoteUser(c *gin.Context) {
	id := c.Param("id")

	err := r.t.PromoteUser(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "promoted"})
}

func (r *userRoutes) GetAllUsers(c *gin.Context) {
	users, err := r.t.GetAllUsers()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, users)
}
