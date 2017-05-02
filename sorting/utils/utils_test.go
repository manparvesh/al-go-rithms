package utils

import (
	"fmt"
	"testing"
)

func TestRead(t *testing.T) {
	for i := 500; i <= 1e5; i += 500 {
		array := GetArrayOfSize(i)
		if len(array) != i {
			fmt.Println(array)
			t.Error()
		}
	}
}
