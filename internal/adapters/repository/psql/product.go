package psql

import (
	"context"
	"github.com/iBoBoTi/gollet-api/internal/core/domain"
	"github.com/iBoBoTi/gollet-api/internal/core/ports"
	"github.com/jackc/pgx/v4/pgxpool"
	"math"
)

type productRepository struct {
	db *pgxpool.Pool
}

func NewProductRepository(db *pgxpool.Pool) ports.ProductRepository {
	return &productRepository{
		db: db,
	}
}

func (r *productRepository) GetProductByID(id int64) (*domain.Product, error) {
	product := domain.Product{}
	queryString := `SELECT * FROM product WHERE id = $1`
	err := r.db.QueryRow(context.Background(), queryString, id).Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *productRepository) GetProductsList(paginator *domain.Paginate) (*domain.Paginate, error) {
	products := make([]domain.Product, 0)
	queryString := `SELECT * FROM product ORDER BY created_at ASC LIMIT $1 OFFSET $2`
	rows, err := r.db.Query(context.Background(), queryString, paginator.Limit, paginator.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		product := domain.Product{}
		err := rows.Scan(&product.ID, &product.Name, &product.Price, &product.CreatedAt, &product.UpdatedAt)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}

	count := 0
	_ = r.db.QueryRow(context.Background(), `SELECT COUNT(*) FROM product`).Scan(&count)
	paginator.TotalRows = count
	paginator.Rows = products
	paginator.TotalPages = int(math.Ceil(float64(count) / float64(paginator.Limit)))

	return paginator, nil
}

func (r *productRepository) CreateProduct(product *domain.Product) (*domain.Product, error) {
	queryString := `INSERT INTO product (name, price) VALUES ($1, $2) RETURNING *`
	result := &domain.Product{}
	row := r.db.QueryRow(context.Background(), queryString, product.Name, product.Price)
	err := row.Scan(&result.ID, &result.Name, &result.Price, &result.CreatedAt, &result.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return result, nil
}
