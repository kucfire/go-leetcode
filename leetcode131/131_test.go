package leetcode131

import (
	"reflect"
	"testing"
)

func TestPartitionReview(t *testing.T) {
	datas := []struct {
		s      string
		result [][]string
	}{
		{"aab", [][]string{
			{"aa", "b"},
			{"a", "a", "b"},
		}},
	}

	for i, data := range datas {
		if actual := PartitionReview(data.s); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestPartition(t *testing.T) {
	datas := []struct {
		s      string
		result [][]string
	}{
		{"aab", [][]string{
			{"aa", "b"},
			{"a", "a", "b"},
		}},
	}

	for i, data := range datas {
		if actual := Partition(data.s); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
