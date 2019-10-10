package simple_arithmetic

/**
 * 冒泡排序升序算法
 * 优点：简单易懂, 传入切片，利用引用传值，开销较小
 * 缺点：时间复杂度高，数组大时，时间开销大
 * 空间复杂度: O(n)
 * 时间复杂度：O(n^2)
 * Author：ncuwen
 */
func BubbleSort(point []float32) {
	var size = len(point)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if point[i] > point[j] {
				point[i], point[j] = point[j], point[i]
			}
		}
	}
}

/**
 * 直接插入排序升序算法(以空间节省时间)
 * 优点: 实现简单,新建切片存储排序后结果，能节省数组元素移动的时间
 * 缺点: 增大空间开销
 * 空间复杂度: O(2n)
 * 时间复杂度: O(n^2)
 * Author: ncuwen
 */
func DirectInsertionSort1(point []float32) []float32 {
	var result = make([]float32, 0, len(point))
	for _, v := range point {
		reSize := len(result)
		index := 0
		for j := reSize - 1; j >= 0; j-- {
			if v > result[j] {
				index = j + 1
				break
			}
		}
		temp := append([]float32{}, result[index:]...)
		result = append(result[:index], v)
		result = append(result, temp...)
	}
	return result
}

/**
 * 直接插入排序升序算法(以时间节省空间)
 * 优点: 实现简单, 直接在原切片上进行处理，能节省空间
 * 缺点: 增大时间开销，需要移动元素
 * 空间复杂度: O(n)
 * 时间复杂度: O(n^2)
 * Author: ncuwen
 */
func DirectInsertionSort2(point []float32) {
	var j = 0
	for i := 1; i < len(point); i++ {
		tmp := point[i]
		for j = i; j > 0 && tmp < point[j-1]; j-- {
			point[j] = point[j-1]
		}
		point[j] = tmp
	}
}

/**
 * 二分插入排序升序算法(以空间节省时间)
 * 优点: 实现简单,新建切片存储排序后结果，能节省数组元素移动的时间，并且使用二分递归查找方式寻找插入位置，节省空间
 * 缺点: 增大空间开销
 * 空间复杂度: O(2n)
 * 时间复杂度: O(n^2)
 * Author: ncuwen
 */
func BinaryInsertionSort(point []float32) []float32 {
	var result = make([]float32, 0, len(point))
	for _, v := range point {
		index := BinarySearchByRecursive(result, v)
		temp := append([]float32{}, result[index:]...)
		result = append(result[:index], v)
		result = append(result, temp...)
	}
	return result
}
