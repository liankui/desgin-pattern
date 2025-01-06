package cracking

import (
	"reflect"
	"testing"
)

// TODO: 未完成
func cutSquares(square1 []int, square2 []int) []float64 {
	s1 := getCenter(square1)
	s2 := getCenter(square2)

	// 中心点 x 轴相同
	if s1[0] == s2[0] {
		minY, maxY := float64(square1[1]), float64(square2[1])
		if minY > maxY {
			minY, maxY = maxY, minY
		}
		return []float64{s1[0], minY, s1[0], maxY}
	}

	// 中心点 y 轴相同
	if s1[1] == s2[1] {
		minX, maxX := float64(square1[0]), float64(square2[0])
		if minX > maxX {
			minX, maxX = maxX, minX
		}
		return []float64{minX, s1[1], maxX, s1[1]}
	}

	// 相交于平行于y轴的边上
	// y - y1 = slope * (x - x1)
	slope := (s2[1] - s1[1]) / (s2[0] - s1[0])
	tmpY1 := slope*(float64(square1[0])-s1[0]) + s1[1]
	if (tmpY1 < s1[1] && tmpY1 < s2[1]) || (tmpY1 > s1[1] && tmpY1 > s2[1]) {
		// 相交于平行于x轴的边上
		x1 := (float64(square1[1])-s1[1])/slope + s1[0]
		x2 := (float64(square2[1])-s1[1])/slope + s1[0]
		if slope > 0 {
			return []float64{x1, float64(square1[1]), x2, float64(square2[1])}
		} else {
			return []float64{x2, float64(square2[1]), x1, float64(square1[1])}
		}
	}
	tmpY2 := slope*(float64(square2[0])-s1[0]) + s1[1]
	if slope > 0 {
		return []float64{float64(square1[0]), tmpY1, float64(square2[0]), tmpY2}
	} else {
		return []float64{float64(square2[0]), tmpY2, float64(square1[0]), tmpY1}
	}
}

func getCenter(square1 []int) []float64 {
	radius := float64(len(square1)) / 2
	return []float64{float64(square1[0]) + radius, float64(square1[1]) + radius}
}

func Test_cutSquares(t *testing.T) {
	tests := []struct {
		square1 []int
		square2 []int
		want    []float64
	}{
		{
			[]int{-1, -1, 2},
			[]int{0, -1, 2},
			[]float64{-1.0, 0, 2.0, 0},
		},
	}

	for _, test := range tests {
		got := cutSquares(test.square1, test.square2)
		if reflect.DeepEqual(got, test.want) == false {
			t.Errorf("got: %v, want: %v", got, test.want)
		}
	}
}
