package leetcode887

import "testing"

func TestSuperEggDrop2(t *testing.T) {
	datas := []struct {
		K      int
		N      int
		result int
	}{
		{1, 2, 2},
		{2, 6, 3},
		{3, 14, 4},
		{6, 10000, 16},
	}

	for _, data := range datas {
		if actual := SuperEggDrop2(data.K, data.N); actual != data.result {
			t.Errorf("got %v; but expect %v", actual, data.result)
		}
	}
}

func TestSuperEggDrop(t *testing.T) {
	datas := []struct {
		K      int
		N      int
		result int
	}{
		{1, 2, 2},
		{2, 6, 3},
		{3, 14, 4},
		{6, 10000, 16},
	}

	for _, data := range datas {
		if actual := SuperEggDrop(data.K, data.N); actual != data.result {
			t.Errorf("got %v; but expect %v", actual, data.result)
		}
	}
}

func TestSuperEggDropReview(t *testing.T) {
	datas := []struct {
		K      int
		N      int
		result int
	}{
		{1, 2, 2},
		{2, 6, 3},
		{3, 14, 4},
		{6, 10000, 16},
	}

	for _, data := range datas {
		if actual := SuperEggDropReview(data.K, data.N); actual != data.result {
			t.Errorf("got %v; but expect %v", actual, data.result)
		}
	}
}
