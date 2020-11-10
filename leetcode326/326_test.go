package leetcode326

import "testing"

func TestIsPowerOfThree(t *testing.T) {
	datas := []struct {
		n      int
		result bool
	}{
		{27, true},
		{0, false},
		{9, true},
		{45, false},
		{1, true},
		{3, true},
		{-3, false},
		{2, false},
	}

	for i, data := range datas {
		if actual := IsPowerOfThree(data.n); actual != data.result {
			t.Errorf("example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
