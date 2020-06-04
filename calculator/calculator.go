package calculator

import (
	"fmt"
	"strconv"
	"strings"
)

func CalculatePrice(price string, square string) (float64, float64) {
	const smtosf = 1.0764e-05

	p, err := strconv.ParseFloat(strings.TrimSpace(price), 64)
	if err != nil {
		fmt.Printf("Not invalid number: %v", err)
	}

	sqmm, err := strconv.ParseFloat(strings.TrimSpace(square), 64)
	if err != nil {
		fmt.Printf("Not invalid number: %v", err)
	}

	sqft := smtosf * sqmm
	pricesqft := sqft * p

	return sqft, pricesqft
}
