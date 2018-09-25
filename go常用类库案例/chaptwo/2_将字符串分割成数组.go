package main

import (
	"fmt"
	"regexp"
	"strings"
)

const refString2  =  "Mary_had a little_lamb"
const refString3 = "加长款*had ,a%little_lamb"
func main(){
	words := strings.Fields(refString2)
	for idx,word := range words {
		fmt.Printf("world %d is : %s \n",idx,word)
	}
	fmt.Println("---------------")
	words = strings.Split(refString2,"_")
	for idx,word := range words {
		fmt.Printf("world %d is : %s \n",idx,word)
	}

	fmt.Println("-----自定义分割函数-----------")
	splitFunc := func(r rune) bool {
		return strings.ContainsRune("*%,_",r)
	}
	words = strings.FieldsFunc(refString3,splitFunc)
	for idx,word := range words {
		fmt.Printf("the world is %d,%s\n",idx,word)
	}

	fmt.Println("---使用正则表达式-------")
	const refString4 = "Mary*had a %little_lamb"
	words = regexp.MustCompile("[*,%_]{1}").Split(refString4,-1)
	fmt.Println(words)
}


