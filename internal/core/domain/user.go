package domain

type User struct {
	BaseModel
	Name           string `json:"name" binding:"required"`
	Email          string `json:"email" binding:"required, email"`
	Password       string `json:"password" binding:"required"`
	HashedPassword string `json:"-,omitempty"`
}
