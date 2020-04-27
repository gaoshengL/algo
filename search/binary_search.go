package search

/*
时间复杂度log2n
*/
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	return BinarySearch(nums, 0, len(nums)-1, target)
}

func BinarySearch(nums []int, left int, right int, tartget int) int {
	if left == right {
		if nums[left] == tartget {
			return left
		} else {
			return -1
		}
	}
	mid := left + (right-left)/2
	for i := left; i <= mid; i++ {
		if nums[i] == tartget {
			return i
		}
	}
	return BinarySearch(nums, mid+1, right, tartget)
}
