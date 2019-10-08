package simple_sort

/**
 * 冒泡排序算法
 * 优点：简单易懂
 * 缺点：类似于C，没有体现go特性，直接传入数组，开销大
 * 空间复杂度: O(n)
 * 时间复杂度：O(n^2)
 * Author：ncuwen
 */
func BubbleSort(point []float32) {
	var size int = len(point)
	for i := 0; i < size - 1; i++ {
		for j := i + 1; j < size; j++ {
			if point[i] > point[j] {
				point[i], point[j] = point[j], point[i]
			}
		}
	}
}


