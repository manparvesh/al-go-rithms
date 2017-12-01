package bucket

import (
	"fmt"
	"testing"

	"github.com/manparvesh/al-go-rithms/sorting/utils"
)

// TestBucketSort tests bucket sort
func TestBucketSort(t *testing.T) {
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

func benchmarkBucketSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		sort(list)
	}
}

func BenchmarkBucketSort100(b *testing.B)    { benchmarkBucketSort(100, b) }
func BenchmarkBucketSort1000(b *testing.B)   { benchmarkBucketSort(1000, b) }
func BenchmarkBucketSort10000(b *testing.B)  { benchmarkBucketSort(10000, b) }
func BenchmarkBucketSort100000(b *testing.B) { benchmarkBucketSort(100000, b) }
