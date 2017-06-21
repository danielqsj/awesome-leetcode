package MaxConsecutiveOnes

import "math"

func findMaxConsecutiveOnes(nums []int) int {
	max, count := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			count++
			max = int(math.Max(float64(count), float64(max)))
		} else {
			count = 0
		}
	}
	return max
}
