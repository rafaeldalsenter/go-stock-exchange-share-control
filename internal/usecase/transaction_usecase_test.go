package usecase

import (
	"errors"
	"go-stock-exchange-shares-control/internal/domain/mocks"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	code     = "TestCode"
	date     = time.Now()
	quantity = 1.4
	value    = 2.3
	tax      = 12.0
	useCase  = transactionUseCase{}
)

func TestNewPurchaseWithDomainError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("New", mock.Anything, mock.Anything).Return("", nil)
	useCase.transactionRepository = repositoryMock

	_, err := useCase.NewPurchase(ctx, "", date, quantity, value, tax)

	assert.Error(err)
}

func TestNewPurchaseWithRepositoryError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("New", mock.Anything, mock.Anything).Return("", errors.New("Unexpected"))
	useCase.transactionRepository = repositoryMock

	_, err := useCase.NewPurchase(ctx, code, date, quantity, value, tax)

	assert.Error(err)
}

func TestNewPurchaseWithSuccess(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("New", mock.Anything, mock.Anything).Return("", nil)
	useCase.transactionRepository = repositoryMock

	_, err := useCase.NewPurchase(ctx, code, date, quantity, value, tax)

	assert.NoError(err)
}

func TestNewSaleWithDomainError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("New", mock.Anything, mock.Anything).Return("", nil)
	useCase.transactionRepository = repositoryMock

	_, err := useCase.NewSale(ctx, "", date, quantity, value, tax)

	assert.Error(err)
}

func TestNewSaleWithRepositoryError(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("New", mock.Anything, mock.Anything).Return("", errors.New("Unexpected"))
	useCase.transactionRepository = repositoryMock

	_, err := useCase.NewSale(ctx, code, date, quantity, value, tax)

	assert.Error(err)
}

func TestNewSaleWithSuccess(t *testing.T) {
	assert := assert.New(t)

	repositoryMock := new(mocks.TransactionRepositoryMock)
	repositoryMock.On("New", mock.Anything, mock.Anything).Return("", nil)
	useCase.transactionRepository = repositoryMock

	_, err := useCase.NewSale(ctx, code, date, quantity, value, tax)

	assert.NoError(err)
}
