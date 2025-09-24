package utils

import "errors"

func FindLongest(slice []int) (result int, err error) {
	sliceLength := len(slice)
	if sliceLength == 0 {
		return 0, errors.New("entry slice is empty")
	}

	result = slice[0]
	maxRank := GetRank(slice[0])

	for i := 1; i < sliceLength; i++ {
		currentRank := GetRank(slice[i])

		if currentRank > maxRank {
			result = slice[i]
			maxRank = currentRank
		}
	}

	return
}
