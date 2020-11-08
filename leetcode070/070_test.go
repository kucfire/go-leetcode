package leetcode070

import "testing"

func TestClimbStairs(t *testing.T) {
	datas := []struct {
		n      int
		result int
	}{
		{2, 2},
		{3, 3},
		{44, 1134903170},
	}

	for i, data := range datas {
		if actual := ClimbStairs(data.n); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}

func TestClimbStairs2(t *testing.T) {
	datas := []struct {
		n      int
		result int
	}{
		{2, 2},
		{3, 3},
		{44, 1134903170},
	}

	for i, data := range datas {
		if actual := ClimbStairs2(data.n); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
