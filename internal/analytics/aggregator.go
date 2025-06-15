package analytics

import (
	"container/heap"
	"sort"

	"abt-analytics/internal/http/response"
	"abt-analytics/internal/model"
)

//--------------------------------------------------------------
//  Aggregator struct & streaming Process
//--------------------------------------------------------------

type Aggregator struct {
	countryRevenueMap map[string]map[string]*response.CountryRevenue
	productStatsMap   map[string]*response.ProductStats
	monthlySalesMap   map[string]int
	regionStatsMap    map[string]*response.RegionStats
}

func NewAggregator() *Aggregator {
	return &Aggregator{
		countryRevenueMap: make(map[string]map[string]*response.CountryRevenue),
		productStatsMap:   make(map[string]*response.ProductStats),
		monthlySalesMap:   make(map[string]int),
		regionStatsMap:    make(map[string]*response.RegionStats),
	}
}

// Stream a single CSV row into in‑memory aggregates (maps only)
func (a *Aggregator) Process(t model.Transaction) {
	// Country × Product
	if a.countryRevenueMap[t.Country] == nil {
		a.countryRevenueMap[t.Country] = make(map[string]*response.CountryRevenue)
	}
	if a.countryRevenueMap[t.Country][t.ProductName] == nil {
		a.countryRevenueMap[t.Country][t.ProductName] = &response.CountryRevenue{
			Country:     t.Country,
			ProductName: t.ProductName,
		}
	}
	rec := a.countryRevenueMap[t.Country][t.ProductName]
	rec.TotalRevenue += t.TotalPrice
	rec.TransactionCount++

	// Product stats (top products endpoint)
	if a.productStatsMap[t.ProductName] == nil {
		a.productStatsMap[t.ProductName] = &response.ProductStats{
			ProductName:   t.ProductName,
			StockQuantity: t.StockQuantity,
		}
	}
	a.productStatsMap[t.ProductName].TotalQuantity += t.Quantity

	// Monthly sales
	month := t.TransactionDate.Format("2006-01")
	a.monthlySalesMap[month] += t.Quantity

	// Region stats
	if a.regionStatsMap[t.Region] == nil {
		a.regionStatsMap[t.Region] = &response.RegionStats{Region: t.Region}
	}
	reg := a.regionStatsMap[t.Region]
	reg.TotalRevenue += t.TotalPrice
	reg.ItemsSold += t.Quantity
}

//--------------------------------------------------------------
//  Country summary – memory‑safe + pagination
//--------------------------------------------------------------

type countryItem struct {
	revenue float64
	summary response.CountrySummary
}

type countryMinHeap []countryItem

func (h countryMinHeap) Len() int           { return len(h) }
func (h countryMinHeap) Less(i, j int) bool { return h[i].revenue < h[j].revenue }
func (h countryMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *countryMinHeap) Push(x any)        { *h = append(*h, x.(countryItem)) }
func (h *countryMinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

type productItem struct {
	revenue float64
	pr      response.ProductRevenueSummary
}

type productMinHeap []productItem

func (h productMinHeap) Len() int           { return len(h) }
func (h productMinHeap) Less(i, j int) bool { return h[i].revenue < h[j].revenue }
func (h productMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *productMinHeap) Push(x any)        { *h = append(*h, x.(productItem)) }
func (h *productMinHeap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

// CountryRevenueSummaryPaginated returns a slice of CountrySummary starting at `offset`
// (0‑based) returning at most `limitCountries` entries. Inside each country we keep the
// top `limitProducts` products by revenue.
func (a *Aggregator) CountryRevenueSummaryPaginated(offset, limitCountries, limitProducts int) []response.CountrySummary {
	// First stage: keep only (offset+limitCountries) countries using min‑heap
	keep := offset + limitCountries
	cHeap := &countryMinHeap{}
	heap.Init(cHeap)

	for country, prodMap := range a.countryRevenueMap {
		// Build top‑N products for this country using a product min‑heap
		pHeap := &productMinHeap{}
		heap.Init(pHeap)
		var countryTotal float64

		for _, p := range prodMap {
			heap.Push(pHeap, productItem{revenue: p.TotalRevenue, pr: response.ProductRevenueSummary{
				ProductName:      p.ProductName,
				TotalRevenue:     p.TotalRevenue,
				TransactionCount: p.TransactionCount,
			}})
			if pHeap.Len() > limitProducts {
				heap.Pop(pHeap) // remove smallest product
			}
			countryTotal += p.TotalRevenue
		}

		// extract products from heap into slice descending
		products := make([]response.ProductRevenueSummary, pHeap.Len())
		for i := len(products) - 1; i >= 0; i-- {
			products[i] = heap.Pop(pHeap).(productItem).pr
		}

		cs := response.CountrySummary{
			Country:             country,
			CountryTotalRevenue: countryTotal,
			Products:            products,
		}

		heap.Push(cHeap, countryItem{revenue: countryTotal, summary: cs})
		if cHeap.Len() > keep {
			heap.Pop(cHeap) // keep heap at most keep items
		}
	}

	// Extract from heap to slice ascending then reverse
	results := make([]response.CountrySummary, cHeap.Len())
	for i := len(results) - 1; i >= 0; i-- {
		results[i] = heap.Pop(cHeap).(countryItem).summary
	}

	// Apply offset / limitCountries window
	if offset >= len(results) {
		return []response.CountrySummary{}
	}
	end := offset + limitCountries
	if end > len(results) {
		end = len(results)
	}
	return results[offset:end]
}

// Convenience wrapper (no pagination, but product cap)
func (a *Aggregator) CountryRevenueSummary(limitProducts int) []response.CountrySummary {
	return a.CountryRevenueSummaryPaginated(0, len(a.countryRevenueMap), limitProducts)
}

//--------------------------------------------------------------
//  Other existing endpoints (top products, monthly sales, regions)
//--------------------------------------------------------------

func (a *Aggregator) TopProducts(n int) []response.ProductStats {
	var result []response.ProductStats
	for _, val := range a.productStatsMap {
		result = append(result, *val)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].TotalQuantity > result[j].TotalQuantity })
	if len(result) > n {
		result = result[:n]
	}
	return result
}

func (a *Aggregator) MonthlySales() []response.MonthlySales {
	var result []response.MonthlySales
	for month, qty := range a.monthlySalesMap {
		result = append(result, response.MonthlySales{Month: month, Volume: qty})
	}
	sort.Slice(result, func(i, j int) bool { return result[i].Volume > result[j].Volume })
	return result
}

func (a *Aggregator) TopRegions(n int) []response.RegionStats {
	var result []response.RegionStats
	for _, val := range a.regionStatsMap {
		result = append(result, *val)
	}
	sort.Slice(result, func(i, j int) bool { return result[i].TotalRevenue > result[j].TotalRevenue })
	if len(result) > n {
		result = result[:n]
	}
	return result
}
