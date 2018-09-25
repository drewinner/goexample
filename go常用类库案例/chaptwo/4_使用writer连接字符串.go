package main

import (
	"bytes"
	"fmt"
)

func main(){
	strings := []string{
		"this","is","even","more","performant",
	}

	buffer := bytes.Buffer{}
	for _,val := range strings {
		buffer.WriteString(val)
	}
	fmt.Println(buffer.String())

	strings2 := []string {
		"this","is","even","more","performant",
	}

	bs := make([]byte,100)
	b1 := 0
	for _,val := range strings2 {
		b1+=copy(bs[b1:],[]byte(val))
	}
	fmt.Println(string(bs[:]))
}
