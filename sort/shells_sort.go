package sort

import "fmt"

/*
递减增量排序
*/

func ShellsSort(source []int) {
	h := 1
	for h < len(source)/3 {
		h = h*3 + 1
	}
	for h >= 1 {
		for i := h; i < len(source); i = i + 1 {
			key := source[i]
			j := i - 1
			for ; j >= h-1 && source[j] > key; j = j - h {
				source[j+1] = source[j]
			}
			source[j+1] = key

		}
		h = h / 3
	}
	fmt.Println(source)
}
