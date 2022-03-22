package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"time"
)

func main() {
	args := os.Args[1:]
	cmd := args[0]
	switch cmd {
	case "add":
		add(time.Now(), args[1])
	}
}

func add(date time.Time, weight string) error {
	f, err := os.OpenFile("weight.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open weight.csv, %w", err)
	}

	wr := csv.NewWriter(f)
	defer wr.Flush()
	err = wr.Write([]string{date.Format(time.RFC3339), weight})
	if err != nil {
		return fmt.Errorf("failed to write weight, %w", err)
	}

	return nil
}
