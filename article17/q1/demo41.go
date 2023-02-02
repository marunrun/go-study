package main

import "fmt"

func main() {
	// 示例1。
	numbers1 := []int{1, 2, 3, 4, 5, 6}
	for i := range numbers1 {
		if i == 3 {
			numbers1[i] |= i // 4 和 3进行按位或操作得到的结果是 7
		}
	}
	fmt.Println(numbers1)
	fmt.Println()

	// 示例2。
	numbers2 := [...]int{1, 2, 3, 4, 5, 6} //数组 值类型
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	// 7 3 5 7 9 11
	fmt.Println(numbers2)
	fmt.Println()

	// 示例3。
	numbers3 := []int{1, 2, 3, 4, 5, 6} //切片 引用类型 number3[i]的结果值一直在变
	maxIndex3 := len(numbers2) - 1
	for i, e := range numbers3 {
		if i == maxIndex3 {
			numbers3[0] += e
		} else {
			numbers3[i+1] += e
		}
	}
	// 22 3 6 10 15 21
	fmt.Println(numbers3)
}
