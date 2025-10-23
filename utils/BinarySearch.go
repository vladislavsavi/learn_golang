package utils

func BinarySearch(s []int, search_value int) (index int) {
	index = -1

	lower_bound := 0
	upper_bound := len(s) - 1

	for lower_bound <= upper_bound {
		mid_point := (upper_bound + lower_bound) / 2
		switch {
		case s[mid_point] == search_value:
			return mid_point
		case s[mid_point] < search_value:
			lower_bound = mid_point + 1
		case s[mid_point] > search_value:
			upper_bound = mid_point - 1
		}
	}
	return index
}
