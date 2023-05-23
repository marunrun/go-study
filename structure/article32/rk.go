package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Print(rk("baddefcbasd", "cba"))
}

func rk(s string, search string) int {

	sHash := hash(search, len(search))

	for i := 0; i < len(s)-len(search); i++ {

		if getH(s, i, len(search)) == sHash {
			return i
		}

	}

	return -1
}

var h = make(map[int]int)

func getH(s string, i int, m int) int {
	// s 是字符串 i 下标  m 是长度

	var res int
	var ok bool

	// 兼容 i = 0
	if res, ok = h[i-1]; !ok {
		h[i] = hash(s[i:m], m)
		return h[i]
	}

	// 后面的字串都可以通过前面的字串计算出来

	h[i] = (res-pow(26, m-1)*int(s[i-1]-'a'))*26 + int(s[i+m-1]-'a')

	return h[i]
}

func hash(s string, m int) (hSum int) {
	for i, s := range s {
		hSum += (int(s) - int('a')) * pow(26, m-i-1)
	}
	return
}

func pow(a, b int) int {
	return int(math.Pow(float64(a), float64(b)))
}
