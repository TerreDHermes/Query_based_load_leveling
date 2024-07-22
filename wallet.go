package backend

type Wallets struct {
	ID      int     `gorm:"primaryKey" json:"id"`
	Balance float64 `json:"balance"`
}
