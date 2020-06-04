package handler

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

func CreateNewFile() {
	_, err := os.Create("File-" + time.Now().String() + ".csv")
	if err != nil {
		log.Printf("Could not create file: %v", err)
	}
}

func SaveToFile(file string, name string, square float64, price float64) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		log.Printf("Could not open file: %v", err)
	}
	defer f.Close()

	record := []string{name, fmt.Sprintf("%.2f", square), fmt.Sprintf("%.2f", price)}
	cw := csv.NewWriter(f)
	cw.Write(record)
	cw.Flush()
}
