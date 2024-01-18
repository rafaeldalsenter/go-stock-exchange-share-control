package domain

import (
	"testing"

	"github.com/bxcodec/faker"
	"github.com/stretchr/testify/assert"
)

func TestNewStockWithValidValues(t *testing.T) {
	assert := assert.New(t)
	code := "validCode"

	var transactions []Transaction
	errf := faker.FakeData(&transactions)
	assert.NoError(errf)

	stock, err := NewStock(code, transactions)

	assert.NoError(err)
	assert.Equal(code, stock.Code)
}

func TestNewStockWithInvalidCode(t *testing.T) {
	assert := assert.New(t)
	code := ""

	var transactions []Transaction
	errf := faker.FakeData(&transactions)
	assert.NoError(errf)

	_, err := NewStock(code, transactions)

	assert.Error(err)
}

func TestNewStockWithoutTransactions(t *testing.T) {
	assert := assert.New(t)
	code := "validCode"

	var transactions []Transaction

	_, err := NewStock(code, transactions)

	assert.Error(err)
}
