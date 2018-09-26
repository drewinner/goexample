package main


import "fmt"

func main(){

	var name string
	fmt.Println("what is your name ?")
	fmt.Scanf("%s\n",&name)

	var age int
	fmt.Println("how old are you ?")
	fmt.Scanf("%d\n",&age)

	fmt.Printf("hello %s,your age %d \n",name,age)
}
