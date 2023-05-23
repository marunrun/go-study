package main

import "fmt"

func main() {

	fmt.Println(hardestWorker(10, [][]int{{0, 3}, {2, 5}, {0, 9}, {1, 15}}))
}

func hardestWorker(n int, logs [][]int) int {
	max := logs[0][1]
	uid := logs[0][0]

	length := len(logs)
	for i := 1; i < length; i++ {
		duration := logs[i][1] - logs[i-1][1]

		if duration > max {
			max = duration
			uid = logs[i][0]
		} else if duration == max {
			uid = min(uid, logs[i][0])
		}
	}

	return uid
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
