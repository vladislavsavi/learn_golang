package utils

func AverageOfEvenNumbers(nums []int) int {
	countOfEven := 0
	sum := 0
	for _, v := range nums {
		if v%2 == 0 {
			sum += v
			countOfEven++
		}
	}

	return sum / countOfEven
}
