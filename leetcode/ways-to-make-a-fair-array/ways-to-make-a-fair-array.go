package main

import "fmt"

func main() {
	fmt.Println(waysToMakeFair([]int{2, 1, 6, 4}))
}

func waysToMakeFair(nums []int) int {

	// 给你一个整数数组 nums 。你需要选择 恰好 一个下标（下标从 0 开始）并删除对应的元素。请注意剩下元素的下标可能会因为删除操作而发生改变。

	var oddSum, evenSum, odd1, even1, res int

	// 先遍历一次 算出奇数 与偶数的总和
	for i, num := range nums {
		if i&1 > 0 {
			oddSum += num
		} else {
			evenSum += num
		}
	}

	// 再遍历一次
	for i, num := range nums {

		// 奇数 或者 偶数 减去当前值 (相当于删掉当前下标)
		if i&1 > 0 {
			oddSum -= num
		} else {
			evenSum -= num
		}

		if oddSum+even1 == evenSum+odd1 {
			res++
		}

		if i&1 > 0 {
			odd1 += num
		} else {
			even1 += num
		}

	}
	return res
}
