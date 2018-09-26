package main

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func main() {
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("create file error = ", err)
		return
	}
	w := tabwriter.NewWriter(f, 15, 0, 1, '-', tabwriter.AlignRight)
	fmt.Fprintln(w, "username\tfirstname\tlastname\t")
	fmt.Fprintln(w, "sohlich\tRadomir\tSohlich\t")
	w.Flush()

}
