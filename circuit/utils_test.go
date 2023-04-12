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
		gates         []Gate
		wantOutStates []TruthTable
		wantOutGates  []Gate
	}{
		{
			title:  "no loops",
			states: []TruthTable{testStateA, testStateB, testStateC, testStateD},
			gates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 1},
				toffoliGate{controlBits: []int{2, 0, 0}, targetBit: 2},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateC, testStateD},
			wantOutGates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 1},
				toffoliGate{controlBits: []int{2, 0, 0}, targetBit: 2},
			},
		},
		{
			title:  "has loops",
			states: []TruthTable{testStateA, testStateB, testStateC, testStateD, testStateB, testStateE},
			gates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 1},
				toffoliGate{controlBits: []int{2, 0, 0}, targetBit: 2},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{1, 1, 2}, targetBit: 0},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateE},
			wantOutGates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{1, 1, 2}, targetBit: 0},
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
		gates         []Gate
		wantOutStates []TruthTable
		wantOutGates  []Gate
	}{
		{
			title:  "no trim",
			states: []TruthTable{testStateA, testStateB, testStateC},
			gates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 1},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateC},
			wantOutGates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 1},
			},
		},

		{
			title:  "trim",
			states: []TruthTable{testStateA, testStateB, testStateC, testStateD, testStateE, testStateC},
			gates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 1},
				toffoliGate{controlBits: []int{2, 0, 0}, targetBit: 2},
				toffoliGate{controlBits: []int{1, 1, 0}, targetBit: 3},
				toffoliGate{controlBits: []int{0, 0, 0}, targetBit: 2},
			},
			wantOutStates: []TruthTable{testStateA, testStateB, testStateC},
			wantOutGates: []Gate{
				toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0},
				toffoliGate{controlBits: []int{2, 1, 0}, targetBit: 1},
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
