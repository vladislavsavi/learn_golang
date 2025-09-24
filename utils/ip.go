package utils

import (
	"fmt"
	"slices"
	"strings"
)

func InvertMap(entryMap map[string]int) map[int][]string {
	result := make(map[int][]string)

	for key, value := range entryMap {
		if _, ok := result[value]; !ok {
			result[value] = []string{}
		}
		result[value] = append(result[value], key)
	}

	return result
}

func PrintMap(entryMap map[int][]string) {

	var sb strings.Builder
	sb.WriteString("{\n")

	if len(entryMap) > 0 {
		mapKeys := make([]int, len(entryMap))

		i := 0
		for key := range entryMap {
			mapKeys[i] = key
			i++
		}

		slices.Sort(mapKeys)

		for _, mapKey := range mapKeys {
			sb.WriteString(fmt.Sprintf("  %d: [\"%s\"],\n", mapKey, strings.Join(entryMap[mapKey], "\", \"")))
		}
	}
	sb.WriteString("}")

	fmt.Println(sb.String())
}
