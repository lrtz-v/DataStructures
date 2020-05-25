package main

import "fmt"

func main() {
	a := make([]int, 3, 3)
	a[0] = 0
	a[1] = 1
	fmt.Println(len(a))
}
