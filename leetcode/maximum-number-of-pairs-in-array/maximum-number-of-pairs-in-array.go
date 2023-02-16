package main

import "fmt"

func main() {
	fmt.Println(numberOfPairs([]int{1, 2, 3, 4, 5}))
}

func numberOfPairs(nums []int) []int {
	cnt := map[int]bool{}
	res := 0
	for _, num := range nums {
		cnt[num] = !cnt[num]
		if !cnt[num] {
			res++
		}
	}
	return []int{res, len(nums) - 2*res}

}
