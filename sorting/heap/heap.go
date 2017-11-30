package heap

func heapify(arr []int, n, root int) {
	largest := root
	l := 2 * root + 1
	r := 2 * root + 2

	if (l < n && arr[l] > arr[largest]) {
		largest = l
	}

	if (r < n && arr[r] > arr[largest]) {
		largest = r
	}

	if (largest != root) {
		arr[root], arr[largest] = arr[largest], arr[root]
		heapify(arr, n, largest)
	}
}

func sort(arr []int) {
	n := len(arr)
	for i := n / 2 - 1; i >= 0; i-- {
		heapify(arr, n, i);
	}

	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0);
	}
}