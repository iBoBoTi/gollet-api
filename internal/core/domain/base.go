package domain

import "time"

type BaseModel struct {
	ID        int64     `json:"id" binding:"required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
