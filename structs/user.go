package structs

type User struct {
	FirstName     string  `json:"first_name"`
	LastName      string  `json:"last_name"`
	Email         string  `json:"email"`
	Contact       string  `json:"contact"`
	City          string  `json:"city"`
	ID            int     `json:"walletBalance"`
	WalletBalance float64 `json:"id"`
}
