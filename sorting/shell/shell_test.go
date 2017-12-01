package shell

import (
	"fmt"
	"testing"

	"github.com/manparvesh/al-go-rithms/sorting/utils"
)

// TestShellSort tests shell sort
func TestShellSort(t *testing.T) {
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

func benchmarkShellSort(n int, b *testing.B) {
	list := utils.GetArrayOfSize(n)
	for i := 0; i < b.N; i++ {
		sort(list)
	}
}

func BenchmarkShellSort100(b *testing.B)    { benchmarkShellSort(100, b) }
func BenchmarkShellSort1000(b *testing.B)   { benchmarkShellSort(1000, b) }
func BenchmarkShellSort10000(b *testing.B)  { benchmarkShellSort(10000, b) }
func BenchmarkShellSort100000(b *testing.B) { benchmarkShellSort(100000, b) }
