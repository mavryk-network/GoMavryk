package parse

import (
	"reflect"
	"testing"
)

func TestPathsRelativeToStruct(t *testing.T) {
	tests := []struct {
		name  string
		paths [][]int
		want  [][]int
	}{
		{
			name:  "empty",
			paths: [][]int{},
			want:  [][]int{},
		},
		{
			name:  "single path - becomes relative [0]",
			paths: [][]int{{0, 1, 2}},
			want:  [][]int{{0}},
		},
		{
			name:  "no common prefix",
			paths: [][]int{{0}, {1}},
			want:  [][]int{{0}, {1}},
		},
		{
			name:  "common prefix [0]",
			paths: [][]int{{0, 0}, {0, 1, 0}, {0, 1, 1}},
			want:  [][]int{{0}, {1, 0}, {1, 1}},
		},
		{
			name:  "common prefix [0, 1]",
			paths: [][]int{{0, 1, 0}, {0, 1, 1, 0}, {0, 1, 1, 1}},
			want:  [][]int{{0}, {1, 0}, {1, 1}},
		},
		{
			name:  "all paths identical",
			paths: [][]int{{0, 1}, {0, 1}, {0, 1}},
			want:  [][]int{{0}, {0}, {0}},
		},
		{
			name:  "prefix is entire first path",
			paths: [][]int{{0}, {0, 1}, {0, 2}},
			want:  [][]int{{0}, {1}, {2}},
		},
		{
			name:  "pair fields",
			paths: [][]int{{0}, {1}},
			want:  [][]int{{0}, {1}},
		},
		{
			name:  "triple fields (already relative)",
			paths: [][]int{{0}, {1, 0}, {1, 1}},
			want:  [][]int{{0}, {1, 0}, {1, 1}},
		},
		{
			name:  "deeply nested common prefix",
			paths: [][]int{{0, 1, 2, 3, 0}, {0, 1, 2, 3, 1}},
			want:  [][]int{{0}, {1}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := pathsRelativeToStruct(tt.paths)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathsRelativeToStruct(%v) = %v, want %v",
					tt.paths, got, tt.want)
			}
		})
	}
}
