package main

import (
	"fmt"
)

func main() {
	var input int
	fmt.Print("enter the number\n")
	fmt.Scanln(&input)

	var sum int = 0
	for input > 0 {
		sum += input % 10
		input /= 10

	}

	if sum > 9 {
		var total int = 0
		for sum > 0 {
			total = total + sum%10
			sum = sum / 10

		}
		fmt.Println(total)

	} else {
		fmt.Println(sum)
	}

}
