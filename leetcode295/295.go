/*
	leetcode tag:295 title:数据流的中位数
*/

package leetcode295

// 这个答案还有问题没解决
import "container/heap"

type MinHeap struct {
	nums []int
}

func NewMinHeap() *MinHeap {
	return new(MinHeap)
}

func (h MinHeap) Len() int {
	return len(h.nums)
}

func (h *MinHeap) Swap(i, j int) {
	(*h).nums[i], (*h).nums[j] = (*h).nums[j], (*h).nums[i]
}

func (h MinHeap) Less(i, j int) bool {
	return h.nums[i] < h.nums[j]
}

func (h MinHeap) Push(x interface{}) {
	h.nums = append(h.nums, x.(int))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h).nums[h.Len()-1]
	(*h).nums = (*h).nums[:h.Len()-1]
	return x
}

type MedianFinder struct {
	minHeap *MinHeap
	maxHeap *MinHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	return MedianFinder{
		minHeap: NewMinHeap(),
		maxHeap: NewMinHeap(),
	}
}

func (this *MedianFinder) AddNum(num int) {
	if this.minHeap.Len() > 0 && num > (this.minHeap.nums)[0] {
		heap.Push(this.minHeap, num)
	} else {
		heap.Push(this.maxHeap, -num)
	}

	if this.minHeap.Len()-this.maxHeap.Len() == 2 {
		heap.Push(this.maxHeap, -(heap.Pop(this.minHeap)).(int))
	} else if this.maxHeap.Len()-this.minHeap.Len() == 2 {
		heap.Push(this.minHeap, -(heap.Pop(this.maxHeap)).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.minHeap.Len() > this.maxHeap.Len() {
		return float64(this.minHeap.nums[0])
	} else if this.minHeap.Len() < this.maxHeap.Len() {
		return -float64(this.maxHeap.nums[0])
	}
	return float64(this.minHeap.nums[0]-this.maxHeap.nums[0]) / float64(2)
}
