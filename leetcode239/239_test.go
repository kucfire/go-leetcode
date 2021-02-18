package leetcode239

import (
	"reflect"
	"testing"
)

func TestMaxSlidingWindow(t *testing.T) {
	datas := []struct {
		nums   []int
		k      int
		result []int
	}{
		{[]int{1, 3, -1, -3, 5, 3, 6, 7}, 3, []int{3, 3, 5, 5, 6, 7}},
		{[]int{1}, 1, []int{1}},
		{[]int{1, -1}, 1, []int{1, -1}},
		{[]int{9, 11}, 2, []int{11}},
		{[]int{4, -2}, 2, []int{4}},
	}

	for i, data := range datas {
		if actual := MaxSlidingWindow(data.nums, data.k); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
