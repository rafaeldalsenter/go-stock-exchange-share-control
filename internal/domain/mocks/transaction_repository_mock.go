package mocks

import (
	"go-stock-exchange-shares-control/internal/domain"

	"github.com/stretchr/testify/mock"
)

type TransactionRepositoryMock struct {
	mock.Mock
}

func (r *TransactionRepositoryMock) New(transaction *domain.Transaction) error {
	args := r.Called(transaction)

	var r0 error
	if rf, ok := args.Get(0).(func(transaction *domain.Transaction) error); ok {
		r0 = rf(transaction)
	} else {
		r0 = args.Error(0)
	}

	return r0
}

func (r *TransactionRepositoryMock) Get(code string) ([]domain.Transaction, error) {
	args := r.Called(code)

	var r0 []domain.Transaction
	if rf, ok := args.Get(0).(func(code string) []domain.Transaction); ok {
		r0 = rf(code)
	} else {
		r0 = args.Get(0).([]domain.Transaction)
	}

	var r1 error
	if rf, ok := args.Get(1).(func(code string) error); ok {
		r1 = rf(code)
	} else {
		r1 = args.Error(1)
	}

	return r0, r1
}
