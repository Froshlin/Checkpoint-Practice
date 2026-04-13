package main

// import "fmt"

func PrintIfNot(str string) string {
	if len(str) != 3 {
		return "Passed"
	}
	return "invalid"
}

// func main() {
// 	fmt.Println(PrintIfNot("i"))
// }
