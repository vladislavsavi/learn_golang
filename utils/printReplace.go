package utils

import "fmt"

func PrintReplaced(text string) {
	runes := []rune(text)

	for index, r := range runes {
		if r == 'у' {
			runes[index] = 'а'
		}
	}
	fmt.Println(string(runes))
}
