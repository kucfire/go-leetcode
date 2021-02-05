package leetcode384

import (
	"math/rand"
	"time"
)

type SolutionReview struct {
	nums []int
}

func ConstructorReview(nums []int) SolutionReview {
	newSolution := SolutionReview{nums: nums}
	return newSolution
}

/** Resets the array to its original configuration and return it. */
func (this *SolutionReview) ResetReview() []int {
	return this.nums
}

/** Returns a random shuffling of the array. */
func (this *SolutionReview) ShuffleReview() []int {
	rand.Seed(time.Now().UnixNano())
	copyid := this.copyReview(this.nums)
	for i := len(copyid); i > 0; i-- {
		index := rand.Intn(i)
		copyid[index], copyid[i-1] = copyid[i-1], copyid[index]
	}
	return copyid
}

func (this *SolutionReview) copyReview(nums []int) []int {
	copyid := make([]int, len(nums))
	for i, num := range nums {
		copyid[i] = num
	}
	return copyid
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
