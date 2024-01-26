package usecase

import (
	"context"
	"errors"
	"go-stock-exchange-shares-control/internal/domain"
	"go-stock-exchange-shares-control/internal/domain/mocks"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	stockUseCase = stockPriceUseCase{}
	ctx          = context.TODO()
)

func TestAveragePurchasePriceWithRepositoryError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("Get", mock.Anything, mock.Anything).Return([]domain.Transaction{}, errors.New("Unexpected"))
	stockUseCase.transactionRepository = repositoryMock

	_, err := stockUseCase.AveragePurchasePrice(ctx, code)

	assert.Error(err)
}

func TestAveragePurchasePriceWithDomainError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("Get", mock.Anything, mock.Anything).Return([]domain.Transaction{}, nil)
	stockUseCase.transactionRepository = repositoryMock

	_, err := stockUseCase.AveragePurchasePrice(ctx, code)

	assert.Error(err)
}

func TestAveragePurchasePriceWithSuccess(t *testing.T) {
	assert := assert.New(t)

	var transactions = []domain.Transaction{
		{
			Quantity: 2,
			Value:    100,
			Type:     domain.Purchase,
		},
		{
			Quantity: 5,
			Value:    12,
			Type:     domain.Purchase,
		},
	}

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("Get", mock.Anything, mock.Anything).Return(transactions, nil)
	stockUseCase.transactionRepository = repositoryMock

	average, err := stockUseCase.AveragePurchasePrice(ctx, code)

	assert.NoError(err)
	assert.Equal(37.14, average)
}
