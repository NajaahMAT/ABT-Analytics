package test

import (
	"testing"
	"time"

	"abt-analytics/internal/analytics"
	"abt-analytics/internal/model"

	"github.com/stretchr/testify/require"
)

func TestAggregator_ProcessAndQueries(t *testing.T) {
	a := analytics.NewAggregator()

	// Insert test data
	transactions := []model.Transaction{
		{
			Country:         "US",
			ProductName:     "Gadget", // Top product
			TotalPrice:      500,
			Quantity:        10,
			StockQuantity:   100,
			TransactionDate: parseDate("2025-02-10"),
			Region:          "NY",
		},
		{
			Country:         "US",
			ProductName:     "Widget",
			TotalPrice:      200,
			Quantity:        5,
			StockQuantity:   50,
			TransactionDate: parseDate("2025-02-15"),
			Region:          "NY",
		},
		{
			Country:         "CA",
			ProductName:     "Gadget",
			TotalPrice:      300,
			Quantity:        3,
			StockQuantity:   80,
			TransactionDate: parseDate("2025-01-05"),
			Region:          "Toronto",
		},
	}

	for _, tx := range transactions {
		a.Process(tx)
	}

	t.Run("CountrySummary pagination", func(t *testing.T) {
		result := a.CountryRevenueSummaryPaginated(0, 2, 2)
		require.Len(t, result, 2)
		require.Equal(t, "US", result[0].Country)
		require.Equal(t, "CA", result[1].Country)
	})

	t.Run("TopProducts", func(t *testing.T) {
		products := a.TopProducts(2)
		require.Len(t, products, 2)
		require.Equal(t, "Gadget", products[0].ProductName, "expected Gadget as top product")
	})

	t.Run("MonthlySales", func(t *testing.T) {
		sales := a.MonthlySales()
		require.GreaterOrEqual(t, len(sales), 1)
	})

	t.Run("TopRegions", func(t *testing.T) {
		regions := a.TopRegions(2)
		require.Len(t, regions, 2)
	})
}

func parseDate(s string) time.Time {
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		panic(err)
	}
	return t
}
