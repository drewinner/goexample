package main

import (
	"errors"
	"fmt"
	"os"
)

func main(){
	file,err := os.OpenFile("abc.txt",os.O_EXCL|os.O_CREATE,0666)
	err = errors.New("this a custom error")
	if err != nil {
		if pathError,ok := err.(*os.PathError);!ok {
			panic(err)
		}else{
			fmt.Printf("%s,%s,%s \n",pathError.Op,pathError.Path,pathError.Err)
		}
	}
	defer file.Close()
}
