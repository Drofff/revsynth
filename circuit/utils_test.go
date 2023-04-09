package circuit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	testStateA = TruthTable{
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
	}
	testStateB = TruthTable{
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
	}
	testStateC = TruthTable{
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
	}
	testStateD = TruthTable{
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
	}
	testStateE = TruthTable{
		Rows: []TruthTableRow{
			{Input: []int{0, 0, 0}, Output: []int{0, 1, 0}},
			{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
			{Input: []int{0, 1, 0}, Output: []int{1, 1, 0}},
			{Input: []int{0, 1, 1}, Output: []int{0, 0, 0}},
			{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
			{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
			{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
			{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
		},
	}
)

func TestCutLoops(t *testing.T) {
	tests := []struct {
		title         string
		states        []TruthTable
		gates         []ToffoliGate
		wantOutStates []TruthTable
		wantOutGates  []ToffoliGate
	}{
		{
			title:  "no loops",
			states: []TruthTable{testStateA, testStateB, testStateC, testStateD},
			gates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateC, testStateD},
			wantOutGates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
			},
		},
		{
			title:  "has loops",
			states: []TruthTable{testStateA, testStateB, testStateC, testStateD, testStateB, testStateE},
			gates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
				{ControlBits: []int{2, 1, 0}, TargetBit: 0},
				{ControlBits: []int{1, 1, 2}, TargetBit: 0},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateE},
			wantOutGates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{1, 1, 2}, TargetBit: 0},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			outStates, outGates := CutLoops(test.states, test.gates)
			assert.Equal(t, test.wantOutStates, outStates)
			assert.Equal(t, test.wantOutGates, outGates)
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
			title:  "no trim",
			states: []TruthTable{testStateA, testStateB, testStateC},
			gates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateC},
			wantOutGates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
			},
		},

		{
			title:  "trim",
			states: []TruthTable{testStateA, testStateB, testStateC, testStateD, testStateE, testStateC},
			gates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
				{ControlBits: []int{2, 0, 0}, TargetBit: 2},
				{ControlBits: []int{1, 1, 0}, TargetBit: 3},
				{ControlBits: []int{0, 0, 0}, TargetBit: 2},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateC},
			wantOutGates: []ToffoliGate{
				{ControlBits: []int{2, 2, 0}, TargetBit: 0},
				{ControlBits: []int{2, 1, 0}, TargetBit: 1},
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
