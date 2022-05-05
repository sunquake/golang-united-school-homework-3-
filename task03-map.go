package homework

import "sort"

func sortMapValues(input map[int]string) (result []string) {
	sorted := []int{}
	for i:=range input {
		sorted = append(sorted, i)
	}
	sort.Ints(sorted)
	for _, v := range sorted {
		result = append(result, v)
	}
	return
}
