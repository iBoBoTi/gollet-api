package ports

import "github.com/gin-gonic/gin"

type UserHandler interface {
	CreateUser(c *gin.Context)
	LoginUser(c *gin.Context)
}

type UserService interface {
	CreateUser()
	LoginUser()
}

type UserRepository interface {
	CreateUser()
	LoginUser()
}
