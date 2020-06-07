package main

import (
	"arithmetics/sorts"
	"fmt"
)

func main() {
	balance := []float32{13, -3, -25, 20, -3, -16, -23, 18, 20, -7, 12, -5, -22, 15, -4, 7}
	//balance = sorts.DirectInsertionSortWithReturn(balance)
	//balance = sorts.BinaryInsertionSort(balance)
	//balance = sorts.MergeSort(balance)
	//fmt.Println(balance)
	//fmt.Println(searchs.BinarySearchEqualByRecursive(balance, 8))
	//low, high, sum := divide_and_conquer.FindMaxSubarray(balance, 0, len(balance) - 1)
	//fmt.Println(low, high, sum)
	//fmt.Println(balance[low: high + 1])
	fmt.Println("原序列： ", balance)
	sorts.HeapSort(balance)
	fmt.Println("新序列： ", balance)
}
