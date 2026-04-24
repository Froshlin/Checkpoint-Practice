package main

import "fmt"

func CamelToSnakeCase(s string) string {
	if len(s) == 0 {
		return ""
	}

	result := ""

	for i := 0; i < len(s); i++ {
		char := s[i]

		if !(char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z') {
			return s
		}

		if i == len(s)-1 && char >= 'A' && char <= 'Z' {
			return s
		}

		if i > 0 && char >= 'A' && char <= 'Z' && s[i-1] >= 'A' && s[i-1] <= 'Z' {
			return s
		}

		if i > 0 && char >= 'A' && char <= 'Z' {
			result += "_"
		}
		result += string(char)
	}
	return result
}

func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
}
