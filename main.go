package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: forma <weight> <note (optional)>")
		os.Exit(1)
	}

	args := os.Args[1:]
	cmd := args[0]
	switch cmd {
	case "add":
		var note string
		if len(args) > 2 {
			note = strings.Join(args[2:], " ")
		}

		m := &Measurement{
			Weight: args[1],
			Date:   time.Now(),
			Note:   note,
		}
		err := add(m)
		if err != nil {
			fmt.Printf("could not add measurement: %v\n", err)
			os.Exit(1)
		}
	}
}

type Measurement struct {
	Date   time.Time
	Weight string
	Note   string
}

func add(m *Measurement) error {
	f, err := os.OpenFile("weight.csv", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to open weight.csv, %w", err)
	}

	wr := csv.NewWriter(f)
	defer wr.Flush()
	err = wr.Write([]string{m.Date.Format(time.RFC3339), m.Weight, m.Note})
	if err != nil {
		return fmt.Errorf("failed to write weight, %w", err)
	}

	return nil
}
