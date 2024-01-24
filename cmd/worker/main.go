package main

import (
	"go-stock-exchange-shares-control/infra/mongo"
	"go-stock-exchange-shares-control/internal/usecase"

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
	_ = usecase.NewTransactionUseCase(tr)

	// todo ler os itens da planilha

	// todo dar um new PUrchase ou Sale para cada,

	// tratar erros ?? para nao inserir duplicado
}
