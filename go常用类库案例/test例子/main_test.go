package main

import (
	"fmt"
	"testing"
)

/**
作为初始化test,并且使用m.Run来初始化
如果没有在TestMain中调用m.Run，则除了TestMain外的其它tests都不会被执行
 */
func TestMain(m *testing.M) {
	fmt.Println("test main first")
	m.Run()
}
func TestPrint(t *testing.T) {
	res := Print()
	if res != 210 {
		t.Errorf("wrong result of Print")
	}
}

/**
跳过此test case
t.SkipNow函数要放到第一行
 */
func TestPrintSkip(t *testing.T) {
	t.SkipNow()
	res := Print()
	if res != 200 {
		t.Errorf("wrong result of Print2")
	}
}

/**
子test cast
 */
func TestPrintSub(t *testing.T) {
	t.Run("a1", func(t *testing.T) { fmt.Println("a1") })
	t.Run("a2", func(t *testing.T) { fmt.Println("a2") })
	t.Run("a3", func(t *testing.T) { fmt.Println("a3") })
}
