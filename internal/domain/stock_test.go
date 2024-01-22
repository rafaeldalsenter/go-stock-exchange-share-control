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

func TestAveragePurchasePriceWithoutTax(t *testing.T) {
	assert := assert.New(t)

	var transactions = []Transaction{
		Transaction{
			Quantity: 5,
			Value:    10,
			Type:     Purchase,
		},
		Transaction{
			Quantity: 3,
			Value:    20,
			Type:     Purchase,
		},
		Transaction{
			Quantity: 1,
			Value:    100,
			Type:     Purchase,
		},
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	average, err := stock.AveragePurchasePrice()

	assert.NoError(err)
	assert.Equal(23.33, average)
}

func TestAveragePurchasePriceWithTax(t *testing.T) {
	assert := assert.New(t)

	var transactions = []Transaction{
		Transaction{
			Quantity: 100,
			Value:    30,
			Tax:      10,
			Type:     Purchase,
		},
		Transaction{
			Quantity: 50,
			Value:    25,
			Tax:      10,
			Type:     Purchase,
		},
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	average, err := stock.AveragePurchasePrice()

	assert.NoError(err)
	assert.Equal(28.47, average)
}

func TestAveragePurchasePriceWithoutRecords(t *testing.T) {
	assert := assert.New(t)

	var transactions []Transaction
	err := faker.FakeData(&transactions)
	assert.NoError(err)

	for i := 0; i < len(transactions); i++ {
		transactions[i].Type = Sale
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	_, err = stock.AveragePurchasePrice()

	assert.Error(err)
}

func TestAverageSellingPriceWithoutRecords(t *testing.T) {
	assert := assert.New(t)

	var transactions []Transaction
	err := faker.FakeData(&transactions)
	assert.NoError(err)

	for i := 0; i < len(transactions); i++ {
		transactions[i].Type = Purchase
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	_, err = stock.AverageSellingPrice()

	assert.Error(err)
}
