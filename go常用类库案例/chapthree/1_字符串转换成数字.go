package main

import (
	"fmt"
	"strconv"
)

func main(){

	const bin = "00001"
	const hex = "2f"
	const intString = "12"
	const floatString = "12.3"

	resInt,err := strconv.ParseInt(hex,16,32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("the integer is : %d\n",resInt)
	res,err := strconv.Atoi(intString)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Parsed integer : %d\n",res)

	res64,err := strconv.ParseInt(hex,16,32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed hexadecima %d \n",res64)

	resBin,err := strconv.ParseInt(bin,2,32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed bin %d\n",resBin)

	resFloat,err := strconv.ParseFloat(floatString,32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("parsed floag : %.5f\n",resFloat)

}