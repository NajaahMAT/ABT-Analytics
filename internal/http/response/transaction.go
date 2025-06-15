package response

type CountrySummary struct {
	Country             string                  `json:"country"`
	CountryTotalRevenue float64                 `json:"country_total_revenue"`
	Products            []ProductRevenueSummary `json:"products"`
}

type ProductRevenueSummary struct {
	ProductName      string  `json:"product_name"`
	TotalRevenue     float64 `json:"total_revenue"`
	TransactionCount int     `json:"transaction_count"`
}

type CountryRevenue struct {
	Country          string  `json:"country"`
	ProductName      string  `json:"product_name"`
	TotalRevenue     float64 `json:"total_revenue"`
	TransactionCount int     `json:"transaction_count"`
}

type ProductStats struct {
	ProductName   string `json:"product_name"`
	TotalQuantity int    `json:"total_quantity"`
	StockQuantity int    `json:"stock_quantity"`
}

type MonthlySales struct {
	Month  string `json:"month"`
	Volume int    `json:"volume"`
}

type RegionStats struct {
	Region       string  `json:"region"`
	TotalRevenue float64 `json:"total_revenue"`
	ItemsSold    int     `json:"items_sold"`
}
