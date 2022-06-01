package domain

type Entry struct {
	BaseModel
	WalletID  int64  `json:"wallet_id"`
	EntryType string `json:"entry_type"`
	Amount    int64  `json:"amount"`
}
