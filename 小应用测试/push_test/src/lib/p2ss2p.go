package lib

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"unsafe"
)

/**
	指针转换成字符串类型
	@author:wanghongli
	@since:2018/09/27
*/
func p2s(conn *net.Conn) int64 {

	if conn == nil {
		log.Println("conn is nil")
		return 0
	}

	strPointerInt := fmt.Sprintf("%d", unsafe.Pointer(conn))
	i, _ := strconv.ParseInt(strPointerInt, 10, 0)
	return i
}

/**
	字符串类型转换为指针类型
	@author:wanghongli
	@since:2018/09/27
 */
func s2p(s string) (conn *net.Conn) {
	var pointer *net.Conn
	pointer = *(**net.Conn)(unsafe.Pointer(&s))
	return pointer
}
