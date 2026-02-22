package handler

import (
	"net/http"

	"rpractice/internal/usecase"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	useCase usecase.UserInterface
}

func NewHandler(v1 *gin.RouterGroup, useCase usecase.UserInterface) {
	handler := &Handler{useCase: useCase}

	v1.POST("/create", handler.CreateUser)
}

func (h *Handler) CreateUser(c *gin.Context) {
	var input struct {
		Name string `json:"name"`
	}

	if err := c.BindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp := h.useCase.CreateUser(input.Name)
	c.JSON(http.StatusOK, gin.H{"Response": resp})
}
