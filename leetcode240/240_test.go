package leetcode240

import "testing"

func TestSearchMatrix(t *testing.T) {
	datas := []struct {
		tasks  [][]int
		target int
		result bool
	}{
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 5, true},
		{[][]int{{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 20, false},
	}

	for i, data := range datas {
		if actual := SearchMatrix2(data.tasks, data.target); actual != data.result {
			t.Errorf("%d, got %v, but exoect %v", i, actual, data.result)
		}
	}
}
