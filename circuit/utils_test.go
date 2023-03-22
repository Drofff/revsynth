package circuit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCutLoops(t *testing.T) {
	tests := []struct {
		title      string
		input      [][]int
		wantOutput [][]int
	}{
		{title: "no cut", input: [][]int{{2, 2}, {1, 3}, {4, 6}, {5, 5}, {8, 1}, {3, 3}}, wantOutput: [][]int{{2, 2}, {1, 3}, {4, 6}, {5, 5}, {8, 1}, {3, 3}}},
		{title: "cut", input: [][]int{{2, 2}, {1, 3}, {4, 6}, {1, 0}, {8, 1}, {4, 6}, {3, 3}, {5, 5}}, wantOutput: [][]int{{2, 2}, {1, 3}, {4, 6}, {3, 3}, {5, 5}}},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			output := CutLoops(test.input)
			assert.Equal(t, test.wantOutput, output)
		})
	}
}

func TestTrim(t *testing.T) {
	tests := []struct {
		title      string
		input      [][]int
		wantOutput [][]int
	}{
		{title: "no trim", input: [][]int{{2, 2}, {1, 3}, {4, 6}, {5, 5}}, wantOutput: [][]int{{2, 2}, {1, 3}, {4, 6}, {5, 5}}},
		{title: "trim", input: [][]int{{2, 2}, {1, 3}, {4, 6}, {1, 3}}, wantOutput: [][]int{{2, 2}, {1, 3}}},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			output := Trim(test.input)
			assert.Equal(t, test.wantOutput, output)
		})
	}
}
