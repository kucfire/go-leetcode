package leetcode046

import "testing"

func TestPermute(t *testing.T) {
	datas := []struct {
		data   []int
		result [][]int
	}{
		{[]int{1, 2, 3}, [][]int{
			{1, 2, 3},
			{1, 3, 2},
			{2, 1, 3},
			{2, 3, 1},
			{3, 1, 2},
			{3, 2, 1},
		}},
	}
	for _, data := range datas {
		t.Errorf("enter %v, got %v", data.data, Permute(data.data))
	}
}
