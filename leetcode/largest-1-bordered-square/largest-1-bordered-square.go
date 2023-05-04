package main

import (
	"fmt"
)

func main() {

	fmt.Println(largest1BorderedSquare([][]int{
		{1, 1, 1}, {1, 0, 1}, {1, 1, 1},
	}))
}

func largest1BorderedSquare(grid [][]int) int {

	// 顶点 0,0 往右 往下

	// 行高
	row := len(grid)
	if row < 1 {
		return 0
	}
	// 列
	col := len(grid[0])

	for x := 0; x < col; x++ {
		y := x
		if grid[x][y] == 1 {

		}

	}

	return 0
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
