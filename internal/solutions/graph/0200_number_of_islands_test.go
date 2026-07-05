package graph

import "testing"

func TestNumIslands(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want int
	}{
		{
			[][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			1,
		},
		{
			[][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			3,
		},
	}
	for _, tt := range tests {
		gridCopy := make([][]byte, len(tt.grid))
		for i := range tt.grid {
			gridCopy[i] = make([]byte, len(tt.grid[i]))
			copy(gridCopy[i], tt.grid[i])
		}
		got := NumIslands(gridCopy)
		if got != tt.want {
			t.Errorf("NumIslands() = %v, want %v", got, tt.want)
		}
	}
}
