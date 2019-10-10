package main

import (
	"fmt"
	"simple_arithmetic"
)

func main() {
	balance := []float32{1000.0, 5.0, 3.4, 7.0, 50.0, 8, 2}
	balance = simple_arithmetic.DirectInsertionSort1(balance)
	//balance = simple_arithmetic.BinaryInsertionSort(balance)
	//simple_arithmetic.DirectInsertionSort2(balance)
	//fmt.Println(simple_arithmetic.BinarySearchByRecursive(balance, 10))
	fmt.Println(balance)
}
