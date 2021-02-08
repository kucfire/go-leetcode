package leetcode215

import "testing"

func TestFindKthLargest(t *testing.T) {
	datas := []struct {
		nums   []int
		k      int
		result int
	}{
		{[]int{3, 2, 1, 5, 6, 4}, 2, 5},
		{[]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4, 4},
	}

	for i, data := range datas {
		if actual := FindKthLargest(data.nums, data.k); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
