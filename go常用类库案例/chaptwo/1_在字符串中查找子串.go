package main

import (
	"fmt"
	"strings"
)

const refString  = "Mary had a little lamb"

func main(){
	lookFor := "lamb"
	contain := strings.Contains(refString,lookFor)
	fmt.Println("the refString contain lamp",contain)

	startsWith := strings.HasPrefix(refString,"Mary")
	fmt.Println("the refString startwith is :",startsWith)

	endWith := strings.HasSuffix(refString,"lamb")
	fmt.Println("the refString enddWith",endWith)
}
