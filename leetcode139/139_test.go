package leetcode139

import "testing"

func TestWordBreak(t *testing.T) {
	datas := []struct {
		s        string
		wordDict []string
		result   bool
	}{
		{"leetcode", []string{"leet", "code"}, true},
		{"applepenapple", []string{"apple", "pen"}, true},
		{"catsandog", []string{"cats", "dog", "sand", "and", "cat"}, false},
	}

	for i, data := range datas {
		if actual := WordBreak(data.s, data.wordDict); actual != data.result {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
