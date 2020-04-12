package main

import (
	"arithmetics/searchs"
	"arithmetics/sorts"
	"fmt"
)

func main() {
	balance := []float32{1000.0, 5.0, 3.4, 7.0, 50.0, 8, 2}
	//balance = sorts.DirectInsertionSortWithReturn(balance)
	//balance = sorts.BinaryInsertionSort(balance)
	balance = sorts.MergeSort(balance)
	fmt.Println(balance)
	fmt.Println(searchs.BinarySearchEqualByRecursive(balance, 8))
}
