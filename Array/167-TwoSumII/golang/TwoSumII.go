/*
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2. Please note that your returned answers (both index1 and index2) are not zero-based.

You may assume that each input would have exactly one solution and you may not use the same element twice.

Input: numbers={2, 7, 11, 15}, target=9
Output: index1=1, index2=2
*/

package main

import (
	"fmt"
)

func twoSum(numbers []int, target int) []int {
	tmp := 0
	for i := len(numbers) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			tmp = numbers[i] + numbers[j]
			if tmp == target {
				return []int{j + 1, i + 1}
			} else if tmp > target {
				break
			} else {
				continue
			}
		}
	}
	return nil
}

func main() {
	fmt.Printf("result: %v", twoSum([]int{2, 7, 11, 15}, 9))
}
