package main

import (
	"errors"
	"fmt"
)

func add(a, b int64) int64 {
	return a + b
}

func divide(a, b int64) (float64, error) {
	var err error
	if b == 0 {
		err = errors.New("divide by zero")
	}
	return float64(a) / float64(b), err
}

func main() {
	var a [2]int64

	fmt.Println("Input 2 numbers: ")

	n, err := fmt.Scan(&a[0], &a[1])

	if err != nil {
		fmt.Println(err)
	}

	if n != 2 {
		fmt.Println("Invalid No of Inputs")
		return
	}

	fmt.Println("Addition: ", add(a[0], a[1]))

	div, err := divide(a[0], a[1])
	if err != nil {
		fmt.Println("Error occurred: ", err)
	} else {
		fmt.Println("Division: ", div)
	}

}
