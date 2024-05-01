package main

import (
	"fmt"
	"time"

	"github.com/Drofff/revsynth/aco"
	"github.com/Drofff/revsynth/circuit"
	"github.com/Drofff/revsynth/cli"
	"github.com/Drofff/revsynth/logging"
)

const pheromonesOutputFile = "pheromones.gv"

func main() {
	conf := aco.Config{
		NumOfAnts:       30,
		NumOfIterations: 15,
		Alpha:           1,
		Beta:            -1,
		EvaporationRate: 0.4,
		DepositStrength: 100,

		LocalLoops:  15,
		SearchDepth: 4,

		AllowedControlBitValues: []int{circuit.ControlBitPositive, circuit.ControlBitIgnore},
	}
	synth := aco.NewSynthesizer(conf,
		[]circuit.GateFactory{circuit.NewToffoliGateFactory()},
		logging.NewLogger(logging.LevelInfo))

	desiredVector := circuit.TruthVector{Inputs: [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 1},
		{0, 0, 1, 0},
		{0, 0, 1, 1},
		{0, 1, 0, 0},
		{0, 1, 0, 1},
		{0, 1, 1, 0},
		{0, 1, 1, 1},
	}, Vector: []int{0, 2, 6, 1, 14, 5, 13, 3}, AdditionalLinesMask: []int{}}

	var res aco.SynthesisResult
	synthIteration := 0
	for {
		fmt.Printf("Running synthesis.. %v iteration\n", synthIteration)
		synthIteration++

		startedAt := time.Now().UnixMilli()

		res = synth.Synthesise(desiredVector)

		processingTime := time.Now().UnixMilli() - startedAt
		fmt.Println("==========================")
		fmt.Printf("Processing time: %v millis\n", processingTime)

		if res.Complexity == 0 {
			break
		}
		fmt.Printf("continue synthesizing, got complexity %v\n", res.Complexity)
	}

	fmt.Printf("Saving pheromoes graph in %v\n", pheromonesOutputFile)
	err := cli.SaveGraph(res.Pheromones, pheromonesOutputFile)
	if err != nil {
		fmt.Printf("ERROR: failed to save pheromones %e", err)
	}

	if len(res.Gates) == 0 {
		fmt.Println("Failed to find a solution. Please try different synthesis parameters")
		return
	}

	fmt.Printf("Result:\n  Complexity=%v\n  NumOfGates=%v\n", res.Complexity, len(res.Gates))
	fmt.Print("\n\n")

	cli.DrawCircuit(len(desiredVector.Inputs[0]), res.Gates)
}
