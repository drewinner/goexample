package main

import "testing"

/**
1,benchmark函数一般以BenchMark开头
2，benchmark的case一般会跑b.N次，而且每次执行都会如此
3，根据实际case的执行时间是否稳定会增加b.N的次数已达到稳定
4，go test -bench=.
*/
func BenchmarkAll(b *testing.B){
	for n:=0;n<b.N;n++{
		Print()
		//aaa(n)
	}
}
/**
稳态状态停止反例
 */
func aaa(n int) int {
	for n>0 {
		n--
	}
	return n
}


