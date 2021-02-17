package leetcode378

import "testing"

func TestKthSmallest(t *testing.T) {
	datas := []struct {
		nums   [][]int
		k      int
		result int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
		{[][]int{{-5}}, 1, -5},
	}

	for i, data := range datas {
		if actual := KthSmallest(data.nums, data.k); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestKthSmallest2(t *testing.T) {
	datas := []struct {
		nums   [][]int
		k      int
		result int
	}{
		{[][]int{{1, 5, 9}, {10, 11, 13}, {12, 13, 15}}, 8, 13},
		{[][]int{{-5}}, 1, -5},
	}

	for i, data := range datas {
		if actual := KthSmallest2(data.nums, data.k); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
