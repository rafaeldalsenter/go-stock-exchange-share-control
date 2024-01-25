package main

import (
	"encoding/csv"
	"go-stock-exchange-shares-control/infra/mongo"
	"go-stock-exchange-shares-control/internal/usecase"
	"go-stock-exchange-shares-control/internal/utils"
	"io"
	"log"
	"os"

	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile("../../config.json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

func main() {
	mongoConnectionString := viper.GetString("mongo.connectionString")
	mongoDatabase := viper.GetString("mongo.database")
	mongoCollection := viper.GetString("mongo.collection")

	tr := mongo.NewTransactionRepositoryMongo(mongoConnectionString, mongoDatabase, mongoCollection)
	t := usecase.NewTransactionUseCase(tr)

	filename := viper.GetString("file.name")

	log.Printf("Starting %s file", filename)

	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	csvReader := csv.NewReader(file)

	for {

		// TODO tratar erros ?? para nao inserir duplicado / idepotencia
		// TODO adicionar assincronismo no processamento
		line, err := csvReader.Read()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal(err)
			continue
		}

		transaction, err := utils.ConverterTransactionValues(line[1], line[2], line[3], line[4], line[5])
		if err != nil {
			log.Printf(err.Error())
			continue
		}

		if transaction.Type == "purchase" {
			_, err := t.NewPurchase(line[0], transaction.Date, transaction.Quantity, transaction.Value, transaction.Tax)
			if err != nil {
				log.Printf(err.Error())
				continue
			}
			log.Printf("Transaction %s/%s imported", line[0], line[1])

		} else if transaction.Type == "sale" {
			_, err := t.NewPurchase(line[0], transaction.Date, transaction.Quantity, transaction.Value, transaction.Tax)
			if err != nil {
				log.Printf(err.Error())
				continue
			}
			log.Printf("Transaction %s/%s imported", line[0], line[1])
		} else {
			log.Printf("Not identified transaction %s type", transaction.Type)
		}
	}
}
