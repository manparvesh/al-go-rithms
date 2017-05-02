package utils

import (
	"io/ioutil"
	"path/filepath"
	"strconv"
	"strings"
)

// GetArrayOfSize returns array of given size
func GetArrayOfSize(n int) []int {
	var ar = make([]int, n)
	absPath, _ := filepath.Abs("../utils/Numbers.txt")
	file, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err)
	}

	str := string(file)
	var sar = strings.Split(str, "\n")

	for i := 0; i < n; i++ {
		x, err2 := strconv.Atoi(sar[i])
		if err2 != nil {
			panic(err2)
		}
		ar[i] = x
	}

	return ar
}
