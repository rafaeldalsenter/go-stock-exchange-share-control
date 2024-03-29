package domain

import (
	"context"
	"errors"
	"time"
)

type TransactionType uint8

const (
	Purchase TransactionType = iota
	Sale
)

type Transaction struct {
	Code     string
	Date     time.Time
	Quantity float64
	Value    float64
	Tax      float64
	Type     TransactionType
}

func NewTransaction(t TransactionType, code string, date time.Time, quantity float64, value float64, tax float64) (*Transaction, error) {

	if code == "" {
		return nil, errors.New("code is required")
	}

	if quantity <= 0 {
		return nil, errors.New("quantity is invalid")
	}

	if value <= 0 {
		return nil, errors.New("value is invalid")
	}

	return &Transaction{
		Code:     code,
		Date:     date,
		Quantity: quantity,
		Value:    value,
		Tax:      tax,
		Type:     t,
	}, nil
}

type NewTransactionUseCase interface {
	NewSale(c context.Context, code string, date time.Time, quantity float64, value float64, tax float64) (string, error)
	NewPurchase(c context.Context, code string, date time.Time, quantity float64, value float64, tax float64) (string, error)
}

type TransactionRepository interface {
	New(c context.Context, transaction *Transaction) (string, error)
	Get(c context.Context, code string) ([]Transaction, error)
}
