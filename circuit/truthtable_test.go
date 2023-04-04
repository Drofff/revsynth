package circuit

import (
	"testing"
)

func TestTruthVector_ToTable(t *testing.T) {
	expectedTable := TruthTable{Rows: []TruthTableRow{
		{Input: []int{0, 0, 0}, Output: []int{0, 1, 0}},
		{Input: []int{0, 0, 1}, Output: []int{1, 1, 0}},
		{Input: []int{0, 1, 0}, Output: []int{0, 0, 0}},
		{Input: []int{0, 1, 1}, Output: []int{1, 0, 1}},
		{Input: []int{1, 0, 0}, Output: []int{1, 1, 1}},
		{Input: []int{1, 0, 1}, Output: []int{0, 1, 1}},
		{Input: []int{1, 1, 0}, Output: []int{1, 0, 0}},
		{Input: []int{1, 1, 1}, Output: []int{0, 0, 1}},
	}}

	v := TruthVector{
		Inputs: [][]int{
			{0, 0, 0},
			{0, 0, 1},
			{0, 1, 0},
			{0, 1, 1},
			{1, 0, 0},
			{1, 0, 1},
			{1, 1, 0},
			{1, 1, 1},
		},
		Vector: []int{2, 6, 0, 5, 7, 3, 4, 1},
	}

	tt := v.ToTable()

	for i := range tt.Rows {
		if !equal(expectedTable.Rows[i].Input, tt.Rows[i].Input) {
			t.Errorf("i=%v, e=%v, a=%v", i, expectedTable.Rows[i].Input, tt.Rows[i].Input)
		}
		if !equal(expectedTable.Rows[i].Output, tt.Rows[i].Output) {
			t.Errorf("i=%v, e=%v, a=%v", i, expectedTable.Rows[i].Output, tt.Rows[i].Output)
		}
	}
}

func TestTruthTable_ToVector(t *testing.T) {
	tt := TruthTable{Rows: []TruthTableRow{
		{Input: []int{0, 0, 0}, Output: []int{0, 1, 0}},
		{Input: []int{0, 0, 1}, Output: []int{1, 1, 0}},
		{Input: []int{0, 1, 0}, Output: []int{0, 0, 0}},
		{Input: []int{0, 1, 1}, Output: []int{1, 0, 1}},
		{Input: []int{1, 0, 0}, Output: []int{1, 1, 1}},
		{Input: []int{1, 0, 1}, Output: []int{0, 1, 1}},
		{Input: []int{1, 1, 0}, Output: []int{1, 0, 0}},
		{Input: []int{1, 1, 1}, Output: []int{0, 0, 1}},
	}}

	v := tt.ToVector()
	if !equal([]int{2, 6, 0, 5, 7, 3, 4, 1}, v.Vector) {
		t.Errorf("invalid vector %v", v.Vector)
	}

	if !equal([]int{1, 0, 0}, v.Inputs[4]) {
		t.Errorf("invalid input %v", v.Inputs[4])
	}

	key := tt.Key()
	if "[2 6 0 5 7 3 4 1]" != key {
		t.Errorf("invalid key: %v", key)
	}
}
