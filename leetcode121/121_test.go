package leetcode121

import "testing"

func TestMaxProfit(t *testing.T) {
	datas := []struct {
		prices []int
		result int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
	}

	for i, data := range datas {
		if actuall := MaxProfit(data.prices); actuall != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actuall, data.result)
		}
	}
}

func TestMaxProfit2(t *testing.T) {
	datas := []struct {
		prices []int
		result int
	}{
		{[]int{7, 1, 5, 3, 6, 4}, 5},
		{[]int{7, 6, 4, 3, 1}, 0},
		{[]int{1, 2}, 1},
	}

	for i, data := range datas {
		if actuall := MaxProfit2(data.prices); actuall != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actuall, data.result)
		}
	}
}
