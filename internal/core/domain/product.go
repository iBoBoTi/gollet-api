package domain

type Product struct {
	BaseModel
	Name  string `json:"name" binding:"required"`
	Price int64  `json:"price" binding:"required,min=1"`
}
