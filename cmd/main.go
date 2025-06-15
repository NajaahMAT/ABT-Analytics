package main

import (
	"abt-analytics/internal/bootstrap"
	http_api "abt-analytics/internal/http"
	"log"
	"net/http"
)

func main() {
	bootstrap.InitSwagger() // call the swagger init function
	aggregator := bootstrap.InitAggregator()

	router := http_api.NewRouter(aggregator)

	log.Println("API server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
