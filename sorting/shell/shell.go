package shell

func gappedInsertionSort(arr []int, gap int) {
	for i := gap; i < len(arr); i += gap {
		value := arr[i]
		j := i - gap
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = value
	}
}

func sort(arr []int) {
	sublist := len(arr) / 2
	for sublist > 0 {
		for start := 0; start < sublist; start++ {
			gappedInsertionSort(arr, sublist)
		}
		sublist /= 2
	}
}