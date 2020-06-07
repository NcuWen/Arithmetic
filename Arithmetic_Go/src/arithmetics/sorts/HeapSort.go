/**
 * @Author: 86938
 * @Description:
 * @File:  HeapSort
 * @Version: 1.0.0
 * @Date: 2020/6/7 17:47
 */
package sorts

type Heap struct {
	arr  []float32 // 用来存储堆的数据
	size int       // 用来标识堆的大小
}

/**
 * @title	HeapSort
 * @description 堆排序(原址)
 * 空间复杂度: O(n)
 * 时间复杂度: O(nlgn)
 * @auth: ncuwen
 * @param: point	需要进行排序的切片
 */
func HeapSort(point []float32) {
	heap := buildMaxHeap(point)

	for heap.size > 0 {
		// 将最大的数值调整到堆的末尾
		heap.arr[0], heap.arr[heap.size-1] = heap.arr[heap.size-1], heap.arr[0]
		// 减少堆的长度
		heap.size--
		// 由于堆顶元素改变了，而且堆的大小改变了，需要重新调整堆，维持堆的性质
		maxHeap(heap, 0)
	}
}

/**
 * @title	buildMaxHeap
 * @description 构建最大堆
 * @auth: ncuwen
 * @param: point	需要构建为最大堆的切片
 */
func buildMaxHeap(point []float32) Heap {

	heap := Heap{
		arr:  point,
		size: len(point),
	}
	for i := len(point) / 2; i >= 0; i-- {
		maxHeap(heap, i)
	}

	return heap
}

/**
 * @title	maxHeap
 * @description 最大堆有序化, 由上至下
 * @auth: ncuwen
 * @param: point	原先的堆
 */
func maxHeap(heap Heap, i int) {

	leftNode := i*2 + 1
	rightNode := i*2 + 2
	maxNode := i
	if leftNode < heap.size && heap.arr[leftNode] > heap.arr[maxNode] {
		maxNode = leftNode
	}
	if rightNode < heap.size && heap.arr[rightNode] > heap.arr[maxNode] {
		maxNode = rightNode
	}

	if maxNode != i {
		heap.arr[i], heap.arr[maxNode] = heap.arr[maxNode], heap.arr[i]
		maxHeap(heap, maxNode)
	}
}
