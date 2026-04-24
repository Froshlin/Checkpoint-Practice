package main

import "fmt"

func HashCode(dec string) string {
	newString := ""
	length := len(dec)

	for i := 0; i < length; i++ {
		value := (int(dec[i]) + length) % 127

		if value < 33 {
			value = value + 33
		}
		newString += string(rune(value))
	}
	return newString
}

func main() {
	fmt.Println(HashCode("A"))
	fmt.Println(HashCode("AB"))
	fmt.Println(HashCode("BAC"))
	fmt.Println(HashCode("Hello World"))
}
