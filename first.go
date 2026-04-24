package main

// import "fmt"

func First(s string) string {
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

// func main() {
// 	fmt.Println(First("Hello World"))
// }
