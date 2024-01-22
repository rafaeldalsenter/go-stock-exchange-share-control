package usecase

import "go-stock-exchange-shares-control/internal/domain"

type stockPriceUseCase struct {
	transactionRepository domain.TransactionRepository
}

func NewStockPriceUseCase(transactionRepository domain.TransactionRepository) domain.StockPriceUseCase {
	return &stockPriceUseCase{
		transactionRepository: transactionRepository,
	}
}

func (st *stockPriceUseCase) AveragePurchasePrice(code string) (float64, error) {
	return 0.0, nil
}

func (st *stockPriceUseCase) AverageSellingPrice(code string) (float64, error) {
	return 0.0, nil
}
