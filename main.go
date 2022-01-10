package main

import (
	"fmt"
	"reflect"
)

func main() {
	var s []int
	var input int
	fmt.Print("enter the number\n")
	fmt.Scanln(&input)
	t := reflect.ValueOf(input).Kind()
	fmt.Println(t)
	var sum int = 0
	for input > 0 {
		sum += input % 10
		s = append(s, input%10)
		input /= 10

	}
	for i, j := 0, len(s)-1; i < j; i, j = i+1, 0 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println(s)
	if sum > 9 {
		var total int = 0
		for sum > 0 {
			total += sum % 10
			sum /= 10

		}
		fmt.Println("sum in unidigit", total)

	} else {
		fmt.Println("sum in unidigit", sum)
	}

}
