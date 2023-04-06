package aco

import (
	"github.com/stretchr/testify/assert"
	"log"
	"testing"
)

func TestChooseRand(t *testing.T) {
	probs := []float64{0.6, 0.3, 0.1}
	res := chooseRand(probs)
	assert.True(t, res >= 0)
	log.Printf("random value %v\n", res)
}
