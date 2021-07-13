package sliceutil

// RemoveDuplicateElement remove duplicate elements in slice
func RemoveDuplicateElement(dataSet []int64) []int64 {
	result := make([]int64, 0, len(dataSet))
	temp := map[int64]struct{}{}
	for _, item := range dataSet {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
