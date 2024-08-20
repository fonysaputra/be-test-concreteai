package domain

type LoginRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

type RegisterRequest struct {
	Email    string `json:"email" example:"user@example.com"`
	Password string `json:"password" example:"password123"`
}

type PaymentAccountRequest struct {
	AccountType   string `json:"accountType" example:"debit"`
	AccountNumber string `json:"accountNumber" example:"2777625252525252" binding:"required"`
	Currancy      string `json:"currency" example:"USD"`
}

type TrxRequest struct {
	AccountNumber string `json:"accountNumber" example:"2777625252525252" binding:"required"`
	Amount        string `json:"amount" example:"5"`
	ToAddress     string `json:"to_address" example:"2777625252525252" binding:"required"`
}
