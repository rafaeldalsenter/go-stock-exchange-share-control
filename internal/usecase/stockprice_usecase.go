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
	transactions, err := st.transactionRepository.Get(code)
	if err != nil {
		return 0.0, err
	}

	stock, err := domain.NewStock(code, transactions)
	if err != nil {
		return 0.0, err
	}

	return stock.AveragePurchasePrice()
}

func (st *stockPriceUseCase) AverageSellingPrice(code string) (float64, error) {
	transactions, err := st.transactionRepository.Get(code)
	if err != nil {
		return 0.0, err
	}

	stock, err := domain.NewStock(code, transactions)
	if err != nil {
		return 0.0, err
	}

	return stock.AverageSellingPrice()
}
