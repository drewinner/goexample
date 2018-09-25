package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

//1,信号是操作系统和正在运行进程之间交流的基本方式
//2,SIGINT,SIGTERM,SIGUP
func main() {
	sChan := make(chan os.Signal, 1)

	signal.Notify(sChan,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,

	)
	exitChan := make(chan int)

	go func() {
		signal := <-sChan
		switch signal {
		case syscall.SIGHUP:
			fmt.Println("the calling terminal has been closed")
			exitChan <- 0
		case syscall.SIGINT:
			fmt.Println("the process has been interrupted by CTRL+C")
			exitChan <- 1
		case syscall.SIGTERM:
			fmt.Println("kill SIGTERM was executed for process")
			exitChan <- 1
		case syscall.SIGQUIT:
			fmt.Println("kill SIGQUIT was executed for process")
			exitChan <- 1
		}
	}()
	code := <-exitChan
	os.Exit(code)
}
