package leetcode038

import "testing"

func TestCountAndSay(t *testing.T) {
	datas := []struct {
		n      int
		result string
	}{
		{1, "1"},
		{4, "1211"},
	}

	for i, data := range datas {
		if actual := CountAndSay(data.n); actual != data.result {
			t.Errorf("example %d actual got %s, but expect %s", i+1, actual, data.result)
		}
	}
}
