package domain

import (
	"time"
)

type PaymentAccount struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UserID        string    `json:"user_id"`
	AccountType   string    `json:"account_type"`
	AccountNumber string    `json:"account_number"`
	Balance       float64   `json:"balance"`
	CurrencyId    int64     `json:"currency_id"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

type Transaction struct {
	ID                   uint      `gorm:"primaryKey"`
	AccountId            uint      `gorm:"not null"`
	Amount               float64   `gorm:"not null"`
	TransactionTimestamp time.Time `gorm:"default:current_timestamp"`
	ToAddress            string    `gorm:"not null"`
	Status               string    `gorm:"not null"` // e.g., 'pending', 'success', 'failed'
	Description          string
	CreatedAt            time.Time
}

type RecurringPayment struct {
	ID          uint      `gorm:"primaryKey"`
	AccountID   uint      `gorm:"not null"`
	Amount      float64   `gorm:"not null"`
	Interval    string    `gorm:"not null"` // e.g., 'daily', 'weekly', 'monthly'
	NextPayment time.Time `gorm:"not null"`
	CreatedAt   time.Time
}

type Currencies struct {
	ID           int64   `gorm:"primaryKey"`
	CurrencyCode string  `gorm:"not null"`
	ExchangeRate float64 `gorm:"not null"`
	CurrencyName string  `gorm:"not null"` // e.g., 'daily', 'weekly', 'monthly'
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
