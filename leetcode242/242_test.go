package leetcode242

import "testing"

func TestIsAnagram(t *testing.T) {
	datas := []struct {
		s, t   string
		result bool
	}{
		{"anagram", "nagaram", true},
		{"rat", "car", false},
	}

	for _, data := range datas {
		if actual := IsAnagram(data.s, data.t); actual != data.result {
			t.Errorf("actually got %v, but expect %v", actual, data.result)
		}
	}
}
