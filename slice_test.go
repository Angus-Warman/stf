package stf

import (
	"fmt"
	"testing"
)

func TestChunk(t *testing.T) {
	input := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}

	tests := []struct {
		numChunks int
		expected  [][]int
	}{
		{
			numChunks: 0,
			expected:  [][]int{},
		},
		{
			numChunks: 1,
			expected: [][]int{
				{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13},
			},
		},
		{
			numChunks: 2,
			expected: [][]int{
				{1, 2, 3, 4, 5, 6, 7},
				{8, 9, 10, 11, 12, 13},
			},
		},
		{
			numChunks: 3,
			expected: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9},
				{10, 11, 12, 13},
			},
		},
		{
			numChunks: 4,
			expected: [][]int{
				{1, 2, 3, 4},
				{5, 6, 7},
				{8, 9, 10},
				{11, 12, 13},
			},
		},
		{
			numChunks: 5,
			expected: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
				{10, 11},
				{12, 13},
			},
		},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("numChunks=%d", tt.numChunks), func(t *testing.T) {
			actual := Chunk(input, tt.numChunks)
			if fmt.Sprint(actual) != fmt.Sprint(tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, actual)
			}
		})
	}
}
