package leetcode171

import "testing"

func TestTitleToNumber(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"A", 1},
		{"AB", 28},
		{"ZY", 701},
	}

	for i, data := range datas {
		if actual := TitleToNumber(data.s); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
