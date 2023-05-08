package aco

import (
	"testing"

	"github.com/Drofff/revsynth/circuit"
	"github.com/stretchr/testify/assert"
)

func TestGenerateTargetBits(t *testing.T) {
	res := generateTargetBits(2, 4)
	assert.Equal(t, [][]int{
		{0, 1},
		{0, 2},
		{0, 3},
		{1, 0},
		{1, 2},
		{1, 3},
		{2, 0},
		{2, 1},
		{2, 3},
		{3, 0},
		{3, 1},
		{3, 2},
	}, res)
}

func TestGenerateControlBits(t *testing.T) {
	res := generateControlBits(2, 4)

	assert.Equal(t, [][]int{
		{2, 2, 2, 2},
		{0, 2, 2, 2},
		{1, 2, 2, 2},
		{2, 0, 2, 2},
		{2, 1, 2, 2},
		{2, 2, 0, 2},
		{2, 2, 1, 2},
		{2, 2, 2, 0},
		{2, 2, 2, 1},
		{0, 0, 2, 2},
		{0, 1, 2, 2},
		{1, 0, 2, 2},
		{1, 1, 2, 2},
		{0, 2, 0, 2},
		{0, 2, 1, 2},
		{1, 2, 0, 2},
		{1, 2, 1, 2},
		{0, 2, 2, 0},
		{0, 2, 2, 1},
		{1, 2, 2, 0},
		{1, 2, 2, 1},
		{0, 0, 2, 2},
		{1, 0, 2, 2},
		{0, 1, 2, 2},
		{1, 1, 2, 2},
		{2, 0, 0, 2},
		{2, 0, 1, 2},
		{2, 1, 0, 2},
		{2, 1, 1, 2},
		{2, 0, 2, 0},
		{2, 0, 2, 1},
		{2, 1, 2, 0},
		{2, 1, 2, 1},
		{0, 2, 0, 2},
		{1, 2, 0, 2},
		{0, 2, 1, 2},
		{1, 2, 1, 2},
		{2, 0, 0, 2},
		{2, 1, 0, 2},
		{2, 0, 1, 2},
		{2, 1, 1, 2},
		{2, 2, 0, 0},
		{2, 2, 0, 1},
		{2, 2, 1, 0},
		{2, 2, 1, 1},
		{0, 2, 2, 0},
		{1, 2, 2, 0},
		{0, 2, 2, 1},
		{1, 2, 2, 1},
		{2, 0, 2, 0},
		{2, 1, 2, 0},
		{2, 0, 2, 1},
		{2, 1, 2, 1},
		{2, 2, 0, 0},
		{2, 2, 1, 0},
		{2, 2, 0, 1},
		{2, 2, 1, 1},
	}, res)
}

