package leetcode053

import "testing"

func TestMaxSubArray(t *testing.T) {
	datas := []struct {
		nums   []int
		result int
	}{
		{[]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}, 6},
	}

	for i, data := range datas {
		if actual := MaxSubArray(data.nums); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
