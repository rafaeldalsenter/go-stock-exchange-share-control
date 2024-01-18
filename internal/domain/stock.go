package domain

import "errors"

type Stock struct {
	Code         string
	Transactions []Transaction
}

func NewStock(code string, transactions []Transaction) (*Stock, error) {

	if code == "" {
		return nil, errors.New("code is required")
	}

	if len(transactions) == 0 {
		return nil, errors.New("stock not have transactions")
	}

	return &Stock{
		Code:         code,
		Transactions: transactions,
	}, nil
}

func (s *Stock) AveragePurchasePrice() (float64, error) {
	return 0, nil
}

func (s *Stock) AverageSellingPrice() (float64, error) {
	return 0, nil
}

type StockAveragePriceUseCase interface {
	AveragePurchasePrice(code string) (float64, error)
	AverageSellingPrice(code string) (float64, error)
}
