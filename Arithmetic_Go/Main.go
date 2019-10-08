package main

import (
	"fmt"
	"simple_sort"
)

func main() {
	var balance = []float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	simple_sort.BubbleSort(balance)
	// for 循环中range对数组进行迭代
	for _,v := range balance {
		fmt.Printf("%.1f  ", v)
	}
}
