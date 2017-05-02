package insertion

import (
	"fmt"
	"testing"

	"github.com/manparvesh/al-go-rithms/sorting/utils"
)

// TestInsertionSort tests bubble sort
func TestInsertionSort(t *testing.T) {
	list := utils.GetArrayOfSize(1e1)
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
