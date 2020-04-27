package sort

import "fmt"

/*
自下而上的归并算法更加适合链表类数组
*/

func MergeSortBottom(source []int) {

	//merge(&source,0,len(source)/2-1,len(source)-1)
	count := 0
	mergesort(&source, 0, len(source)-1)
	fmt.Println(count)

}

func MergeSortTop(source []int) {
	for sz := 1; ; sz = sz + sz {
		if sz <= len(source) {
			for i := 0; i < len(source)-sz; i = i + sz + 1 {
				mergesort(&source, i, i+sz)
			}
		} else {
			mergesort(&source, 0, len(source)-1)
			break
		}

	}
	fmt.Println(source)

}

func mergesort(source *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}
	mid := (lo + hi) / 2
	mergesort(source, lo, mid)
	mergesort(source, mid+1, hi)
	merge(source, lo, mid, hi)
}

func merge(source *[]int, lo int, mid int, hi int) {
	i := lo
	j := mid + 1
	rs := make([]int, len(*source))
	for k := lo; k <= hi; k++ {
		rs[k] = (*source)[k]
	}
	for k := lo; k <= hi; k++ {
		if i > mid {
			(*source)[k] = rs[j]
			j++
		} else {
			if j > hi {
				(*source)[k] = rs[i]
				i++
			} else {
				if rs[i] < rs[j] {
					(*source)[k] = rs[i]
					//*count=*count+1
					//fmt.Println(*count)
					i++
				} else {

					(*source)[k] = rs[j]
					j++
				}
			}
		}

	}
}
