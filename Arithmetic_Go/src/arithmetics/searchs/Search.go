/**
 * @Author: Ncuwen
 * @Description:
 * @File:  Search
 * @Version: 1.0.0
 * @Date: 2019/10/11 11:01
 */

package searchs

/**
 * @title	BinarySearchFirstGreaterByRecursive
 * @description	二分查找算法: 在指定切片(有序)中查找第一个大于给定值的元素索引
 * 递归实现
 * @Param: point []float32 指定的切片
 * @Param: v float32 给定的值
 * @Return: index int 查找符合条件的元素在切片中的索引
 */
func BinarySearchFirstGreaterByRecursive(point []float32, v float32) int {
	size := len(point)
	if size == 0 {
		// 切片长为空时返回 0
		return 0
	}
	return bsFirstGreaterByRecursive(point, v, 0, size-1)
}

/**
 * @title	bsFirstGreaterByRecursive
 * @description	二分查找算法的递归部分: 在指定切片(有序)中查找第一个大于给定值的元素索引
 * @Param: point []float32 指定的切片
 * @Param: v float32 给定的值
 * @Param: low, high int 查找的指定上下标
 * @Return: index int 查找符合条件的元素在切片中的索引
 */
func bsFirstGreaterByRecursive(point []float32, v float32, low, high int) int {
	if low == high && high != 0 {
		return high + 1
	}
	if low > high {
		// 下标大于上标
		return 0
	}
	mid := (low + high) / 2
	if point[mid] <= v && point[mid+1] >= v {
		return mid + 1
	} else if point[mid] < v {
		return bsFirstGreaterByRecursive(point, v, mid+1, high)
	} else {
		return bsFirstGreaterByRecursive(point, v, low, mid-1)
	}
}

/**
 * @title	BinarySearchEqualByRecursive
 * @description	二分查找算法: 在指定切片(有序)中查找第一个大于给定值的元素索引
 * 递归实现
 * @Param: point []float32 指定的切片
 * @Param: v float32 给定的值
 * @Return: index int 查找符合条件的元素在切片中的索引
 */
func BinarySearchEqualByRecursive(point []float32, v float32) int {
	size := len(point)
	if size == 0 {
		// 切片长为空时返回 -1
		return -1
	}
	return bsEqualByRecursive(point, v, 0, size-1)
}

/**
 * @title	bsEqualByRecursive
 * @description	二分查找算法的递归部分: 在指定切片(有序)中查找等于给定值的元素索引
 * @Param: point []float32 指定的切片
 * @Param: v float32 给定的值
 * @Param: low, high int 查找的指定上下标
 * @Return: index int 查找符合条件的元素在切片中的索引
 */
func bsEqualByRecursive(point []float32, v float32, low, high int) int {
	if low > high {
		// 下标大于上标
		return -1
	}
	mid := (low + high) / 2
	if point[mid] == v {
		return mid
	} else if point[mid] < v {
		return bsEqualByRecursive(point, v, mid+1, high)
	} else {
		return bsEqualByRecursive(point, v, low, mid-1)
	}
}
