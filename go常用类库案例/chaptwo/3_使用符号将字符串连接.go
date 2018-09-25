package main

import (
	"fmt"
	"strings"
)

func main(){
	const selectBase = "select * from user where %s "
	var refStringSlice = []string{
		" FIRST_NAME = 'jack' ",
		" INSURANGE_NO = 333444555 ",
		" EFFECTIVE_FROM = SYSDATE ",
	}
	sentence := strings.Join(refStringSlice,"AND")
	fmt.Printf(selectBase+"\n",sentence)
}