package leetcode013

import "testing"

func TestRomanToInt(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"III", 3},
		{"IV", 4},
		{"IX", 9},
		{"LVIII", 58},
		{"MCMXCIV", 1994},
	}

	for i, data := range datas {
		if actual := RomanToInt(data.s); actual != data.result {
			t.Errorf("example %d actually got %d, but expect %d", i+1, actual, data.result)
		}
	}
}
