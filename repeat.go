package main

// import "fmt"

func Repeat(s string) string {
	result := ""

	for i := 0; i < len(s); i++ {
		repInd := s[i]
		if repInd >= 'a' && repInd <= 'z' {
			for j := 0; j < int(repInd-'a')+1; j++ {
				result += string(repInd)
			}
		} else if repInd >= 'A' && repInd <= 'Z' {
			for j := 0; j < int(repInd-'A')+1; j++ {
				result += string(repInd)
			}
		} else {
			result += string(repInd)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(Repeat("AbcdeFGh"))
// }
