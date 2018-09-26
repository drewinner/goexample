package main

import (
	"fmt"
	"math"
)

func main() {
	const da = 0.2999999999999999
	const db = 0.3

	daStr := fmt.Sprintf("%.10f", da)
	dbStr := fmt.Sprintf("%.10f", db)

	fmt.Printf("strings %s = %s equals:%v\n", daStr, dbStr, daStr == dbStr)
	fmt.Printf("Number equals:%v\n", da == db)

	fmt.Printf("Number equals with TOLERANGE:%v\n", equals(da, db))
}

const TOLERANGE = 1e-8

func equals(numA, numB float64) bool {

	delta := math.Abs(numA - numB)

	if delta < TOLERANGE {
		return true
	}
	return false
}
