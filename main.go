package main

import (
	"fmt"
	"time"

	"drofff.com/revsynth/aco"
	"drofff.com/revsynth/circuit"
	"drofff.com/revsynth/logging"
)

func main() {
	conf := aco.Config{
		NumOfAnts:       20,
		NumOfIterations: 10,
		Alpha:           2.0,
		Beta:            1.0,
		EvaporationRate: 0.5,
		DepositStrength: 100,

		LocalLoops:  10,
		SearchDepth: 15,

		Logger: logging.NewLogger(logging.LevelDebug),
	}
	synth := aco.NewSynth(conf)

	startedAt := time.Now().UnixMilli()
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

	processingTime := time.Now().UnixMilli() - startedAt
	fmt.Printf("Result complexity = %v, processing time %v millis\n", res.Complexity, processingTime)
}
