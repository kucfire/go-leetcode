package leetcode014

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	datas := []struct {
		strs   []string
		result string
	}{
		{[]string{"flower", "flow", "flight"}, "fl"},
		{[]string{"dog", "racecar", "car"}, ""},
	}

	for i, data := range datas {
		if actual := LongestCommonPrefix(data.strs); actual != data.result {
			t.Errorf("example %d actually got %s, but expect %s", i, actual, data.result)
		}
	}
}