func TestExploreVisibility(t *testing.T) {
	wantRes := []struct {
		controlBits []int
		targetBits  []int
		dist        int
	}{
		{controlBits: []int{2, 2, 2}, targetBits: []int{0}, dist: 7},
		{controlBits: []int{2, 0, 2}, targetBits: []int{0}, dist: 5},
		{controlBits: []int{2, 1, 2}, targetBits: []int{0}, dist: 6},
		{controlBits: []int{2, 2, 0}, targetBits: []int{0}, dist: 3},
		{controlBits: []int{2, 2, 1}, targetBits: []int{0}, dist: 8},
		{controlBits: []int{2, 0, 0}, targetBits: []int{0}, dist: 3},
		{controlBits: []int{2, 0, 1}, targetBits: []int{0}, dist: 6},
		{controlBits: []int{2, 1, 0}, targetBits: []int{0}, dist: 4},
		{controlBits: []int{2, 1, 1}, targetBits: []int{0}, dist: 6},
		{controlBits: []int{2, 0, 0}, targetBits: []int{0}, dist: 3},
		{controlBits: []int{2, 1, 0}, targetBits: []int{0}, dist: 4},
		{controlBits: []int{2, 0, 1}, targetBits: []int{0}, dist: 6},
		{controlBits: []int{2, 1, 1}, targetBits: []int{0}, dist: 6},
		{controlBits: []int{2, 2, 2}, targetBits: []int{1}, dist: 8},
		{controlBits: []int{0, 2, 2}, targetBits: []int{1}, dist: 6},
		{controlBits: []int{1, 2, 2}, targetBits: []int{1}, dist: 6},
		{controlBits: []int{2, 2, 0}, targetBits: []int{1}, dist: 4},
		{controlBits: []int{2, 2, 1}, targetBits: []int{1}, dist: 8},
		{controlBits: []int{0, 2, 0}, targetBits: []int{1}, dist: 4},
		{controlBits: []int{0, 2, 1}, targetBits: []int{1}, dist: 6},
		{controlBits: []int{1, 2, 0}, targetBits: []int{1}, dist: 4},
		{controlBits: []int{1, 2, 1}, targetBits: []int{1}, dist: 6},
		{controlBits: []int{0, 2, 0}, targetBits: []int{1}, dist: 4},
		{controlBits: []int{1, 2, 0}, targetBits: []int{1}, dist: 4},
		{controlBits: []int{0, 2, 1}, targetBits: []int{1}, dist: 6},
		{controlBits: []int{1, 2, 1}, targetBits: []int{1}, dist: 6},
		{controlBits: []int{2, 2, 2}, targetBits: []int{2}, dist: 8},
		{controlBits: []int{0, 2, 2}, targetBits: []int{2}, dist: 6},
		{controlBits: []int{1, 2, 2}, targetBits: []int{2}, dist: 6},
		{controlBits: []int{2, 0, 2}, targetBits: []int{2}, dist: 6},
		{controlBits: []int{2, 1, 2}, targetBits: []int{2}, dist: 6},
		{controlBits: []int{0, 0, 2}, targetBits: []int{2}, dist: 5},
		{controlBits: []int{0, 1, 2}, targetBits: []int{2}, dist: 5},
		{controlBits: []int{1, 0, 2}, targetBits: []int{2}, dist: 5},
		{controlBits: []int{1, 1, 2}, targetBits: []int{2}, dist: 5},
		{controlBits: []int{0, 0, 2}, targetBits: []int{2}, dist: 5},
		{controlBits: []int{1, 0, 2}, targetBits: []int{2}, dist: 5},
		{controlBits: []int{0, 1, 2}, targetBits: []int{2}, dist: 5},
		{controlBits: []int{1, 1, 2}, targetBits: []int{2}, dist: 5},
	}

	state := circuit.TruthTable{Rows: []circuit.TruthTableRow{
		{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
		{Input: []int{0, 0, 1}, Output: []int{0, 0, 1}},
		{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
		{Input: []int{0, 1, 1}, Output: []int{0, 1, 1}},
		{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
		{Input: []int{1, 0, 1}, Output: []int{1, 0, 1}},
		{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
		{Input: []int{1, 1, 1}, Output: []int{1, 1, 1}},
	}}
	desiredState := circuit.TruthTable{Rows: []circuit.TruthTableRow{
		{Input: []int{0, 0, 0}, Output: []int{0, 0, 0}},
		{Input: []int{0, 0, 1}, Output: []int{1, 1, 1}},
		{Input: []int{0, 1, 0}, Output: []int{0, 1, 0}},
		{Input: []int{0, 1, 1}, Output: []int{0, 0, 0}},
		{Input: []int{1, 0, 0}, Output: []int{1, 0, 0}},
		{Input: []int{1, 0, 1}, Output: []int{0, 1, 0}},
		{Input: []int{1, 1, 0}, Output: []int{1, 1, 0}},
		{Input: []int{1, 1, 1}, Output: []int{0, 1, 1}},
	}}

	res := exploreVisibility(state, desiredState, circuit.NewToffoliGateFactory())

	for i := 0; i < len(res); i++ {
		assert.Equal(t, wantRes[i].controlBits, res[i].controlBits)
		assert.Equal(t, wantRes[i].targetBits, res[i].targetBits)
		assert.Equal(t, wantRes[i].dist, res[i].distance)
	}
}
