package domain

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	date     = time.Now()
	quantity = 1.4
	value    = 2.3
	tax      = 12.0
	tr       = Purchase
)

func TestNewTransactionWithValidValues(t *testing.T) {
	assert := assert.New(t)

	transaction, err := NewTransaction(tr, date, quantity, value, tax)

	assert.NoError(err)
	assert.Equal(date, transaction.Date)
	assert.Equal(quantity, transaction.Quantity)
	assert.Equal(value, transaction.Value)
	assert.Equal(tax, transaction.Tax)
	assert.Equal(tr, transaction.Type)
}

func TestNewTransactionWithInvalidQuantity(t *testing.T) {
	assert := assert.New(t)
	invalidQuantity := float64(-1)

	_, err := NewTransaction(tr, date, invalidQuantity, value, tax)

	assert.Error(err)
}

func TestNewTransactionWithInvalidValue(t *testing.T) {
	assert := assert.New(t)
	invalidValue := float64(-1)

	_, err := NewTransaction(tr, date, quantity, invalidValue, tax)

	assert.Error(err)
}
