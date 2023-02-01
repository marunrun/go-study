package main

import "fmt"

func main() {

	fmt.Println(decodeMessage("the quick brown fox jumps over the lazy dog", "vkbs bs t suepuv"))
}

func decodeMessage(key string, message string) string {

	keyMap := make(map[int32]int32, len(key))
	keyMap[' '] = ' '
	max := 'a'

	for _, v := range key {
		if _, ok := keyMap[v]; !ok {
			keyMap[v] = max
			max++
		}
	}

	var res []byte

	for _, v := range message {
		res = append(res, byte(keyMap[v]))
	}

	return string(res)
}
