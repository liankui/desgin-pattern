package cracking

import (
	"reflect"
	"testing"
)

// masterMind
// 时间复杂度：O（n）
// 空间复杂度：O（n）
func masterMind(solution string, guess string) []int {
	hit, pseudoHit := 0, 0
	solutions := make(map[uint8]int, 4)
	guesses := make(map[uint8]int, 4)
	for i := 0; i < len(solution); i++ {
		if solution[i] == guess[i] {
			hit++
		} else {
			solutions[solution[i]]++
			guesses[guess[i]]++
		}
	}

	for k, v := range solutions {
		if count, ok := guesses[k]; ok {
			if v > count {
				v = count
			}
			pseudoHit += v
		}
	}

	return []int{hit, pseudoHit}
}

func masterMind2(solution string, guess string) []int {
	hit, allHit := 0, 0
	solutions := make(map[uint8]int, 4)
	for i := 0; i < len(solution); i++ {
		solutions[solution[i]]++
	}

	for i := 0; i < 4; i++ {
		if solution[i] == guess[i] {
			hit++
		}

		if count, ok := solutions[guess[i]]; ok && count > 0 {
			allHit++
			solutions[guess[i]]--
		}
	}

	return []int{hit, allHit - hit}
}

func Test_masterMind(t *testing.T) {
	tests := []struct {
		solution string
		guess    string
		want     []int
	}{
		{
			solution: "RGBY",
			guess:    "GGRR",
			want:     []int{1, 1},
		},
		{
			solution: "RRRR",
			guess:    "RRRR",
			want:     []int{4, 0},
		},
		{
			solution: "RRRR",
			guess:    "GGRR",
			want:     []int{2, 0},
		},
		{
			solution: "RRRG",
			guess:    "GGRR",
			want:     []int{1, 2},
		},
		{
			solution: "RRRR",
			guess:    "GGGG",
			want:     []int{0, 0},
		},
	}

	for _, test := range tests {
		if got := masterMind2(test.solution, test.guess); !reflect.DeepEqual(got, test.want) {
			t.Errorf("masterMind(%s, %s) = %v, want %v", test.solution, test.guess, got, test.want)
		}
	}
}
