package loader

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"abt-analytics/internal/model"
)

// LoadTransactions now processes and aggregates CSV data without storing all rows in memory.
func LoadTransactions(path string, process func(model.Transaction)) error {
	file, err := os.Open(path)
	if err != nil {
		return fmt.Errorf("open csv: %w", err)
	}
	defer file.Close()

	r := csv.NewReader(file)
	r.ReuseRecord = true

	// Skip header
	if _, err := r.Read(); err != nil {
		return fmt.Errorf("read header: %w", err)
	}

	layout := "2006-01-02"
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("parse csv: %w", err)
		}

		price, _ := strconv.ParseFloat(rec[8], 64)
		qty, _ := strconv.Atoi(rec[9])
		stock, _ := strconv.Atoi(rec[10])
		tDate, _ := time.Parse(layout, rec[1])
		aDate, _ := time.Parse(layout, rec[11])

		tx := model.Transaction{
			TransactionID:   rec[0],
			TransactionDate: tDate,
			UserID:          rec[2],
			Country:         rec[3],
			Region:          rec[4],
			ProductID:       rec[5],
			ProductName:     rec[6],
			Category:        rec[7],
			Price:           price,
			Quantity:        qty,
			TotalPrice:      price * float64(qty),
			StockQuantity:   stock,
			AddedDate:       aDate,
		}
		process(tx) // Pass to aggregator function
	}
	return nil
}
