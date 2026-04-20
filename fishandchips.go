package main

// import "fmt"

func FishAndChips(n int) string {
	if n < 0 {
		return "Invalid input"
	}

	if (n%3 == 0) && (n%5 == 0) {
		return "FishAndChips"
	}

	if n%3 == 0 {
		return "Fish"
	}

	if n%5 == 0 {
		return "Chips"
	}

	return "Na only Garri Dey o"
}

// func main() {
// 	fmt.Println(FishAndChips(15))
// }
