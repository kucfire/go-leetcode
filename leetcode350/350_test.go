package leetcode350

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	datas := []struct {
		nums1, nums2, result []int
	}{
		{[]int{1, 2, 2, 1}, []int{2, 2}, []int{2, 2}},
		{[]int{4, 9, 5}, []int{9, 4, 9, 8, 4}, []int{4, 9}},
	}

	for _, data := range datas {
		if actual := Intersect(data.nums1, data.nums2); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("actually got %v, but expect %v", actual, data.result)
		}
	}
}
