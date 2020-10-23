package leetcode122

import "testing"

func TestMaxProfit(t *testing.T) {
	datas := []struct {
		prices []int
		result int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 7},
		{[]int{1, 2, 3, 4, 5}, 4},
		{[]int{7, 6, 4, 3, 1}, 0},
	}

	for _, data := range datas {
		if actual := MaxProfit(data.prices); actual != data.result {
			t.Errorf("actually got %d; but expect %d", actual, data.result)
		}
	}
}
