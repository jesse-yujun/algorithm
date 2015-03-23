// test
package main

import (
	"fmt"
)

func main() {
	src := []int{1, 2, 2, 2, 3, 5, 5, 7}
	for i := 0; i < len(src); i++ {
		fmt.Println(src[i])
	}
	fmt.Println("begin")
	head := 0
	next := 1
	for next < len(src) {
		if src[head] == src[next] {
			next++
		} else {
			src[head+1] = src[next]
			head++
			next++
		}
	}

	fmt.Println(head)
	fmt.Println("end")

	for i := 0; i < len(src); i++ {
		fmt.Println(src[i])
	}
}
