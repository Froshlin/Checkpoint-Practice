package main

// import "fmt"

func PrintIf(str string) string {
	if len(str) == 3 {
		return "Passed"
	}
	return "invalid"
}

// func main() {
// 	fmt.Println(PrintIf("who"))
// }
