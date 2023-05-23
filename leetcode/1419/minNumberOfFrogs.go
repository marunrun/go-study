package main

import (
	"fmt"
)

func main() {
	fmt.Println(minNumberOfFrogs("croakcroak"))
}

func minNumberOfFrogs(croakOfFrogs string) int {

	c, r, o, a, k, res := 0, 0, 0, 0, 0, 0

	for _, s := range croakOfFrogs {
		if s == 'c' {
			if k > 0 {
				k--
			} else {
				res++
			}
			c++
		} else if s == 'r' {
			c--
			r++
		} else if s == 'o' {
			r--
			o++
		} else if s == 'a' {
			o--
			a++
		} else if s == 'k' {
			a--
			k++
		}

		if c < 0 || r < 0 || o < 0 || a < 0 {
			break
		}

	}

	if c != 0 || r != 0 || o != 0 || a != 0 {
		return -1
	}

	return res
}
