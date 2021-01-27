package leetcode088

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	datas := []struct {
		nums1, nums2 []int
		m, n         int
		result       []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, []int{}, 1, 0, []int{1}},
	}

	for i, data := range datas {
		if actual := Merge(data.nums1, data.m, data.nums2, data.n); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestMergeReviwe(t *testing.T) {
	datas := []struct {
		nums1, nums2 []int
		m, n         int
		result       []int
	}{
		{[]int{1, 2, 3, 0, 0, 0}, []int{2, 5, 6}, 3, 3, []int{1, 2, 2, 3, 5, 6}},
		{[]int{1}, []int{}, 1, 0, []int{1}},
	}

	for i, data := range datas {
		if actual := MergeReview(data.nums1, data.m, data.nums2, data.n); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
