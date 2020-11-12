/*
	leetcode tag:268 title:丢失的数字
*/

package leetcode268

import "fmt"

// MissingNumber : method of bit operation
func MissingNumber(nums []int) int {
	var result int
	for i, num := range nums {
		result ^= num ^ i
		fmt.Println(result)
	}
	return result ^ len(nums)
}

// MissingNumber2 : normal method
func MissingNumber2(nums []int) int {
	var result int
	for i, num := range nums {
		result ^= num ^ i
		fmt.Println(result)
	}
	return result ^ len(nums)
}
