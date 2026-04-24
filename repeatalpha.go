package main

import "fmt"

func RepeatAlpha(arg string) string {
	result := ""

	for i := 0; i < len(arg); i++ {
		char := arg[i]

		if char >= 'a' && char <= 'z' {
			for j := 0; j < int(char-'a')+1; j++ {
				result += string(char)
			}
		} else if char >= 'A' && char <= 'Z' {
			for j := 0; j < int(char-'A')+1; j++ {
				result += string(char)
			}
		} else {
			result += string(char)
		}
	}
	return result
}

func main() {
	fmt.Println(RepeatAlpha("AbcdefghI"))
}
