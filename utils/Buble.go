package utils

// O(n^2)
func BubleSort(s []int) {
	n := len(s) - 1
	swapped := true

	for i := 0; i < n; i++ {
		for j := 0; j < n-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}
