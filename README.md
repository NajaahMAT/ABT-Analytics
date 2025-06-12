# ABT-Analytics
📘 Project Description:
ABTAnalytics is a high-performance, extensible analytics dashboard built using Go for backend processing and a modern frontend framework for data visualization. It provides actionable insights derived from ABT Corporation’s business data, including product sales, transactions, and customer regions.

The system ingests and processes large datasets to deliver fast, clear, and interactive visual summaries across multiple business dimensions. It is designed for scalability, testability, and future extension of analytics features.

🔍 Key Features:
Country-Level Revenue Table:
Displays each country's revenue breakdown by product, total revenue, and transaction counts, sorted by revenue.

Top 20 Products Visualization:
Highlights the most frequently purchased products with their stock availability.

Monthly Sales Trends:
An interactive chart showcasing peak sales periods by month.

Top 30 Regions by Revenue:
Visual representation of the highest revenue-generating regions and total items sold.

⚙️ Technical Highlights:
Backend:

Developed in Go for optimal performance and concurrency.

REST APIs expose processed metrics efficiently.

Modular and testable codebase with unit tests and test coverage.

Frontend:

Built with a lightweight charting framework (e.g., React + Recharts).

Responsive, clean UI for a seamless user experience.

Data Handling:

Aggregation and transformation pipelines built in Go.

Efficient sorting, filtering, and grouping of business data.

Performance:

Optimized to ensure dashboard loads in under 10 seconds.

Caching and in-memory structures used for repeated query responses.

📁 Deliverables:
Fully working source code (backend + frontend)

Dataset processing and transformation scripts

Unit tests and test coverage report

Complete setup instructions in README.md

