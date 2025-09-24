package main

import (
	"fmt"
	"project/utils"
)

func main() {
	slice := []int{0, 1, 12, 13, 15}

	if result, err := utils.FindLongest(slice); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	if result, err := utils.FindLongest([]int{}); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
