package utils

import "fmt"

func IntersectSlices(slice1, slice2 []int) ([]int, error) {
	if slice1 == nil || slice2 == nil {
		return nil, fmt.Errorf("slices cannot be nil")
	}

	result := []int{}
	i, j := 0, 0

	for i < len(slice1) && j < len(slice2) {
		if slice1[i] < slice2[j] {
			i++
		} else if slice1[i] > slice2[j] {
			j++
		} else {
			result = append(result, slice1[i])
			i++
			j++
		}
	}

	return result, nil
}
