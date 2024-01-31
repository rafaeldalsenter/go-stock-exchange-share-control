package domain

import (
	"testing"
	"time"

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
		{
			Quantity: 5,
			Value:    10,
			Type:     Purchase,
		},
		{
			Quantity: 3,
			Value:    20,
			Type:     Purchase,
		},
		{
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
		{
			Quantity: 100,
			Value:    30,
			Tax:      10,
			Type:     Purchase,
		},
		{
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

func TestSalesResultNegativeWithoutTax(t *testing.T) {
	assert := assert.New(t)

	var transactions = []Transaction{
		{
			Quantity: 10,
			Value:    11,
			Type:     Sale,
			Date:     getDateByString("2024-01-18T10:20:50.52Z"),
		},
		{
			Quantity: 10,
			Value:    10.25,
			Type:     Purchase,
			Date:     getDateByString("2024-01-16T07:20:50.52Z"),
		},
		{
			Quantity: 12,
			Value:    11.25,
			Type:     Purchase,
			Date:     getDateByString("2024-01-17T07:20:50.52Z"),
		},
		{
			Quantity: 13,
			Value:    12.75,
			Type:     Purchase,
			Date:     getDateByString("2024-01-18T07:20:50.52Z"),
		},
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	salesResult, err := stock.SalesResult()

	assert.NoError(err)
	assert.Equal(-5.20, salesResult[0].Result)
}

func TestSalesResultPositiveWithoutTax(t *testing.T) {
	assert := assert.New(t)

	var transactions = []Transaction{
		{
			Quantity: 12,
			Value:    25,
			Type:     Sale,
			Date:     getDateByString("2024-01-18T10:20:50.52Z"),
		},
		{
			Quantity: 10,
			Value:    10.25,
			Type:     Purchase,
			Date:     getDateByString("2024-01-16T07:20:50.52Z"),
		},
		{
			Quantity: 12,
			Value:    11.25,
			Type:     Purchase,
			Date:     getDateByString("2024-01-17T07:20:50.52Z"),
		},
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	salesResult, err := stock.SalesResult()

	assert.NoError(err)
	assert.Equal(170.4, salesResult[0].Result)
}

func TestSalesResultWithTax(t *testing.T) {
	assert := assert.New(t)

	var transactions = []Transaction{
		{
			Quantity: 20,
			Value:    10,
			Tax:      0.5,
			Type:     Sale,
			Date:     getDateByString("2024-01-18T10:20:50.52Z"),
		},
		{
			Quantity: 10,
			Value:    11.8,
			Tax:      0.5,
			Type:     Purchase,
			Date:     getDateByString("2024-01-16T07:20:50.52Z"),
		},
		{
			Quantity: 12,
			Value:    9.8,
			Tax:      0.5,
			Type:     Purchase,
			Date:     getDateByString("2024-01-17T07:20:50.52Z"),
		},
		{
			Quantity: 9,
			Value:    7.0,
			Tax:      0.5,
			Type:     Purchase,
			Date:     getDateByString("2024-01-17T08:20:50.52Z"),
		},
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	salesResult, err := stock.SalesResult()

	assert.NoError(err)
	assert.Equal(5.9, salesResult[0].Result)
}

func TestSalesResultWithMultiples(t *testing.T) {
	assert := assert.New(t)

	var transactions = []Transaction{
		{
			Quantity: 12,
			Value:    25,
			Type:     Sale,
			Date:     getDateByString("2024-01-18T10:20:50.52Z"),
		},
		{
			Quantity: 10,
			Value:    5,
			Type:     Sale,
			Date:     getDateByString("2024-01-20T10:20:50.52Z"),
		},
		{
			Quantity: 10,
			Value:    10.25,
			Type:     Purchase,
			Date:     getDateByString("2024-01-16T07:20:50.52Z"),
		},
		{
			Quantity: 12,
			Value:    11.25,
			Type:     Purchase,
			Date:     getDateByString("2024-01-17T07:20:50.52Z"),
		},
	}

	stock, err := NewStock(code, transactions)
	assert.NoError(err)

	salesResult, err := stock.SalesResult()

	assert.NoError(err)
	assert.Equal(170.4, salesResult[0].Result)
	assert.Equal(-58.0, salesResult[1].Result)

}

func getDateByString(value string) time.Time {
	tim, err := time.Parse(time.RFC3339, value)
	if err != nil {
		panic(err)
	}

	return tim
}
