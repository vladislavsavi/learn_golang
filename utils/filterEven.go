package utils

func FilterEven(nums ...int) []int {
	i := 0
	for _, v := range nums {
		if v%2 == 0 {

			nums[i] = v
			i++
		}
	}
	return nums[:i]
}
