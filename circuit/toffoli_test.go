package circuit

import (
	"testing"
)

func TestToffoliGate_Apply(t *testing.T) {
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

	tt1 := toffoliGate{controlBits: []int{2, 2, 0}, targetBit: 0}.Apply(tt)
	tt2 := toffoliGate{controlBits: []int{2, 2, 2}, targetBit: 1}.Apply(tt1)
	tt3 := toffoliGate{controlBits: []int{0, 0, 2}, targetBit: 2}.Apply(tt2)

	for i := range expectedResult {
		if !equal(expectedResult[i].Input, tt3.Rows[i].Input) {
			t.Errorf("input: unexpected row %v, e=%v, a=%v", i, expectedResult[i].Input, tt3.Rows[i].Input)
		}

		if !equal(expectedResult[i].Output, tt3.Rows[i].Output) {
			t.Errorf("output: expected row %v, e=%v, a=%v", i, expectedResult[i].Output, tt3.Rows[i].Output)
		}
	}
}
