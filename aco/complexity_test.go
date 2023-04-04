package aco

import (
	"testing"

	"drofff.com/revsynth/circuit"
	"github.com/stretchr/testify/assert"
)

var desiredTable = circuit.TruthTable{Rows: []circuit.TruthTableRow{
	{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
	{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
	{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
	{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
	{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
	{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
	{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
	{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
}}

func TestCalcComplexity(t *testing.T) {

	hasDistTable := circuit.TruthTable{Rows: []circuit.TruthTableRow{
		{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
		{Input: []int{0, 0, 1}, Output: []int{1, 1, 1}},
		{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
		{Input: []int{0, 1, 1}, Output: []int{0, 0, 0}},
		{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
		{Input: []int{1, 0, 1}, Output: []int{0, 1, 0}},
		{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
		{Input: []int{1, 1, 1}, Output: []int{0, 1, 1}},
	}}
	res := CalcComplexity(hasDistTable, desiredTable)
	assert.Equal(t, 4, res)

	noDistTable := circuit.TruthTable{Rows: []circuit.TruthTableRow{
		{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
		{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
		{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
		{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
		{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
		{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
		{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
		{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
	}}
	res = CalcComplexity(noDistTable, desiredTable)
	assert.Equal(t, 0, res)
}
