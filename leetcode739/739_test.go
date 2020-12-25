package leetcode739

import (
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	datas := []struct {
		T      []int
		result []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}

	for i, data := range datas {
		if actual := DailyTemperatures(data.T); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestDailyTemperatures2(t *testing.T) {
	datas := []struct {
		T      []int
		result []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}

	for i, data := range datas {
		if actual := DailyTemperatures2(data.T); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestDailyTemperatures3(t *testing.T) {
	datas := []struct {
		T      []int
		result []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
	}

	for i, data := range datas {
		if actual := DailyTemperatures3(data.T); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("Example %d actually got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
