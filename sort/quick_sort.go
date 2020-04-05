package sort

import "fmt"

func QuickSort(source []int) {
	lo := 0
	hi := len(source) - 1
	sort(&source, lo, hi)
	fmt.Println(source)
}

func sort(source *[]int, lo int, hi int) {
	if lo >= hi {
		return
	}
	p := partition(source, lo, hi)
	sort(source, lo, p-1)
	sort(source, p+1, hi)

}

func partition(source *[]int, lo int, hi int) int {
	i := lo
	j := hi + 1
	s := (*source)
	for {
		for {
			i++
			if s[i] > s[lo] || i == hi {
				break
			}
		}
		for {
			j--
			if s[j] < s[lo] || j == lo {
				break
			}
		}
		if i >= j {
			break
		}
		exchange(source, i, j)
	}
	exchange(source, lo, j)
	return j
}

func exchange(source *[]int, i int, j int) {
	s := (*source)
	tmpi := s[i]
	tmpj := s[j]
	(*source)[i] = tmpj
	(*source)[j] = tmpi
}
