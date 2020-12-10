package utils

func Min(nums []int) int {
	min, _ := MinWithIndex(nums)
	return min
}
func MinWithIndex(nums []int) (min int, index int) {
	min = nums[0]
	index = 0
	for i, num := range nums {
		if num < min {
			min = num
			index = i
		}
	}
	return min, index
}

func Max(nums []int) int {
	max, _ := MaxWithIndex(nums)
	return max
}
func MaxWithIndex(nums []int) (max int, index int) {
	max = nums[0]
	index = 0
	for i, num := range nums {
		if num > max {
			max = num
			index = i
		}
	}
	return max, index
}
