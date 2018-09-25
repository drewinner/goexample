package pracitice

import (
	"fmt"
	"os"
)

func GetArgs(){

	args := os.Args
	fmt.Println(args)

	programName := args[0]
	fmt.Println("The binary name is %s",programName)

	if len(args) > 1{
		otherArgs := args[1:]
		for idx,itemArg := range(otherArgs){
			fmt.Printf("Arg %d = %s \n",idx,itemArg)
		}
	}

}
