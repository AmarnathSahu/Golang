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

func GeneratePermutations(word string, prefix string) {
	if len(word) == 0 {
		fmt.Printf("%v ", prefix)
	} else {
		for index, character := range word {
			remaining := word[0:index] + word[index+1:]
			GeneratePermutations(remaining, prefix+string(character))
		}
	}
}

func ReverseString(str string) string {
	runes := []rune(str)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func GetIntersectionFromTwoList(list1 []int, list2 []int) []int {
	var intersectionList []int
	lookup := make(map[int]struct{})

	for _, value := range list1 {
		lookup[value] = struct{}{}
	}

	for _, value := range list2 {
		if _, exists := lookup[value]; exists {
			intersectionList = append(intersectionList, value)
		}
	}

	return intersectionList
}

func FirstMaximumAndSecondMaxiMum(list []int) (int, int) {
	var firstMaximum, secondMaximum int
	for _, value := range list {
		if value > firstMaximum {
			secondMaximum = firstMaximum
			firstMaximum = value
		} else if secondMaximum < value {
			secondMaximum = value
		}
	}

	return firstMaximum, secondMaximum
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
	GeneratePermutations(str, "")

	str = "ABCD"
	reverserString := ReverseString(str)
	fmt.Printf("\nHere is the %v, the generated reverse of %v \n", reverserString, str)

	list1, list2 := []int{1, 2, 5, 6, 7}, []int{3, 8, 9, 5, 4}
	intersectionList := GetIntersectionFromTwoList(list1, list2)
	fmt.Printf("Intersection of two lists %v and %v is %v \n", list1, list2, intersectionList)

	firstMaximum, secondMaximum := FirstMaximumAndSecondMaxiMum(list1)
	fmt.Printf("List %v ::: First maximum value %v and Second maximum value %v \n", list1, firstMaximum, secondMaximum)
	firstMaximum, secondMaximum = FirstMaximumAndSecondMaxiMum(list2)
	fmt.Printf("List %v ::: First maximum value %v and Second maximum value %v \n", list2, firstMaximum, secondMaximum)

}
