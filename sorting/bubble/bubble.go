package bubble

func sort(arr {}int){
	for item_count := len(arr) - 1; ; item_count-- {
		swap := false
		for i:= 1; i<=item_count; i++ {
			if arr[i-1] > arr[i] {
				arr[i-1], arr[i] = arr[i], arr[i-1]
				swap = true
			}
		}
		if swap == false {
			break;
		}
	}
}