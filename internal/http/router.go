package httpapi

import (
	"net/http"

	"abt-analytics/internal/analytics"
	handler "abt-analytics/internal/http/handlers"

	httpSwagger "github.com/swaggo/http-swagger"
)

// NewRouter wires all routes to their handlers and returns an http.Handler.
func NewRouter(ag *analytics.Aggregator) http.Handler {
	h := handler.NewHandler(ag) // Uses exported constructor

	mux := http.NewServeMux()

	mux.HandleFunc("/api/revenue/country/summary", h.CountrySummary)
	mux.HandleFunc("/api/products/top", h.TopProducts)
	mux.HandleFunc("/api/sales/monthly", h.MonthlySales)
	mux.HandleFunc("/api/regions/top", h.TopRegions)

	//swagger endpoint
	mux.Handle("/swagger/", httpSwagger.Handler())

	return mux
}
