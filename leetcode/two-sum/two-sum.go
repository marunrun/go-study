package main

import "fmt"

func main() {
	sum := twoSum([]int{2, 7, 11, 15}, 18)

	fmt.Printf("%v", sum)

}

func twoSum(nums []int, target int) []int {

	length := len(nums)

	for i := 0; i < length; i++ {
		for x := i + 1; x < length; x++ {
			if nums[i]+nums[x] == target {
				return []int{i, x}
			}
		}
	}

	return []int{0, 0}
}

func towSum2(nums []int, target int) []int {

	hashtable := make(map[int]int)

	for i, x := range nums {
		if p, ok := hashtable[target-x]; ok {
			return []int{p, i}
		}
		hashtable[x] = i
	}

	return nil

}
