package BinarySearch

func BSearch(nums []int, target int) int {
	ptr := len(nums) / 2
	start := 0
	finish := len(nums) - 1

	for start <= finish {
		if nums[ptr] == target {
			return ptr
		} else if nums[ptr] > target {
			finish = ptr - 1
			ptr = (start + finish) / 2
		} else {
			start = ptr + 1
			ptr = (start + finish) / 2
		}
	}

	return -1
}
