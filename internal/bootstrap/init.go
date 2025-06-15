package bootstrap

import (
	"log"

	"abt-analytics/docs"
	"abt-analytics/internal/analytics"
	"abt-analytics/internal/loader"
)

// InitAggregator loads transactions and returns an aggregated instance
func InitAggregator() *analytics.Aggregator {
	aggregator := analytics.NewAggregator()
	err := loader.LoadTransactions("data/transactions.csv", aggregator.Process)
	if err != nil {
		log.Fatalf("failed to load transactions: %v", err)
	}
	return aggregator
}

// InitSwagger sets up swagger metadata info
func InitSwagger() {
	docs.SwaggerInfo.Title = "ABT Analytics API"
	docs.SwaggerInfo.Description = "API documentation for ABT Analytics"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Host = "localhost:8080"
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"}
}
