package utils

import "fmt"

func IsPalindrome(arr [10]int) {
	for i := 0; i < 5; i++ {
		if arr[i] != arr[9-i] {
			fmt.Println("Не палиндром!")
			return
		}
	}

	fmt.Println("Это палиндром!")
}
