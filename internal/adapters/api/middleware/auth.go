package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/gollet-api/internal/adapters/api/response"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"net/http"
	"strings"
)

const (
	authorizationHeaderKey = "authorization"
)

func AuthMiddleWare(findUserByEmail func(string) (*domain.User, error), tokenMaker ports.TokenMaker) gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.Request.Header.Get(authorizationHeaderKey)
		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is empty")
			respondAndAbort(c, "authorization header is empty", http.StatusUnauthorized, nil, []string{err.Error()})
			return
		}
		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("authorization header is invalid")
			respondAndAbort(c, "authorization header is invalid", http.StatusUnauthorized, nil, []string{err.Error()})
			return
		}
		authorizationType := strings.ToLower(fields[0])
		if authorizationType != "bearer" {
			err := errors.New("unsupported authorization type")
			respondAndAbort(c, "unsupported authorization type", http.StatusUnauthorized, nil, []string{err.Error()})
			return
		}
		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			respondAndAbort(c, "invalid token", http.StatusUnauthorized, nil, []string{err.Error()})
			return
		}
		// find user by email then add to context
		user, err := findUserByEmail(payload.Email)
		if err != nil {
			respondAndAbort(c, "user not found", http.StatusUnauthorized, nil, []string{err.Error()})
			return
		}
		c.Set("user", user)
		c.Next()
	}
}

func respondAndAbort(c *gin.Context, message string, status int, data interface{}, errs []string) {
	response.JSON(c, message, status, data, errs)
	c.Abort()
}
