package leetcode005

import "testing"

func TestlongestPalindrome(t *testing.T) {
	datas := []struct {
		s      string
		result string
	}{
		{"babad", "bab"},
		{"cbbd", "bb"},
	}

	for _, data := range datas {
		if actual := longestPalindrome(data.s); actual != data.result {
			t.Errorf("actually got %s, not get expect %s", actual, data.result)
		}

	}
}
