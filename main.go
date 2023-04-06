package main

import (
	"fmt"

	"drofff.com/revsynth/aco"
	"drofff.com/revsynth/circuit"
)

func main() {
	conf := aco.Config{
		NumOfAnts:       20,
		NumOfIterations: 100,
		Alpha:           1.5,
		Beta:            2.0,
		EvaporationRate: 0.5,
		DepositStrength: 100,

		LocalLoops:  15,
		SearchDepth: 10,
	}
	synth := aco.NewSynth(conf)

	desiredVector := circuit.TruthVector{Inputs: [][]int{
		{0, 0, 0},
		{0, 0, 1},
		{0, 1, 0},
		{0, 1, 1},
		{1, 0, 0},
		{1, 0, 1},
		{1, 1, 0},
		{1, 1, 1},
	}, Vector: []int{1, 0, 3, 2, 5, 7, 4, 6}}
	res := synth.Synthesise(desiredVector)
	fmt.Printf("Result complexity = %v\n", res.Complexity)
}
