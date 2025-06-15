package http

import (
	"encoding/json"
	"net/http"
	"strconv"

	"abt-analytics/internal/analytics"
)

// Handler bundles shared deps (here: the in‑memory Aggregator).
type Handler struct {
	ag *analytics.Aggregator
}

// Exported constructor
func NewHandler(ag *analytics.Aggregator) *Handler {
	return &Handler{ag: ag}
}

// ---- helpers ------------------------------------------------

func writeJSON(w http.ResponseWriter, v any) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(v)
}

func parseQueryInt(val string, def int) int {
	if n, err := strconv.Atoi(val); err == nil && n > 0 {
		return n
	}
	return def
}

// CountrySummary godoc
// @Summary      Get country revenue summary
// @Description  Returns a paginated list of countries with revenue & top products
// @Tags         revenue
// @Accept       json
// @Produce      json
// @Param        page     query int false "Page number"
// @Param        size     query int false "Page size"
// @Param        products query int false "Number of top products per country"
// @Success      200 {array} response.CountrySummary
// @Router       /api/revenue/country/summary [get]
func (h *Handler) CountrySummary(w http.ResponseWriter, r *http.Request) {
	q := r.URL.Query()
	page := parseQueryInt(q.Get("page"), 1)
	size := parseQueryInt(q.Get("size"), 20)
	productLimit := parseQueryInt(q.Get("products"), 10)

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * size
	writeJSON(w, h.ag.CountryRevenueSummaryPaginated(offset, size, productLimit))
}

// TopProducts godoc
// @Summary      Get top products globally
// @Description  Returns a fixed list of top 20 products globally
// @Tags         products
// @Accept       json
// @Produce      json
// @Success      200 {array} response.ProductStats
// @Router       /api/products/top [get]
func (h *Handler) TopProducts(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.ag.TopProducts(20))
}

// MonthlySales godoc
// @Summary      Get monthly sales data
// @Description  Returns aggregated monthly sales data
// @Tags         sales
// @Accept       json
// @Produce      json
// @Success      200 {array} response.MonthlySales
// @Router       /api/sales/monthly [get]
func (h *Handler) MonthlySales(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.ag.MonthlySales())
}

// TopRegions godoc
// @Summary      Get top sales regions
// @Description  Returns a list of top 30 sales regions
// @Tags         regions
// @Accept       json
// @Produce      json
// @Success      200 {array} response.RegionStats
// @Router       /api/regions/top [get]
func (h *Handler) TopRegions(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, h.ag.TopRegions(30))
}
