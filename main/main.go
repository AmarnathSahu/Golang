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

func Find_Longest_And_Smallest_Words_In_List(list []string) (string, string) {
	maximum, minimum := math.MinInt, math.MaxInt
	longestWord, smallestWord := "", ""
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

func generatePermutations(word string, prefix string) {
	if len(word) == 0 {
		fmt.Printf("%v ", prefix)
	} else {
		for index, character := range word {
			remaining := word[0:index] + word[index+1:]
			generatePermutations(remaining, prefix+string(character))
		}
	}
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
	fmt.Printf("Longest word is %v And smallest word is %v \n", longestWord, smallestWord)

	str := "ABC"
	fmt.Printf("Below is the generated permutation for %v \n", str)
	generatePermutations(str, "")
}
