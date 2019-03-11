/* ====================================================
#   Copyright (C)2018 All rights reserved.
#
#   Author        : tao.lv
#   Email         : 547617251lrtz@gmail.com
#   File Name     : main.go
#   Last Modified : 2018-05-14 11:46
#   Describe      :
#
# ====================================================*/

package main

import (
	"fmt"
	"time"
)

func fac(a, b int) {
	fmt.Println(a, b)
	time.Sleep(time.Second)
	if a+b > 100 {
		return
	}
	fac(b, a+b)
}

func main() {
	fac(1, 1)
}
