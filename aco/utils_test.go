package aco

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChooseRand(t *testing.T) {
	probs := []float64{0.6, 0.3, 0.1}
	res := chooseRand(probs)
	assert.True(t, res >= 0)
	t.Logf("rand: %v\n", res)
}

func TestHaveSameElements(t *testing.T) {
	res := haveSameElements([]int{2, 3}, []int{1, 2, 5})
	assert.False(t, res)

	res = haveSameElements([]int{2, 3, 5}, []int{5, 3, 1})
	assert.False(t, res)

	res = haveSameElements([]int{5, 2, 9}, []int{2, 9, 5})
	assert.True(t, res)
}
