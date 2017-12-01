package heap

import (
	"fmt"
	"testing"

	"github.com/manparvesh/al-go-rithms/sorting/utils"
)

// TestHeapSort tests heap sort
func TestHeapSort(t *testing.T) {
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

func benchmarkHeapSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		sort(list)
	}
}

func BenchmarkHeapSort100(b *testing.B)    { benchmarkHeapSort(100, b) }
func BenchmarkHeapSort1000(b *testing.B)   { benchmarkHeapSort(1000, b) }
func BenchmarkHeapSort10000(b *testing.B)  { benchmarkHeapSort(10000, b) }
func BenchmarkHeapSort100000(b *testing.B) { benchmarkHeapSort(100000, b) }
