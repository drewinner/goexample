package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main(){

	file,err := os.Open("D:/go_project/goexample/goexample/go常用类库案例/chaptwo/data.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	for {
		record,e := reader.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}

	file1,err1 := os.Open("D:/go_project/goexample/goexample/go常用类库案例/chaptwo/data_uncommon.csv")
	if err1 != nil {
		panic(err1)
	}

	defer file1.Close()

	reader1 := csv.NewReader(file1)
	reader1.Comma = ';'
	for {
		record,e := reader1.Read()
		if e != nil {
			fmt.Println(e)
			break
		}
		fmt.Println(record)
	}
}