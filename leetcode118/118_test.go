package leetcode118

import (
	"reflect"
	"testing"
)

func TestGenerate(t *testing.T) {
	datas := []struct {
		numRows int
		result  [][]int
	}{
		{5, [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
	}

	for i, data := range datas {
		if actual := Generate(data.numRows); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("example %d actual got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
