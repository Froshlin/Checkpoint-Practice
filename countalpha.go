package main

import (
	/*"fmt"*/
	"unicode"
)

func CountAlpha(s string) int {
	count := 0
	for _, alpha := range s {
		if unicode.IsLetter(alpha) {
			count++
		}
	}
	return count
}

// func main() {
// 	fmt.Println(CountAlpha("Oluwatimilehin"))
// }
