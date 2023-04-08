package aco

import "drofff.com/revsynth/logging"

type Config struct {
	NumOfAnts       int
	NumOfIterations int
	// Alpha controls the relative importance of pheromone trails.
	Alpha float64
	// Beta controls the relative importance of distance (visibility).
	Beta float64
	// EvaporationRate controls how quickly the pheromone trails evaporate over time.
	EvaporationRate float64
	// DepositStrength the strength of the pheromone deposit.
	DepositStrength float64

	// LocalLoops defines number of leaps an ant makes attempting to get to the food (zero-state).
	LocalLoops int
	// SearchDepth defines number of gates an ant can use within a leap (local loop).
	SearchDepth int

	Logger logging.Logger
}
