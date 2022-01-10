package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input1 string
	fmt.Print("enter the number\n")
	fmt.Scanln(&input1)
	input := random(input1)
	fmt.Println(input)
}
func random(input1 string) int {
	input, err := strconv.Atoi(input1)
	if err != nil {
		fmt.Println("Enter the valid input")
	}
	for input/10 > 0 {
		input = input/10 + input%10
	}
	return input

}
