package leetcode204

import "testing"

func TestcountPrimes(t *testing.T) {
	datas := []struct {
		n      int
		result int
	}{
		{10, 4},
		{0, 0},
		{1, 0},
	}

	for i, data := range datas {
		if actual := countPrimes(data.n); actual != data.result {
			t.Errorf("example %d actual got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
