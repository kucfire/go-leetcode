package leetcode621

import "testing"

func TestleastInterval(t *testing.T) {
	// result
	result := 8

	data := []byte{'A', 'A', 'A', 'B', 'B', 'B'}
	coolingTime := 2
	actual := leastInterval(data, coolingTime)

	if result != actual {
		t.Errorf("result %d; but got %d", result, actual)
	}

}
