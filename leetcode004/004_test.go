package leetcode004

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {
	datas := []struct {
		nums1, nums2 []int
		result       float64
	}{
		{[]int{1, 3}, []int{2}, 2.00000},
		{[]int{1, 2}, []int{3, 4}, 2.50000},
		{[]int{0, 0}, []int{0, 0}, 0.00000},
		{[]int{}, []int{1}, 1.00000},
		{[]int{2}, []int{}, 2.00000},
	}

	for i, data := range datas {
		if actual := FindMedianSortedArrays(data.nums1, data.nums2); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestReverseFindMedianSortedArrays(t *testing.T) {
	datas := []struct {
		nums1, nums2 []int
		result       float64
	}{
		{[]int{1, 3}, []int{2}, 2.00000},
		{[]int{1, 2}, []int{3, 4}, 2.50000},
		{[]int{0, 0}, []int{0, 0}, 0.00000},
		{[]int{}, []int{1}, 1.00000},
		{[]int{2}, []int{}, 2.00000},
		{[]int{2}, []int{1, 3, 4}, 2.50000},
	}

	for i, data := range datas {
		if actual := ReverseFindMedianSortedArrays(data.nums1, data.nums2); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
