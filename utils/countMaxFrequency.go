package utils

func CountMaxFrequency(slice []int) int {
	numsMap := map[int]int{}
	var mostPopularNum int

	for _, num := range slice {
		_, isExistCurrentNum := numsMap[num]
		_, isExistMostPopular := numsMap[mostPopularNum]

		if isExistCurrentNum {
			numsMap[num]++
		} else {
			numsMap[num] = 1
		}

		if !isExistMostPopular {
			mostPopularNum = num
		}

		if numsMap[mostPopularNum] < numsMap[num] {
			mostPopularNum = num
		}
	}
	return numsMap[mostPopularNum]
}
