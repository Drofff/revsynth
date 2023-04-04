package aco

import (
	"log"

	"drofff.com/revsynth/circuit"
)

func CalcComplexity(actualState, desiredState circuit.TruthTable) int {
	actualVector := actualState.ToVector().Vector
	desiredVector := desiredState.ToVector().Vector

	if len(actualVector) != len(desiredVector) {
		log.Fatalf("critical error, different table dimension: actual %v desired %v\n", len(actualVector), len(desiredVector))
	}

	dist := 0
	for i := 0; i < len(desiredVector); i++ {
		if actualVector[i] != desiredVector[i] {
			dist++
		}
	}
	return dist
}
