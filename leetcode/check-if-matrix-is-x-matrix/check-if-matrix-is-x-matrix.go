package main

import "fmt"

func main() {
	fmt.Println(checkXMatrix([][]int{
		{2, 0, 0, 1}, {0, 3, 1, 0}, {0, 5, 2, 0}, {4, 0, 0, 2},
	}))
}

func checkXMatrix(grid [][]int) bool {
	if len(grid) < 1 {
		return false
	}

	// 长度
	n := len(grid[0])

	for i := 0; i < n; i++ {
		x := i
		y := i

		if grid[x][y] == 0 {
			return false
		}

		for x++; x < n; x++ {
			// 对角线
			if x+y == n-1 {
				if grid[x][y] == 0 {
					return false
				}
			} else {
				if grid[x][y] != 0 {
					return false
				}
			}
		}
		x = i

		for y++; y < n; y++ {
			// 对角线
			if x+y == n-1 {
				if grid[x][y] == 0 {
					return false
				}
			} else {
				if grid[x][y] != 0 {
					return false
				}
			}
		}

	}

	return true
}

func checkXMatrix2(grid [][]int) bool {
	for i, row := range grid {
		for j, x := range row {
			if i == j || i+j == len(grid)-1 {
				if x == 0 {
					return false
				}
			} else if x != 0 {
				return false
			}
		}
	}
	return true
}
