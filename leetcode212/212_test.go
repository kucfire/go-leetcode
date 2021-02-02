package leetcode212

import (
	"reflect"
	"testing"
)

func TestFindWords(t *testing.T) {
	datas := []struct {
		board  [][]byte
		words  []string
		result []string
	}{
		{[][]byte{
			{'o', 'a', 'a', 'n'},
			{'e', 't', 'a', 'e'},
			{'i', 'h', 'k', 'r'},
			{'i', 'f', 'l', 'v'}},
			[]string{"oath", "pea", "eat", "rain"},
			[]string{"oath", "eat"},
		},
		{
			[][]byte{
				{'a', 'b'},
				{'c', 'd'},
			},
			[]string{"abcb"},
			[]string{},
		},
		{
			[][]byte{{'a', 'b'}},
			[]string{"ba"},
			[]string{"ba"},
		},
		{
			[][]byte{
				{'a', 'b', 'c'},
				{'a', 'e', 'd'},
				{'a', 'f', 'g'},
			},
			[]string{"abcdefg", "gfedcbaaa", "eaabcdgfa", "befa", "dgc", "ade"},
			[]string{"abcdefg", "befa", "eaabcdgfa", "gfedcbaaa"},
		},
		{
			[][]byte{
				{'a', 'a'},
			},
			[]string{"aaa"},
			[]string{},
		},
	}

	for i, data := range datas {
		if actual := FindWords(data.board, data.words); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %v actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
