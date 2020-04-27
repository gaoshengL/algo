package main

import (
	"algo/sort"
	"fmt"
)

func main() {
	//relation := [][]int{{0,2},{2,1},{3,4},{2,3},{1,4},{2,0},{0,4}}
	//fmt.Println(search.NumberOfSubarrays([]int{1,1,1,1,1},1))
	//fmt.Println(search.NumberOfSubarrays([]int{2,2,2,1,2,2,1,2,2,2},2))
	//fmt.Println(search.NumberOfSubarrays([]int{1,1,2,1,1},3))
	//nums:=[]int{1,-2,-3}
	//s:="1234"
	//fmt.Println(s[1:5])
	//fmt.Println(getMaxRepetitions("baba",11,"baab",1))
	//fmt.Println(getMaxRepetitions("lovelivenanjomusicforever",100000,"lovelivenanjomusicforever",100000))
	//fmt.Println(getMaxRepetitions("lovelive",1000,"lovelive",999))
	//fmt.Println(cal("baab",4))
	sort.MergeSortBottom([]int{7, 5, 6, 4})

}

func cal(s string, n int) string {
	sy := s
	for i := 1; i < n; i = i * 10 {
		if i*10 > n {
			st := ""
			for k := i; k < n; k++ {
				st = st + s
			}
			sy = sy + st
		} else {

			sy = sy + sy
		}
	}
	return sy
}

func getMaxRepetitions(s1 string, n1 int, s2 string, n2 int) int {
	S1 := ""
	S2 := ""
	S1 = cal(s1, n1)
	S2 = cal(s2, n2)
	fmt.Println(len(S1) / len(s1))
	fmt.Println(len(S2) / len(s2))
	n := int(len(S1) / len(S2))
	for j := n; j >= 1; j-- {
		S3 := ""
		for i := 0; i < j; i++ {
			S3 = S3 + S2
		}
		num := len(S3)
		for nu := num; nu >= 1; nu-- {
			flag2 := true
			idx := 0
			for o := 0; o < len(S3); o = o + nu {
				flag := false
				k := o + nu
				if o+nu > len(S3) {
					k = len(S3)
				}
				for i := idx; i < len(S1); i = i + nu {
					g := i + nu
					if i+nu >= len(S1) {
						g = len(S1)
					}
					x := S3[o:k]
					y := S1[i:g]
					if x == y {
						flag = true
						idx = i + nu
						break
					}
				}
				if flag {
					continue
				} else {
					flag2 = false
					break
				}
			}
			if flag2 {
				return j
			}

		}

	}
	return 0
}

func findMinFibonacciNumbers(k int) int {
	f := []int{1, 1}
	i := 0
	j := k
	z := j
	for {
		if j = find(f, z); j <= z {
			z = z - j
			i++
		} else {
			break
		}

	}
	return i

}

func find(nums []int, k int) int {
	if nums[len(nums)-1]+nums[len(nums)-2] > k {
		return nums[len(nums)-1]
	} else {
		if nums[len(nums)-1]+nums[len(nums)-2] == k {
			return nums[len(nums)-1] + nums[len(nums)-2]
		}
	}
	nums = append(nums, nums[len(nums)-1]+nums[len(nums)-2])
	return find(nums, k)
}

func numWays(n int, relation [][]int, k int) int {
	rs := 0
	for _, infonew := range relation {
		if infonew[0] == 0 {
			way(n, relation, infonew, k, &rs)
		}
	}
	return rs
}

func way(n int, relation [][]int, info []int, k int, rs *int) {
	if k == 1 {
		if info[1] == n-1 {
			*rs = *rs + 1
		}
		return
	}
	for _, infonew := range relation {
		if infonew[0] == info[1] {
			way(n, relation, infonew, k-1, rs)
		}
	}
}
