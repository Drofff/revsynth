package aco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcComplexity(t *testing.T) {

	desiredVector := []int{1, 2, 3, 4, 5, 6, 7, 8}

	diffLength := []int{1, 2, 9, 4, 5}
	res := CalcComplexity(diffLength, desiredVector)
	assert.Equal(t, 4, res)

	hasDist := []int{1, 9, 3, 1, 5, 10, 7, 4}
	res = CalcComplexity(hasDist, desiredVector)
	assert.Equal(t, 4, res)

	noDist := []int{1, 2, 3, 4, 5, 6, 7, 8}
	res = CalcComplexity(noDist, desiredVector)
	assert.Equal(t, 0, res)
}
