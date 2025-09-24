package utils

func DeletingFromSlice(slice []int) []int {
	startLength := len(slice)
	if startLength == 0 {
		return slice
	}

	if slice[startLength-1] > 10 {
		slice = slice[:startLength-1]
	}

	if len(slice) >= 3 && cap(slice) > 5 {
		slice = append(slice[:2], slice[3:]...)
	}

	if startLength == len(slice)+2 {
		slice = slice[1:]
	}

	if len(slice) > 0 {
		slice = slice[:len(slice):len(slice)]
	}

	return slice
}
