package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
)

type ProductHandler interface {
	GetProductByID(c *gin.Context)
	GetProductsList(c *gin.Context)
	CreateProduct(c *gin.Context)
}

type ProductService interface {
	GetProductByID(id int64) (*domain.Product, error)
	GetProductsList(page int) (*domain.Paginate, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
}

type ProductRepository interface {
	GetProductByID(id int64) (*domain.Product, error)
	GetProductsList(paginator *domain.Paginate) (*domain.Paginate, error)
	CreateProduct(product *domain.Product) (*domain.Product, error)
}
