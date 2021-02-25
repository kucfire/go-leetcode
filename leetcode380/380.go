/*
	leetcode tag:380 title:常熟时间插入
*/

package leetcode380

import "math/rand"

// 抄的
type RandomizedSet struct {
	rMap   map[int]int
	rSlice []int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		rMap:   make(map[int]int),
		rSlice: make([]int, 0),
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.rMap[val]; ok {
		return false
	}
	this.rSlice = append(this.rSlice, val)
	this.rMap[val] = len(this.rSlice) - 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	if _, ok := this.rMap[val]; !ok {
		return false
	}

	index := this.rMap[val]
	this.rSlice[index] = this.rSlice[len(this.rSlice)-1]
	this.rMap[this.rSlice[index]] = index
	// 删除队尾，同时从map中删除
	this.rSlice = this.rSlice[:len(this.rSlice)-1]
	delete(this.rMap, val)

	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	if len(this.rSlice) == 0 {
		return -1
	}
	index := rand.Intn(len(this.rSlice))

	return this.rSlice[index]
}
