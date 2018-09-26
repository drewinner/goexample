package main

import (
	"fmt"
	"strconv"
)

func main() {
	const bin = "10111"
	const hex = "1A"
	const oct = "12"
	const dec = "10"
	const floatNum = 16.123557

	//二进制转十六进制
	v,_ := ConvertInt(bin,2,16)
	fmt.Printf("Binary value %s converted to hex : %s\n",bin,v)

	v,_ = ConvertInt(hex,16,10)
	fmt.Printf("hex value %s converted to dec : %s \n ",hex,v)

	// Converts oct value into hex
	v, _ = ConvertInt(oct, 8, 16)
	fmt.Printf("Oct value %s converted to hex: %s\n", oct, v)
	// Converts dec value into oct
	v, _ = ConvertInt(dec, 10, 8)
	fmt.Printf("Dec value %s converted to oct: %s\n", dec, v)
}

func ConvertInt(val string,base ,toBase int ) (string ,error){
	i,err := strconv.ParseInt(val,base,64)
	if err != nil {
		return "",err
	}
	return strconv.FormatInt(i,toBase),nil
}
