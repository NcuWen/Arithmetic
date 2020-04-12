/**
 * @Author: Ncuwen
 * @Description:
 * @File:  Sort
 * @Version: 1.0.0
 * @Date: 2019/10/12 11:01
 */

package sorts

import (
	"arithmetics/searchs"
)

/**
 * @title	BubbleSort
 * @description	冒泡排序升序算法
 * 优点：简单易懂, 传入切片，利用引用传值，开销较小
 * 缺点：时间复杂度高，数组大时，时间开销大
 * 空间复杂度: O(n)
 * 时间复杂度：O(n^2)
 * @auth: ncuwen
 * @param: point	需要进行排序的切片
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
 * @title	DirectInsertionSortWithReturn
 * @description	直接插入排序升序算法(以空间节省时间)
 * 优点: 实现简单,新建切片存储排序后结果，能节省数组元素移动的时间
 * 缺点: 增大空间开销
 * 空间复杂度: O(2n)
 * 时间复杂度: O(n^2)
 * @auth: ncuwen
 * @param: point	需要进行排序的切片
 * @Return: 返回排序完成的切片
 */
func DirectInsertionSortWithReturn(point []float32) []float32 {
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
 * @title	DirectInsertionSortNoReturn
 * @description	直接插入排序升序算法(以时间节省空间)
 * 优点: 实现简单, 直接在原切片上进行处理，能节省空间
 * 缺点: 增大时间开销，需要移动元素
 * 空间复杂度: O(n)
 * 时间复杂度: O(n^2)
 * @auth: ncuwen
 * @param: point	需要进行排序的切片
 */
func DirectInsertionSortNoReturn(point []float32) {
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
 * @title	BinaryInsertionSort
 * @description	二分插入排序升序算法(以空间节省时间)
 * 优点: 实现简单,新建切片存储排序后结果，能节省数组元素移动的时间，并且使用二分递归查找方式寻找插入位置，节省空间
 * 缺点: 增大空间开销
 * 空间复杂度: O(2n)
 * 时间复杂度: O(n^2)
 * @auth: ncuwen
 * @param: point	需要进行排序的切片
 */
func BinaryInsertionSort(point []float32) []float32 {
	var result = make([]float32, 0, len(point))
	for _, v := range point {
		index := searchs.BinarySearchFirstGreaterByRecursive(result, v)
		temp := append([]float32{}, result[index:]...)
		result = append(result[:index], v)
		result = append(result, temp...)
	}
	return result
}

/**
 * @title	merge
 * @description 归并算法, 将两个已经有序的子序列归并为一个有序的序列
 * 空间复杂度: O(2n)
 * 时间复杂度: O(n)
 * @auth: ncuwen
 * @param: left	需要进行归并的左子序列
 * @param: right	需要进行归并的右子序列
 */
func merge(left, right []float32) []float32 {
	result := make([]float32, 0)
	i, j := 0, 0 // left和right的index位置
	lSize, rSize := len(left), len(right)
	for i < lSize && j < rSize {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
			continue
		}
		result = append(result, right[j])
		j++
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	return result
}

/**
 * @title	MergeSort
 * @description 归并排序算法(递归)
 * 空间复杂度: O(2n)
 * 时间复杂度: O(n)
 * @auth: ncuwen
 * @param: point	需要进行排序的切片
 * @return: 已排序的切片
 */
func MergeSort(point []float32) []float32 {
	if len(point) < 2 {
		return point
	} else {
		i := len(point) / 2
		left := MergeSort(point[0:i])
		right := MergeSort(point[i:])
		result := merge(left, right)
		return result
	}
}
