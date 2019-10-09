package main

import (
	"fmt"
	"simple_arithmetic"
)

func main() {
	balance := []float32{1000.0, 5.0, 3.4, 7.0, 50.0}
	balance = simple_arithmetic.DirectInsertionSort1(balance)
	fmt.Println(balance)
}
