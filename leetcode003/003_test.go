package leetcode003

import "testing"

func TestLengthOfLongestSubstring(t *testing.T) {
	datas := []struct {
		s      string
		result int
	}{
		{"abcabcbb", 3},
		{"bbbbb", 1},
		{"pwwkew", 3},
		{"au", 2},
	}

	for _, data := range datas {
		if actual := LengthOfLongestSubstring(data.s); actual != data.result {
			t.Errorf("actually got %d; but expect get %d", actual, data.result)
		}
	}
}
