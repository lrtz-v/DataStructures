/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : tao.lv
#   Email         : 547617251lrtz@gmail.com
#   File Name     : testcgo.go
#   Last Modified : 2018-05-10 11:35
#   Describe      :
#
# ====================================================*/
package main

import "C"
import (
	"fmt"
	"unsafe"
)

func main() {
	var s = "hello"
	C.foo(C.CString(s))

	var i int = 5
	C.bar(unsafe.Pointer(&i))

	var i32 int32 = 7
	var p *uint32 = (*uint32)(unsafe.Pointer(&i32))
	fmt.Println(*p)
}
