package main

import "fmt"

func RetainFirstHalf(str string) string {
	// Converting it to rune first
	newString := []rune(str)

	// Check the lenght of the string that has been converted to rune then divide it by 2
	finalStr := newString[:len(newString)/2]

	// Now convert the finalString after division back to strings
	return string(finalStr)
}

func main() {
	fmt.Println(RetainFirstHalf("Welcome"))
}
