package main

// import "fmt"

func RepeatAlpha(arg string) string {
	result := ""

	for i := 0; i < len(arg); i++ {
		strInd := arg[i]

		if strInd >= 'a' && strInd <= 'z' {
			for j := 0; j < int(strInd-'a')+1; j++ {
				result += string(strInd)
			}
		} else if strInd >= 'A' && strInd <= 'Z' {
			for j := 0; j < int(strInd-'A')+1; j++ {
				result += string(strInd)
			}
		} else {
			result += string(strInd)
		}
	}
	return result
}

// func main() {
// 	fmt.Println(RepeatAlpha("AbcdefghI"))
// }
