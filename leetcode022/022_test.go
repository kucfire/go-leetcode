package leetcode022

import (
	"reflect"
	"testing"
)

func TestGenerateParenthesis(t *testing.T) {
	datas := []struct {
		n      int
		result []string
	}{
		{3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
	}

	for i, data := range datas {
		if actual := GenerateParenthesis(data.n); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestGenerateParenthesis2(t *testing.T) {
	datas := []struct {
		n      int
		result []string
	}{
		{3, []string{
			"((()))",
			"(()())",
			"(())()",
			"()(())",
			"()()()",
		}},
	}

	for i, data := range datas {
		if actual := GenerateParenthesis2(data.n); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
