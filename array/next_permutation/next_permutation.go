package nextPermutation

// nextPermutation implementation of next_permutation as in C++ STL
func nextPermutation(arr []int) bool {
	if len(arr) < 2 {
		return false
	}

	// eg:  1 3 5 4 7 8
	// ans: 1 3 5 7 4 8

	// iterate from right and find the position at which the numbers stop decreasing
	pivot := len(arr) - 1
	for pivot > 0 && arr[pivot-1] >= arr[pivot] {
		pivot--
	}
	// ar: 1 3 5 4 7 8
	//       |

	// if reached the end (already sorted)
	if pivot == 0 {
		return false
	}

	// find successor to the pivot
	successor := len(arr) - 1
	for arr[successor] <= arr[pivot-1] {
		successor--
	}
	// ar: 1 3 5 4 7 8
	//           |

	// swap pivot and successor
	arr[pivot-1], arr[successor] = arr[successor], arr[pivot-1]
	// ar: 1 3 8 4 7 5

	// reverse the latter part of array
	successor = len(arr) - 1
	for successor > pivot {
		arr[pivot], arr[successor] = arr[successor], arr[pivot]
		pivot, successor = pivot+1, successor-1
	}
	// ar: 1 3 5 7 4 8
	return true
}
