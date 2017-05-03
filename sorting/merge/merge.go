package merge

func sort(arr []int) {
	if len(arr) < 2 {
		return
	}

	mid := len(arr) / 2

	// divide and conquer
	sort(arr[:mid])
	sort(arr[mid:])

	// if last of first part is smaller than first for second part, it means this part is sorted.
	if arr[mid-1] <= arr[mid] {
		return
	}

	var s = make([]int, len(arr)/2)
	copy(s, arr[:mid])

	left, right := 0, mid

	for i := 0; ; i++ {
		if s[left] <= arr[right] {
			arr[i] = s[left]
			left++

			if left == mid {
				break
			}
		} else {
			arr[i] = arr[right]
			right++

			if right == len(arr) {
				copy(arr[i+1:], s[left:mid])
				break
			}
		}
	}
}
