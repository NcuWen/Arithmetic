/**
 * @Author: 86938
 * @Description:
 * @File:  MaxSubarray
 * @Version: 1.0.0
 * @Date: 2020/4/19 13:35
 */
package divide_and_conquer

import "arithmetics/constants"

/**
 * @title	findMaxCrossingSubarray
 * @description	寻找最大子数组问题中的寻找过中点的最大子数组函数
 * 总体理念: 对于给定的数组, 其中点坐标mid, 即向左寻找到最大子数组和, 再向右寻找到最大子数组, 返回左右边界及整个数组即可
 * 空间复杂度: O(n)
 * 时间复杂度: O(n)
 * @auth: ncuwen
 * @param: A	原数组
 * @param: low	数组起始下标
 * @param: high	数组末尾下标
 * @param: mid	数组中点下标, 有 low <= mid <= high
 * @return: 返回一个下标元组划分跨越中点的最大子数组的边界, 并返回最大子数组中值的和
 */
func findMaxCrossingSubarray(A []int, low int, mid int, high int) (int, int, int) {
	leftSum := constants.INT_MIN
	maxLeft := mid
	sum := 0
	for i := mid; i >= low; i-- {
		sum += A[i]
		if sum > leftSum {
			leftSum = sum
			maxLeft = i
		}
	}
	rightSum := constants.INT_MIN
	maxRight := mid + 1
	sum = 0
	for j := mid + 1; j <= high; j++ {
		sum += A[j]
		if sum > rightSum {
			rightSum = sum
			maxRight = j
		}
	}
	return maxLeft, maxRight, leftSum + rightSum
}

/**
 * @title	FindMaxSubarray
 * @description	寻找最大子数组
 * 总体理念: 分治思想
 * 空间复杂度: O(n)
 * 时间复杂度: O(nlg(n))
 * @auth: ncuwen
 * @param: A	原数组
 * @param: low	数组起始下标
 * @param: high	数组末尾下标
 */
func FindMaxSubarray(A []int, low int, high int) (int, int, int) {
	if high == low {
		return low, high, A[low]
	} else {
		mid := (low + high) / 2
		leftLow, leftHigh, leftSum := FindMaxSubarray(A, low, mid)
		rightLow, rightHigh, rightSum := FindMaxSubarray(A, mid+1, high)
		crossLow, crossHigh, crossSum := findMaxCrossingSubarray(A, low, mid, high)
		if leftSum >= rightSum && leftSum >= crossSum {
			return leftLow, leftHigh, leftSum
		} else {
			if rightSum >= crossSum {
				return rightLow, rightHigh, rightSum
			} else {
				return crossLow, crossHigh, crossSum
			}
		}
	}
}
