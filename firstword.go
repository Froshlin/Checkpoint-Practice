package main

import "fmt"

func FirstWord(s string) string {
	i := 0

	for i < len(s) && s[i] == ' ' {
		i++
	}

	if i == len(s) {
		return "\n"
	}

	start := i

	for i < len(s) && s[i] != ' ' {
		i++
	}
	return s[start:i]
}

func main() {
	fmt.Println(FirstWord("Hello world. When will a new world start"))
}
