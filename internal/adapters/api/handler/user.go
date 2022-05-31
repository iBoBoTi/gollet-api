package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
)

type userHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) ports.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) CreateUser(c *gin.Context) {}

func (u *userHandler) LoginUser(c *gin.Context) {}
