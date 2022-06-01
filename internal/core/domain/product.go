package domain

type Product struct {
	BaseModel
	Name  string `json:"name"`
	Price int64  `json:"price"`
}
