package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func PrintMagic(numSlice []int) {
	n := len(numSlice)
	if n == 0 {
		fmt.Println("[]")
		return
	}

	result := make([]int, n)

	// Проход 1: Считаем произведения всех элементов слева
	leftProduct := 1
	for i := 0; i < n; i++ {
		result[i] = leftProduct
		leftProduct *= numSlice[i]
	}

	// Проход 2: Считаем произведения всех элементов справа
	rightProduct := 1
	for i := n - 1; i >= 0; i-- {
		result[i] *= rightProduct
		rightProduct *= numSlice[i]
	}

	// Форматируем и выводим результат
	resultSlice := make([]string, n)
	for i, v := range result {
		resultSlice[i] = strconv.Itoa(v)
	}
	fmt.Println("[" + strings.Join(resultSlice, ", ") + "]")
}
