package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {

	prc := exec.Command("ls", "-a")
	out := bytes.NewBuffer([]byte{})
	prc.Stdout = out
	//同步执行
	err := prc.Run()
	if err != nil {
		fmt.Println(err)
	}

	if prc.ProcessState.Success() {
		fmt.Println("Process run successful with output :\n")
		fmt.Println(out.String())
	}

	prc2 := exec.Command("ls", "-a")
	//异步执行
	err2 := prc2.Start()
	if err2 != nil {
		fmt.Println(err2)
	}
	prc2.Wait()
	//prc2.ProcessState.Pid()
	if prc2.ProcessState.Success() {
		fmt.Println("Process run successful with output :\n")
		//fmt.Println(out.String())
	}

}
