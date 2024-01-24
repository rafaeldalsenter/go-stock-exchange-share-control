package domain

import (
	"errors"
	"math"
)

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

func (s *Stock) averagePrice(transactionType TransactionType) (float64, error) {
	sumQuantity := 0.0
	sumValue := 0.0

	for _, transaction := range s.Transactions {

		if transaction.Type != transactionType {
			continue
		}

		sumValue += (transaction.Quantity * transaction.Value) + transaction.Tax
		sumQuantity += transaction.Quantity
	}

	if sumValue == 0.0 {
		return 0.0, errors.New("There are no valid " + string(transactionType) + " records")
	}

	ratio := math.Pow(10, float64(2))
	return math.Round((sumValue/sumQuantity)*ratio) / ratio, nil
}

func (s *Stock) AveragePurchasePrice() (float64, error) {
	return s.averagePrice(Purchase)
}

func (s *Stock) AverageSellingPrice() (float64, error) {
	return s.averagePrice(Sale)
}

func (s *Stock) Profit() (float64, error) {
	// TODO calcular lucro até agora
	// rodar todas as vendas e ir calculando o lucro de cada uma (so considerando o estado até entao das compras)
	return 0, nil
}

type StockPriceUseCase interface {
	AveragePurchasePrice(code string) (float64, error)
	AverageSellingPrice(code string) (float64, error)
}
