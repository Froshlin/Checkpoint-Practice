package main

import (
	"unicode"
)

func CheckNumber(arg string) bool {
	for _, num := range arg {
		if unicode.IsDigit(num) {
			return true
		}
	}
	return false
}

// func main() {
// 	fmt.Println(CheckNumber("eww"))
// }
