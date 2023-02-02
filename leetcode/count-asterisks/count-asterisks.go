package main

import "fmt"

func main() {
	fmt.Println(countAsterisks("l|*e*et|c**o|*de|"))
}

func countAsterisks(s string) (res int) {
	flag := true
	for _, v := range s {

		if flag && v == '*' {
			res++
		}

		if v == '|' {
			flag = !flag
		}
	}

	return res
}
