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
