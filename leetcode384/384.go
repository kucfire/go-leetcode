/*
	leetcode tag:384 title:打乱数组
*/

package leetcode384

import (
	"math/rand"
	"time"
)

type Solution struct {
	nums []int
}

func Constructor(nums []int) Solution {
	newDefinition := Solution{nums: nums}
	return newDefinition
}

/** Resets the array to its original configuration and return it. */
func (this *Solution) Reset() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *Solution) Shuffle() []int {
	rand.Seed(time.Now().UnixNano())
	copied := this.copy()
	// 随机数逐层递减，
	for i := len(copied); i > 0; i-- {
		index := rand.Intn(i)
		tmp := copied[index]
		copied[index] = copied[i-1]
		copied[i-1] = tmp
	}
	return copied
}

func (this *Solution) Shuffle2() []int {
	res := make([]int, len(this.nums))
	copy(res, this.nums)
	for i := len(res); i > 0; i-- {
		rand := rand.Intn(i)
		res[i-1], res[rand] = res[rand], res[i-1]
	}
	return res
}

func (this *Solution) copy() []int {
	copied := make([]int, len(this.nums))
	for i, v := range this.nums {
		copied[i] = v
	}
	return copied
}
