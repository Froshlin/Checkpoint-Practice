package main

import "fmt"

func LastWord(s string) string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			endWord := i

			for i >= 0 && s[i] != ' ' {
				i--
			}
			return s[i+1 : endWord+1]
		}
	}
	return "\n"
}

func main() {
	fmt.Println(LastWord("Hello world ."))
}
