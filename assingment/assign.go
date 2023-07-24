package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome")

	num := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range num {
		// fmt.Println(num[i])

		if num[i]%2 == 0 {
			fmt.Println("Even", num[i])
		} else {
			fmt.Println("Odd", num[i])
		}
	}

}
