package main

import (
	"bufio"
	"fmt"
	"os"
)
//可以读取多行
func main(){
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		txt := sc.Text()
		fmt.Printf("Echo : %s\n",txt)
	}
}
