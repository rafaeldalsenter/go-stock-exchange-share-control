package main

import (
	"go-stock-exchange-shares-control/infra/mongo"
	"go-stock-exchange-shares-control/internal/controllers"
	"go-stock-exchange-shares-control/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
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
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	mongoConnectionString := viper.GetString("mongo.connectionString")
	mongoDatabase := viper.GetString("mongo.database")
	mongoCollection := viper.GetString("mongo.collection")

	tr := mongo.NewTransactionRepositoryMongo(mongoConnectionString, mongoDatabase, mongoCollection)
	s := usecase.NewStockPriceUseCase(tr)
	t := usecase.NewTransactionUseCase(tr)
	c := controllers.NewController(s, t)

	r.Post("/stock/{code}/transaction", controllers.ControllerBase(c.TransactionPost))
	r.Get("/stock/{code}/average-purchase", controllers.ControllerBase(c.AveragePurchaseGet))

	http.ListenAndServe(viper.GetString("server.address"), r)
}
