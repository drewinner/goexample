package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main(){
	ex,err := os.Executable()
	if err != nil {
		panic(err)
	}
	fmt.Println(ex)

	exPath := filepath.Dir(ex)
	fmt.Println("Executable path:"+exPath)

	realPath,err := filepath.EvalSymlinks(exPath)
	if err != nil {
		panic(err)
	}
	fmt.Println("Symlinks evaluated:"+realPath)
}