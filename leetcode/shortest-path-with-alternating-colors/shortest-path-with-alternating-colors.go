package main

import "fmt"

func main() {

	fmt.Println(shortestAlternatingPaths(3, [][]int{
		{0, 1},
	}, [][]int{
		{1, 2},
	}))
}

func shortestAlternatingPaths(n int, redEdges [][]int, blueEdges [][]int) (res []int) {
	if n < 1 {
		return
	}

	res[0] = 0

	// 1 是red  2 是 blue
	allEdges := make([][]int, n)

	for _, redEdge := range redEdges {
		if allEdges[redEdge[0]] == nil {
			allEdges[redEdge[0]] = make([]int, n)
		}
		allEdges[redEdge[0]][redEdge[1]] = 1
	}

	for _, blueEdge := range blueEdges {
		if allEdges[blueEdge[0]] == nil {
			allEdges[blueEdge[0]] = make([]int, n)
		}
		allEdges[blueEdge[0]][blueEdge[1]] = 2
	}

	for i := 1; i < n; i++ {
		// 直连 直接从0 到 i 最短路径就是 1
		if allEdges[0][i] != 0 {
			res[i] = 1
		} else if res[i-1] == -1 {
			// 如果没有直连
			res[i] = -1

		}

	}

	return
}
