package main

import (
	"fmt"
	"math/big"
)

func main(){
	var da float64 = 0.2999992
	var db float64 = 0.2999991

	var prec uint = 32
	var prec2 uint = 16

	fmt.Printf("Compare floag64 with '==' equals:%v\n",da == db)

	daB := big.NewFloat(da).SetPrec(prec)
	dbB := big.NewFloat(db).SetPrec(prec)

	fmt.Printf("A:%v \n",daB)
	fmt.Printf("B:%v \n",dbB)
	fmt.Printf("Comparing big.Float with precision : %d : %v\n",prec,daB.Cmp(dbB) == 0)

	daB = big.NewFloat(da).SetPrec(prec2)
	dbB = big.NewFloat(db).SetPrec(prec2)

	fmt.Printf("A:%v \n",daB)
	fmt.Printf("B:%v \n",dbB)
	fmt.Printf("Comparing big.Float with precision : %d : %v\n",prec2,daB.Cmp(dbB) == 0)





}
