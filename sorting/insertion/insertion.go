package insertion

func sort(arr []int) {
	for i := 1; i < len(arr); i++ {
		value := arr[i]
		j := i - 1
		for j >= 0 && arr[j] > value {
			arr[j+1] = arr[j]
			j--
		}
		arr[j+1] = value
	}
}
