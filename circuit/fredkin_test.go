package circuit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFredkinGate_Apply(t *testing.T) {
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

	expectedResult := TruthTable{
		Rows: []TruthTableRow{
			{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
			{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
			{Input: []int{0, 1, 0}, Output: []int{1, 0, 0}},
			{Input: []int{0, 1, 1}, Output: []int{1, 0, 1}},
			{Input: []int{1, 0, 0}, Output: []int{0, 1, 0}},
			{Input: []int{1, 0, 1}, Output: []int{0, 1, 1}},
			{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
			{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
		},
	}

	tt1 := fredkinGate{targetBits: []int{1, 2}, controlBits: []int{0, 2, 2}}.Apply(tt)
	tt2 := fredkinGate{targetBits: []int{0, 1}, controlBits: []int{2, 2, 2}}.Apply(tt1)
	tt3 := fredkinGate{targetBits: []int{0, 2}, controlBits: []int{2, 0, 2}}.Apply(tt2)

	assert.Equal(t, expectedResult.ToVector().Vector, tt3.ToVector().Vector)
}
