package leetcode011

import "testing"

func TestMaxArea(t *testing.T) {
	datas := []struct {
		height []int
		result int
	}{
		{[]int{1, 8, 6, 2, 5, 4, 8, 3, 7}, 49},
	}

	for _, data := range datas {
		if actual := MaxArea(data.height); actual != data.result {
			t.Errorf("actually got %d; but expect %d", actual, data.result)
		}
	}
}
