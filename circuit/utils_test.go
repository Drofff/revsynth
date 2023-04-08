package circuit

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
		title         string
		states        []TruthTable
		gates         []ToffoliGate
		wantOutStates []TruthTable
		wantOutGates  []ToffoliGate
	}{
		{
			title: "no trim",
			states: []TruthTable{
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 1, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 0, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{1, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{0, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
			},
			gates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
				{ControlBits: []int{1, 1, 0}, TargetBit: 3},
			},
			wantOutStates: []TruthTable{
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 1, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 0, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{1, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{0, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
			},
			wantOutGates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
				{ControlBits: []int{1, 1, 0}, TargetBit: 3},
			},
		},

		{
			title: "trim",
			states: []TruthTable{
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 1, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 0, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{1, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{0, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 1, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 0, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{1, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{0, 1, 1}},
					},
				},
			},
			gates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
				{ControlBits: []int{1, 1, 0}, TargetBit: 3},
				{ControlBits: []int{1, 0, 0}, TargetBit: 1},
				{ControlBits: []int{2, 1, 0}, TargetBit: 4},
			},
			wantOutStates: []TruthTable{
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 1, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{0, 0, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
					},
				},
				{
					Rows: []TruthTableRow{
						{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
						{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
						{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
						{Input: []int{0, 1, 1}, Output: []int{1, 1, 1}},
						{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
						{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
						{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
						{Input: []int{1, 1, 1}, Output: []int{0, 1, 1}},
					},
				},
			},
			wantOutGates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			outStates, outGates := Trim(test.states, test.gates)
			assert.Equal(t, test.wantOutStates, outStates)
			assert.Equal(t, test.wantOutGates, outGates)
		})
	}
}
