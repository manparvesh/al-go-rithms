package selection

func sort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		min := i

		// get minimum int, staring from i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}

		// if different, swap
		if min != i {
			arr[i], arr[min] = arr[min], arr[i]
		}
	}
	return arr
}
