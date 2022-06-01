package domain

type Wallet struct {
	BaseModel
	UserID   int64  `json:"user_id"`
	Balance  int64  `json:"balance"`
	Currency string `json:"currency"`
}
