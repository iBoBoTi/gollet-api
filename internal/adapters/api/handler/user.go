package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/gollet-api/internal/adapters/api/response"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"
)

type userHandler struct {
	userService ports.UserService
}

func NewUserHandler(userService ports.UserService) ports.UserHandler {
	return &userHandler{
		userService: userService,
	}
}

func (u *userHandler) SignUpUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		response.JSON(c, "invalid request body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	result, err := u.userService.SignUpUser(&user)
	if err != nil {
		if pgerr, ok := err.(*pgconn.PgError); ok {
			switch pgerr.Code {
			case "23505":
				response.JSON(c, "Error", http.StatusForbidden, nil, []string{err.Error()})
				return
			}
		}
		response.JSON(c, "Error", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	userResponse := domain.UserResponse{
		ID:        result.ID,
		Name:      result.Name,
		Email:     result.Email,
		CreatedAt: result.CreatedAt,
		UpdatedAt: result.UpdatedAt,
	}
	response.JSON(c, "Success", http.StatusOK, userResponse, nil)
}

func (u *userHandler) LoginUser(c *gin.Context) {
	var req domain.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.JSON(c, "invalid request body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	userResponse, err := u.userService.LoginUser(&req)
	if err != nil {
		switch err {
		case pgx.ErrNoRows:
			response.JSON(c, "user not found", http.StatusNotFound, nil, []string{err.Error()})
			return
		case domain.ErrInvalidPassword:
			response.JSON(c, "invalid password", http.StatusUnauthorized, nil, []string{err.Error()})
			return
		default:
			fmt.Println("here I am 500")
			response.JSON(c, "Error", http.StatusInternalServerError, nil, []string{err.Error()})
			return
		}
	}

	response.JSON(c, "Success", http.StatusOK, userResponse, nil)
}

func (u *userHandler) CreditUserWallet(c *gin.Context) {
	user := c.MustGet("user").(*domain.User)
	amt := c.Query("amount")
	if amt == "" || amt == "0" {
		amt = "0"
	}
	creditAmount, err := strconv.Atoi(amt)
	if err != nil {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	wallet, err := u.userService.CreditUserWallet(user.ID, int64(creditAmount))
	if err != nil {
		response.JSON(c, "Error", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	response.JSON(c, "Success", http.StatusOK, wallet, nil)
}

func (u *userHandler) DebitUserWallet(c *gin.Context) {
	user := c.MustGet("user").(*domain.User)
	amt := c.Query("amount")
	if amt == "" || amt == "0" {
		amt = "0"
	}
	debitAmount, err := strconv.Atoi(amt)
	if err != nil {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
		return
	}

	wallet, err := u.userService.DebitUserWallet(user.ID, int64(debitAmount))
	if err != nil {
		if err.Error() == "insufficient balance" {
			response.JSON(c, "insufficient balance", http.StatusForbidden, nil, []string{err.Error()})
			return
		}
		response.JSON(c, "Error", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	response.JSON(c, "Success", http.StatusOK, wallet, nil)
}
