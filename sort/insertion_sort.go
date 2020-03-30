package sort

import "fmt"

/*
增量排序
*/

func InsertionSort(source []int) {
	for j := 1; j < len(source); j++ {
		key := source[j]
		i := j - 1
		for i >= 0 && source[i] > key {
			source[i+1] = source[i]
			i = i - 1
		}
		source[i+1] = key
	}
	fmt.Println(source)
}
