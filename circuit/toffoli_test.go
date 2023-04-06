package circuit

import (
	"testing"
)

func TestUpdateTruthTable(t *testing.T) {
	tt := TruthTable{
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

	expectedResult := []TruthTableRow{
		{Input: []int{0, 0, 0}, Output: []int{0, 1, 0}},
		{Input: []int{0, 0, 1}, Output: []int{1, 1, 0}},
		{Input: []int{0, 1, 0}, Output: []int{0, 0, 0}},
		{Input: []int{0, 1, 1}, Output: []int{1, 0, 1}},
		{Input: []int{1, 0, 0}, Output: []int{1, 1, 1}},
		{Input: []int{1, 0, 1}, Output: []int{0, 1, 1}},
		{Input: []int{1, 1, 0}, Output: []int{1, 0, 0}},
		{Input: []int{1, 1, 1}, Output: []int{0, 0, 1}},
	}

	tt1 := UpdateTruthTable(tt, ToffoliGate{ControlBits: []int{2, 2, 0}, TargetBit: 0})
	tt2 := UpdateTruthTable(tt1, ToffoliGate{ControlBits: []int{2, 2, 2}, TargetBit: 1})
	tt3 := UpdateTruthTable(tt2, ToffoliGate{ControlBits: []int{0, 0, 2}, TargetBit: 2})

	for i := range expectedResult {
		if !equal(expectedResult[i].Input, tt3.Rows[i].Input) {
			t.Errorf("input: unexpected row %v, e=%v, a=%v", i, expectedResult[i].Input, tt3.Rows[i].Input)
		}

		if !equal(expectedResult[i].Output, tt3.Rows[i].Output) {
			t.Errorf("output: expected row %v, e=%v, a=%v", i, expectedResult[i].Output, tt3.Rows[i].Output)
		}
	}
}
