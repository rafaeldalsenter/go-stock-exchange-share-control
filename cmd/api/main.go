package api

import (
	"go-stock-exchange-shares-control/infra/mongo"
	"go-stock-exchange-shares-control/internal/controllers"
	"go-stock-exchange-shares-control/internal/usecase"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	// build
	tr := mongo.NewTransactionRepositoryMongo("")
	s := usecase.NewStockPriceUseCase(tr)
	t := usecase.NewTransactionUseCase(tr)
	c := controllers.NewController(s, t)

	r.Post("/stock/{code}/transaction", controllers.ControllerBase(c.TransactionPost))
	r.Get("/stock/{code}/average-purchase", controllers.ControllerBase(c.AveragePurchaseGet))
	r.Get("/stock/{code}/average-selling", controllers.ControllerBase(c.AverageSellingGet))

	http.ListenAndServe(":3000", r)
}
