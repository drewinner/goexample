package main

import (
	"fmt"
	"math"
	"runtime"
)
var valA float64 = 3.35554444
func main() {
	// Bad assumption on rounding
	// the number by casting it to
	// integer.
	intVal := int(valA)
	fmt.Printf("Bad rounding by casting to int: %v\n", intVal)
	fRound := Round(valA)
	fmt.Printf("Rounding by custom function: %v\n", fRound)
	//from 1.10 can use
	fmt.Println(math.Round(valA))
	fmt.Println(runtime.Version())
}
// Round returns the nearest integer.
func Round(x float64) float64 {
	t := math.Trunc(x)
	if math.Abs(x-t) >= 0.5 {
		return t + math.Copysign(1, x)
	}
	return t
}