package algorithm

// QuickSort recursively sorts an array using the quicksort algorithm
func QuickSort(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[len(arr)/2]
	var left, right, equal []int

	for _, num := range arr {
		if num < pivot {
			left = append(left, num)
		} else if num > pivot {
			right = append(right, num)
		} else {
			equal = append(equal, num)
		}
	}

	return append(append(QuickSort(left), equal...), QuickSort(right)...)
}
