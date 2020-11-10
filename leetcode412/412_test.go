package leetcode412

import (
	"reflect"
	"testing"
)

func TestfizzBuzz(t *testing.T) {
	datas := []struct {
		n      int
		result []string
	}{
		{15, []string{"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz"}},
	}

	for i, data := range datas {
		if actual := fizzBuzz(data.n); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("example %d actuall got %v, but expect %v", i+1, actual, data.result)
		}
	}
}

func TestfizzBuzz2(t *testing.T) {
	datas := []struct {
		n      int
		result []string
	}{
		{15, []string{"1",
			"2",
			"Fizz",
			"4",
			"Buzz",
			"Fizz",
			"7",
			"8",
			"Fizz",
			"Buzz",
			"11",
			"Fizz",
			"13",
			"14",
			"FizzBuzz"}},
	}

	for i, data := range datas {
		if actual := fizzBuzz2(data.n); !reflect.DeepEqual(actual, data.result) {
			t.Errorf("example %d actuall got %v, but expect %v", i+1, actual, data.result)
		}
	}
}
