package main

import "fmt"

func main() {
	str := "Go爱好者"
	for i, c := range str {
		fmt.Printf("%d: %q %x [% x]\n", i, c, c, []byte(string(c)))
	}
}
