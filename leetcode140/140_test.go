package leetcode140

import (
	"reflect"
	"testing"
)

func TestWordBreak(t *testing.T) {
	datas := []struct {
		s        string
		wordDict []string
		result   []string
	}{
		{"catsanddog",
			[]string{"cat", "cats", "and", "sand", "dog"},
			[]string{"cats and dog", "cat sand dog"}},
		{"pineapplepenapple",
			[]string{"apple", "pen", "applepen", "pine", "pineapple"},
			[]string{"pine apple pen apple", "pineapple pen apple", "pine applepen apple"}},
		{"catsandog",
			[]string{"cats", "dog", "sand", "and", "cat"},
			[]string{}},
	}

	for i, data := range datas {
		if actual := WordBreak(data.s, data.wordDict); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
