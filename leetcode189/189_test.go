package leetcode189

import (
	"reflect"
	"testing"
)

func TestRotateReview(t *testing.T) {
	datas := []struct {
		nums   []int
		k      int
		result []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		{[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
	}

	for i, data := range datas {
		if RotateReview(data.nums, data.k); !reflect.DeepEqual(data.result, data.nums) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, data.nums, data.result)
		}
	}
}
