package utils

import (
	"fmt"
	"math/rand/v2"
)

func RollDice(n int) {
	var firstDice int = 0
	var secondDice int = 0
	var tryCount int = 0

	for {
		firstDice = rand.IntN(6) + 1
		secondDice = rand.IntN(6) + 1
		tryCount++
		if firstDice+secondDice == n {
			fmt.Printf("Выпало %d и %d, в сумме %d, на это потребовалось %d бросков.\n", firstDice, secondDice, n, tryCount)
			break
		} else {
			fmt.Printf("Выпало %d и %d, в сумме %d, бросаем еще раз.\n", firstDice, secondDice, firstDice+secondDice)
		}
	}
}
