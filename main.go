package main

import (
	"fmt"
	"strconv"
)

func main() {
	conv(123)
}

func conv(num int) {
	var s string = ""

	for num > 0 {
		s = strconv.Itoa(num%10) + s
		num = num / 10
	}
	fmt.Printf("%[1]T, %[1]s\n", s)
}
