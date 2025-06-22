package main

import (
	"fmt"
	"math"
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

func Find_Longest_And_Smallest_Words_In_List(list []string) (longestWord string, smallestWord string) {
	maximum, minimum := math.MinInt, math.MaxInt
	for _, value := range list {

		if len(value) >= maximum {
			maximum = len(value)
			longestWord = value
		}

		if len(value) <= minimum {
			minimum = len(value)
			smallestWord = value
		}
	}
	return longestWord, smallestWord
}

func main() {
	input := "AABBBDDDD"
	result := Program_To_Get_Occurence_Character(input)
	fmt.Println(result)

	list := []int{1, 2, 3, 4, 5}
	list = Swap_Elemnts_Of_List(list, 0, 1)
	fmt.Println(list)

	listOfWords := []string{"Vacant", "Testamony", "According"}

	longestWord, smallestWord := Find_Longest_And_Smallest_Words_In_List(listOfWords)

	fmt.Printf("Longest word %v and Smallest word %v ", longestWord, smallestWord)
}
