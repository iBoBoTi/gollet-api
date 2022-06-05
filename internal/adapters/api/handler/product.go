package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/iBoBoTi/gollet-api/internal/adapters/api/response"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"github.com/jackc/pgx/v4"
	"net/http"
	"strconv"
)

type productHandler struct {
	productService ports.ProductService
}

func NewProductHandler(productService ports.ProductService) ports.ProductHandler {
	return &productHandler{
		productService: productService,
	}
}

func (h *productHandler) GetProductByID(c *gin.Context) {
	roleID := c.Param("id")
	id, err := strconv.Atoi(roleID)
	if err != nil {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
	}

	role, err := h.productService.GetProductByID(int64(id))
	fmt.Println(pgx.ErrNoRows)
	if err != nil {
		switch err == pgx.ErrNoRows {
		case true:
			response.JSON(c, "product not found", http.StatusNotFound, nil, []string{err.Error()})
			return
		default:
			response.JSON(c, "failed to fetch product", http.StatusInternalServerError, nil, []string{err.Error()})
			return
		}
	}

	response.JSON(c, "success finding role", http.StatusOK, role, nil)
}

func (h *productHandler) GetProductsList(c *gin.Context) {
	p := c.Query("page")
	if p == "" || p == "0" {
		p = "1"
	}
	page, err := strconv.Atoi(p)
	if err != nil {
		response.JSON(c, "invalid_request", http.StatusBadRequest, nil, nil)
	}
	paginatedProducts, err := h.productService.GetProductsList(page)
	if err != nil {
		response.JSON(c, "failed to find products", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}

	response.JSON(c, "success retrieving products", http.StatusOK, paginatedProducts, nil)
}

func (h *productHandler) CreateProduct(c *gin.Context) {
	var product domain.Product
	if err := c.ShouldBindJSON(&product); err != nil {
		response.JSON(c, "invalid_request_body", http.StatusBadRequest, nil, []string{err.Error()})
		return
	}

	resultProduct, err := h.productService.CreateProduct(&product)
	if err != nil {
		response.JSON(c, "failed to create product", http.StatusInternalServerError, nil, []string{err.Error()})
		return
	}
	response.JSON(c, "success creating product", http.StatusCreated, resultProduct, nil)

}
