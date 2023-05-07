package main

import (
	"fmt"
	"time"

	"github.com/Drofff/revsynth/aco"
	"github.com/Drofff/revsynth/circuit"
	"github.com/Drofff/revsynth/cli"
	"github.com/Drofff/revsynth/logging"
)

func main() {
	conf := aco.Config{
		NumOfAnts:       100,
		NumOfIterations: 30,
		Alpha:           2.0,
		Beta:            1.5,
		EvaporationRate: 0.4,
		DepositStrength: 100,

		LocalLoops:  20,
		SearchDepth: 10,
	}
	synth := aco.NewSynthesizer(conf, []circuit.GateFactory{circuit.NewToffoliGateFactory()}, logging.NewLogger(logging.LevelInfo))

	fmt.Println("Running synthesis..")
	startedAt := time.Now().UnixMilli()
	desiredVector := circuit.TruthVector{Inputs: [][]int{
		{0, 0, 0, 1},
		{0, 1, 0, 1},
		{1, 0, 0, 1},
		{1, 1, 0, 1},
	}, Vector: []int{0, 2, 2, 1}, AdditionalLinesMask: []int{1, 1, 0, 0}}
	res := synth.Synthesise(desiredVector)

	processingTime := time.Now().UnixMilli() - startedAt

	fmt.Println("==========================")
	fmt.Printf("Processing time: %v millis\n", processingTime)

	if len(res.Gates) == 0 {
		fmt.Println("Failed to find a solution. Please try different synthesis parameters")
		return
	}

	fmt.Printf("Result:\n  Complexity=%v\n  NumOfGates=%v\n", res.Complexity, len(res.Gates))
	fmt.Print("\n\n")

	cli.DrawCircuit(len(desiredVector.Inputs[0]), res.Gates)
}
