package main

import (
	"fmt"
	"time"

	"drofff.com/revsynth/aco"
	"drofff.com/revsynth/circuit"
	"drofff.com/revsynth/cli"
	"drofff.com/revsynth/logging"
)

func main() {
	conf := aco.Config{
		NumOfAnts:       30,
		NumOfIterations: 15,
		Alpha:           2.0,
		Beta:            1.0,
		EvaporationRate: 0.5,
		DepositStrength: 100,

		LocalLoops:  4,
		SearchDepth: 6,

		Logger: logging.NewLogger(logging.LevelInfo),
	}
	synth := aco.NewSynth(conf)

	fmt.Println("Running synthesis..")
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

	fmt.Println("==========================")
	fmt.Printf("Processing time: %v millis\n", processingTime)
	fmt.Printf("Result:\n  Complexity=%v\n  NumOfGates=%v\n", res.Complexity, len(res.Gates))
	fmt.Print("\n\n")

	cli.DrawCircuit(len(desiredVector.Inputs[0]), res.Gates)
}
