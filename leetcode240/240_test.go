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
		}, 25, false},
	}

	for _, data := range datas {
		if actual := SearchMatrix(data.tasks, data.target); actual != data.result {
			t.Errorf("got %v, but exoect %v", actual, data.result)
		}
	}
}
