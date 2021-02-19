package leetcode227

import "testing"

func TestCalculate(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"3+2*2", 7},
		{" 3/2 ", 1},
		{" 3+5 / 2 ", 5},
	}

	for i, data := range datas {
		if actual := Calculate(data.s); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
