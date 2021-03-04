package leetcode279

import "testing"

func TestNumSquares(t *testing.T) {
	datas := []struct {
		n      int
		result int
	}{
		{12, 3},
		{13, 2},
	}

	for i, data := range datas {
		if actual := NumSquares(data.n); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, data.n, data.result)
		}
	}
}
