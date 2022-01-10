package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1 string
	fmt.Print("enter the number\n")
	fmt.Scanln(&input1)
	input, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Println("Enter the valid input")
	} else {
		for input/10 > 0 {
			input = input/10 + input%10
		}
	}
	fmt.Println("sum in unitdigt", input)

}
