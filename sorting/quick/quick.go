package quick

func partition(arr []int) int {
	lo := 0
	hi := len(arr)-1
	pivot := arr[hi]

	for j := lo; j < hi; j++ {
		if arr[j] < pivot {
			arr[j], arr[lo] = arr[lo], arr[j]
			lo++
		}
	}

	arr[lo], arr[hi] = arr[hi], arr[lo]
	return lo
}

func sort(arr []int) {
	if len(arr) < 2 {
		return
	}

	pivot := partition(arr)
	sort(arr[:pivot])
	sort(arr[pivot:])
}
