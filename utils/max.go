package utils

import (
	"fmt"
)

func Max(nums []int) (int, error) {
	if len(nums) == 0 {
		return 0, fmt.Errorf("slice is nil or empty")
	}

	result := nums[0]

	for _, v := range nums {
		if v > result {
			result = v

		}
	}
	return result, nil
}
