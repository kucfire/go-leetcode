package leetcode461

import "testing"

func TestHammingDistance(t *testing.T) {
	datas := []struct {
		x, y   int
		result int
	}{
		{1, 4, 2},
	}

	for i, data := range datas {
		if actual := HammingDistance(data.x, data.y); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
