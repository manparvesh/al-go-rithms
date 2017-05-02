package bubble

import (
	"fmt"
	"testing"

	"../utils"
)

// TestBubbleSort tests bubble sort
func TestBubbleSort(t *testing.T) {
	list := utils.GetArrayOfSize(10)
	sort(list)
	fail := false
	for i := 0; i < len(list)-2; i++ {
		if list[i] > list[i+1] {
			fmt.Println("Error!")
			// Output: Error!
			t.Error()
			fail = true
		}
	}
	if !fail {
		fmt.Println("Success!")
		// Output: Success!
	}
}
