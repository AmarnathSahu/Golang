package main

import (
	"fmt"
)

func Program_To_Get_Occurence_Character(input string) string {
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
	return result
}

func Swap_Elemnts_Of_List(list []int, i, j int) []int {
	list[i], list[j] = list[j], list[i]
	return list
}

func main() {
	input := "AABBBDDDD"
	result := Program_To_Get_Occurence_Character(input)
	fmt.Println(result)

	list := []int{1, 2, 3, 4, 5}
	list = Swap_Elemnts_Of_List(list, 0, 1)
	fmt.Println(list)
}
