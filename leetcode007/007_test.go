package leetcode007

import "testing"

func TestReverse(t *testing.T) {
	datas := []struct {
		x      int
		result int
	}{
		{123, 321},
		{-123, -321},
		{120, 21},
		{1534236469, 0},
	}

	for _, data := range datas {
		if actual := Reverse(data.x); actual != data.result {
			t.Errorf("actually got %d, but expect %d!\n", actual, data.result)
		}
	}
}
