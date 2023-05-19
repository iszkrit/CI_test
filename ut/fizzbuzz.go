package main

import (
	"fmt"
	"strconv"
)

func FizzBuzz(numbers []int) []string {
	
	var newNumbers []string

	for i := 0; i < len(numbers); i++ {
		n:= numbers[i]
		switch {
		case n%15 == 0:
			newNumbers = append(newNumbers, "FizzBuzz")
		case n%5 == 0:
			newNumbers = append(newNumbers, "Buzz")
		case n%3 == 0:
			newNumbers = append(newNumbers, "Fizz")
		default:
			newNumbers = append(newNumbers, strconv.Itoa(n))

		}
	}	
	return newNumbers

}

func main() {
	var test = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Printf("[]string : %v\n", FizzBuzz(test))
}
