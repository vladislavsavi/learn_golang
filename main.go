package main

import (
	"fmt"
	"project/utils"
)

func main() {
	slice := []int{2, 12, 3, 1, 11, 10}
	fmt.Println(slice)

	utils.BubleSort(slice)
	fmt.Println(slice)

	fmt.Println(utils.BinarySearch(slice, 11))
}
