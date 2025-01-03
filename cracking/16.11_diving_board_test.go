package cracking

import (
	"reflect"
	"testing"
)

func divingBoard(shorter int, longer int, k int) []int {
	if k == 0 {
		return []int{}
	}

	if shorter == longer {
		return []int{shorter * k}
	}

	res := make([]int, k+1)
	for i := 0; i <= k; i++ {
		res[i] = shorter*(k-i) + longer*i
	}

	return res
}

func TestDivingBoard(t *testing.T) {
	tests := []struct {
		shorter int
		longer  int
		k       int
		want    []int
	}{
		{1, 2, 3, []int{3, 4, 5, 6}},
		{2, 2, 3, []int{6}},
		{1, 2, 0, []int{}},
		{1, 2, 1, []int{1, 2}},
	}

	for _, test := range tests {
		got := divingBoard(test.shorter, test.longer, test.k)
		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}
