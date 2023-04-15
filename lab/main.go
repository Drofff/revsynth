package main

import (
	"fmt"
	"io/fs"
	"log"
	"math"
	"os"
	"time"

	"drofff.com/revsynth/aco"
	"drofff.com/revsynth/circuit"
	"drofff.com/revsynth/logging"
)

type lab struct {
	gateFactory   circuit.GateFactory
	logger        logging.Logger
	targetCircuit circuit.TruthVector
}

type experimentSettings struct {
	filename           string
	config             aco.Config
	configModifier     func(aco.Config) aco.Config
	modificationsCount int
}

type synthesisResult struct {
	numOfGates int
	complexity int
	duration   int64
}

const iterationsPerConfig = 3

var desiredVector = circuit.TruthVector{Inputs: [][]int{
	{0, 0, 0},
	{0, 0, 1},
	{0, 1, 0},
	{0, 1, 1},
	{1, 0, 0},
	{1, 0, 1},
	{1, 1, 0},
	{1, 1, 1},
}, Vector: []int{7, 5, 2, 4, 6, 1, 0, 3}}

func summarizeResults(results []synthesisResult) synthesisResult {
	summary := synthesisResult{}
	resultsLen := float64(len(results))
	for _, result := range results {
		summary.numOfGates += int(math.Round(float64(result.numOfGates) / resultsLen))
		summary.complexity += int(math.Round(float64(result.complexity) / resultsLen))
		summary.duration += int64(math.Round(float64(result.duration) / resultsLen))
	}
	return summary
}

func (l lab) runSynthesis(config aco.Config) synthesisResult {
	s := aco.NewSynthesizer(config, l.gateFactory, l.logger)

	results := make([]synthesisResult, 0)
	for i := 0; i < iterationsPerConfig; i++ {
		startedAt := time.Now().UnixMilli()
		res := s.Synthesise(l.targetCircuit)

		duration := time.Now().UnixMilli() - startedAt
		results = append(results, synthesisResult{
			numOfGates: len(res.Gates),
			complexity: res.Complexity,
			duration:   duration,
		})
	}

	return summarizeResults(results)
}

func formatResult(config aco.Config, res synthesisResult) string {
	return fmt.Sprintf("%v,%v,%v,%v,%v,%v,%v,%v,%v,%v,%v\n",
		config.NumOfAnts,
		config.NumOfIterations,
		config.Alpha,
		config.Beta,
		config.EvaporationRate,
		config.DepositStrength,
		config.LocalLoops,
		config.SearchDepth,
		res.numOfGates,
		res.complexity,
		res.duration)
}

func main() {
	l := lab{
		gateFactory:   circuit.NewToffoliGateFactory(),
		logger:        logging.NewLogger(logging.LevelInfo),
		targetCircuit: desiredVector,
	}

	for ei, experiment := range experiments {
		fmt.Printf("starting experiment #%v (%v)\n", ei, experiment.filename)

		experimentConfig := experiment.config
		experimentResult := ""
		for mod := 0; mod <= experiment.modificationsCount; mod++ {
			fmt.Printf("modification %v\n", mod)
			modRes := l.runSynthesis(experimentConfig)
			experimentResult += formatResult(experimentConfig, modRes)

			experimentConfig = experiment.configModifier(experimentConfig)
		}

		err := os.WriteFile("./lab/"+experiment.filename+".csv", []byte(experimentResult), fs.ModePerm)
		if err != nil {
			log.Fatalln("error writing results to file:", err)
		}
	}
}
