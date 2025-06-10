package main

import "fmt"

func main() {
	input := "AABBBDDDD"
	MapData := make(map[string]int)
	for _, ch := range input {
		str := string(ch)
		val, ok := MapData[str]
		if ok {
			val++
		} else {
			val = 1
		}
		MapData[str] = val
	}
	result := ""
	for key, value := range MapData {
		result += fmt.Sprintf("%s%d", key, value)
	}
	fmt.Println(result)
}
