package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maxTotalFruits([][]int{
		{2, 8},
		{6, 3},
		{8, 6},
	}, 5, 4))
}

func maxTotalFruits(fruits [][]int, startPos int, k int) int {

	n := len(fruits)

	// 左侧边界
	left := sort.Search(n, func(i int) bool {
		return fruits[i][0] >= startPos-k
	})

	right, s := left, 0

	// left 到 startPos 之间的水果
	for ; right < n && fruits[right][0] <= startPos; right++ {
		s += fruits[right][1]
	}
	ans := s
	for ; right < n && fruits[right][0] <= startPos+k; right++ {

		// 加上右侧的水果数量
		s += fruits[right][1]

		// 先往右走,再回到左侧边界
		// 先往左走,再回到右侧边界  如果两边都不行,说明步数超过k 需要减去左侧的水果数量
		for fruits[right][0]*2-fruits[left][0]-startPos > k &&
			fruits[right][0]-fruits[left][0]*2+startPos > k {

			s -= fruits[left][1]
			left++
		}

		ans = max(ans, s)
	}

	return ans
}

func max(a, b int) int {

	if a > b {
		return a
	}
	return b
}
