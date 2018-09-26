package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main(){

	f,err := os.Open("D:/go_project/goexample/goexample/go常用类库案例/chapfive/file.txt")
	if err != nil {
		panic(err)
	}
	c,err := ioutil.ReadAll(f)
	if err !=nil {
		panic(err)
	}

	fmt.Printf("###file content ###\n%s\n",string(c))
	f.Close()

	f,err = os.OpenFile("D:/go_project/goexample/goexample/go常用类库案例/chapfive/file1.txt",os.O_CREATE|os.O_RDWR,os.ModePerm)
	if err != nil {
		panic(err)
	}

	io.WriteString(f,"test string")
	f.Close()
}
