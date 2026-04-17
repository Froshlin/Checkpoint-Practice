package main

import (
	"fmt"
	"unicode"
)

func camelToSnake(s string) string {
	result := ""

	for i, char := range s {
		if unicode.IsUpper(char) {
			if i != 0 {
				result = result + "_"
			}
			result += string(unicode.ToLower(char))
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(camelToSnake("HelloWorld")) // hello_world
}
