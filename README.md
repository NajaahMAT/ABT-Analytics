# ABT Analytics

📘 **Project Description**  
ABTAnalytics is a high-performance, extensible analytics dashboard built using Go for backend processing and a modern frontend framework for data visualization. It provides actionable insights derived from ABT Corporation’s business data, including product sales, transactions, and customer regions.

The system ingests and processes large datasets to deliver fast, clear, and interactive visual summaries across multiple business dimensions. It is designed for scalability, testability, and future extension of analytics features.

---

## 🔍 Key Features

- **Country-Level Revenue Table**  
  Displays each country's revenue breakdown by product, total revenue, and transaction counts.

- **Top 20 Products Visualization**  
  Highlights the most frequently purchased products along with stock availability.

- **Monthly Sales Trends**  
  An interactive chart showcasing peak sales periods by month.

- **Top 30 Regions by Revenue**  
  Visualizes the highest revenue-generating regions and total items sold.

---

## ⚙️ Technical Highlights

### Backend
- Developed in Go for optimal performance and concurrency.
- REST APIs expose processed metrics efficiently.
- Modular and testable codebase with unit test support.

### Frontend
- Built using a lightweight charting framework (e.g., HTML, Barchart).

### Data Handling
- Aggregation and transformation pipelines written in Go.
- Efficient sorting, filtering, and grouping of business data.

### Performance
- Dashboard loads within 10s for ~100k rows.
- Caching and in-memory structures used for repeated query responses.

---

## 📁 Deliverables

- Fully working source code (backend + frontend)
- Dataset processing and transformation scripts
- Unit tests and test coverage report
- Complete setup instructions (see below)

---

# 📖 API Reference

All endpoints are available at:  
**Base URL:** `http://localhost:8080`

Interactive Swagger UI:  
**🔗 Swagger Docs:** [`http://localhost:8080/swagger/`](http://localhost:8080/swagger/)

---

## Endpoints Overview

| Endpoint                             | Method | Description                                                      | Query / Path Parameters                        |
|--------------------------------------|--------|------------------------------------------------------------------|------------------------------------------------|
| `/api/revenue/country/summary`      | GET    | Paginated revenue table for each country, including top‑N products. | `page` (int), `size` (int), `products` (int)   |
| `/api/products/top`                 | GET    | Retrieves the 20 most‑sold products across all countries.        | –                                              |
| `/api/sales/monthly`               | GET    | Returns aggregated monthly sales volumes.                        | –                                              |
| `/api/regions/top`                 | GET    | Fetches the top 30 revenue-generating regions and items sold.    | –                                              |

---

## Instruction

Note: 
1. API Postman collection added in the docs folder
2. Data csv file cannot be pushed to git. Because of that it has compressed , Decomposs the file before using it.



## Quick Start
```bash
git clone <repo>
cd abt-analytics
go mod tidy
go run ./cmd &
```
Open `frontend/index.html` in your browser.

## Running Tests
```bash
chmod +x test.sh
./test.sh
```




