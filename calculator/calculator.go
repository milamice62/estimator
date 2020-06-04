package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculatePrice(unitPrice string, quantity string) float64 {
	up, err := strconv.ParseFloat(strings.TrimSpace(unitPrice), 64)
	if err != nil {
		fmt.Printf("Not invalid number: %v", err)
	}

	qa, err := strconv.Atoi(strings.TrimSpace(quantity))
	if err != nil {
		fmt.Printf("Not invalid number: %v", err)
	}

	price := up * float64(qa)
	return price
}
