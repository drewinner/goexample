package pracitice

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

type ArrayValue []string

func AccordFlagsGetArgs() {
	retry := flag.Int("retry",-1,"Defines max retry count")
	var logPrefix string
	flag.StringVar(&logPrefix,"prefix","","Logger prefix")
	var arr ArrayValue
	flag.Var(&arr,"array","input array to iterate through.")
	flag.Parse()
	logger := log.New(os.Stdout,logPrefix,log.Ldate)

	retryCount := 0
	for retryCount<*retry {
		logger.Println("retrying connection")
		logger.Println("Sending array %v\n",arr)
		retryCount++
	}
}

func (s *ArrayValue) String() string {
	return fmt.Sprintf("%v", *s)
}

func (a *ArrayValue) Set(s string) error {
	*a = strings.Split(s,",")
	return nil
}

