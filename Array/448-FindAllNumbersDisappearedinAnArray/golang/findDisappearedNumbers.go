package findDisappearedNumbers

import "math"

func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		val := int(math.Abs(float64(nums[i]))) - 1
		if nums[val] > 0 {
			nums[val] = -nums[val]
		}
	}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
