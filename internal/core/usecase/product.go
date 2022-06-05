package usecase

import (
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"os"
	"strconv"
)

type productService struct {
	productRepo ports.ProductRepository
}

func NewProductService(productRepo ports.ProductRepository) ports.ProductService {
	return &productService{
		productRepo: productRepo,
	}
}

func (p *productService) GetProductByID(id int64) (*domain.Product, error) {
	return p.productRepo.GetProductByID(id)
}

func (p *productService) GetProductsList(page int) (*domain.Paginate, error) {
	limit, _ := strconv.Atoi(os.Getenv("PAGE_LIMIT"))

	paginate := domain.NewPaginate(limit, page)
	paginate.Offset = (page - 1) * limit
	return p.productRepo.GetProductsList(paginate)
}

func (p *productService) CreateProduct(product *domain.Product) (*domain.Product, error) {
	return p.productRepo.CreateProduct(product)
}
