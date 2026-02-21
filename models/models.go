package models

// Represents a user's net balance
type Balance struct {
	UserID string
	Amount float64
}

// Represents the final optimized transactions
type Transaction struct {
	From   string  `json:"from"`
	To     string  `json:"to"`
	Amount float64 `json:"amount"`
}
