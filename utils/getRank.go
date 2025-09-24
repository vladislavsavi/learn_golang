package utils

import "math"

func GetRank(num int) int {
	if num == 0 {
		return 1
	}

	absNum := math.Abs(float64(num))

	return int(math.Log10(absNum)) + 1
}
