package main

import (
	"fmt"
	"regexp"
	"strings"
)

func main() {
	const refString = "Mary had a little lamb"
	const refString2 = "lamb lamb lamb"

	out := strings.Replace(refString, "lamb", "wolf", -1)
	fmt.Println(out)

	out = strings.Replace(refString2, "lamb", "wolf", 2)
	fmt.Println(out)

	fmt.Println("-------replacer替换字符串--------")

	const refString3 = "Mary had a little lamb"
	replacer := strings.NewReplacer("lamb", "wolf", "Mary", "Jack")
	out = replacer.Replace(refString3)
	fmt.Println(out)

	fmt.Println("-------使用正则表达式替换字符串------")

	const refString4 = "Mary had a little lamb"
	regex := regexp.MustCompile("l[a-z]+")
	out = regex.ReplaceAllString(refString4,"replacement")
	fmt.Println(out)
}
