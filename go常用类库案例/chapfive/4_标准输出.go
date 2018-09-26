package main

import (
	"fmt"
	"io"
	"os"
)

func main(){

	io.WriteString(os.Stdout,"this is string to standard output\n")
	io.WriteString(os.Stderr,"this is string to standard err output \n")

	buf := []byte{0xAF,0xFF,0xFE}
	for i:=0;i<200;i++{
		if _,e := os.Stdout.Write(buf); e != nil {
			panic(e)
		}
	}
	fmt.Fprintln(os.Stdout,"\n")
}
