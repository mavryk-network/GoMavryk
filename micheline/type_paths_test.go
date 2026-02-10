package micheline

import (
	"reflect"
	"testing"
)

func TestRightCombPath(t *testing.T) {
	tests := []struct {
		name  string
		base  []int
		index int
		n     int
		want  []int
	}{
		{
			name:  "single element",
			base:  []int{},
			index: 0,
			n:     1,
			want:  []int{},
		},
		{
			name:  "pair: first element",
			base:  []int{},
			index: 0,
			n:     2,
			want:  []int{0},
		},
		{
			name:  "pair: second element",
			base:  []int{},
			index: 1,
			n:     2,
			want:  []int{1},
		},
		{
			name:  "triple: first element",
			base:  []int{},
			index: 0,
			n:     3,
			want:  []int{0},
		},
		{
			name:  "triple: second element",
			base:  []int{},
			index: 1,
			n:     3,
			want:  []int{1, 0},
		},
		{
			name:  "triple: third element",
			base:  []int{},
			index: 2,
			n:     3,
			want:  []int{1, 1},
		},
		{
			name:  "quad: elements",
			base:  []int{},
			index: 0,
			n:     4,
			want:  []int{0},
		},
		{
			name:  "quad: second",
			base:  []int{},
			index: 1,
			n:     4,
			want:  []int{1, 0},
		},
		{
			name:  "quad: third",
			base:  []int{},
			index: 2,
			n:     4,
			want:  []int{1, 1, 0},
		},
		{
			name:  "quad: fourth",
			base:  []int{},
			index: 3,
			n:     4,
			want:  []int{1, 1, 1},
		},
		{
			name:  "with base path",
			base:  []int{0, 1},
			index: 1,
			n:     3,
			want:  []int{0, 1, 1, 0},
		},
		{
			name:  "five elements: last",
			base:  []int{},
			index: 4,
			n:     5,
			want:  []int{1, 1, 1, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rightCombPath(tt.base, tt.index, tt.n)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightCombPath(%v, %d, %d) = %v, want %v",
					tt.base, tt.index, tt.n, got, tt.want)
			}
		})
	}
}
