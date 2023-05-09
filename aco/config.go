package aco

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

	// AllowedControlBitValues is a list of circuit.ControlBitValues a control bit can be assigned
	// in the synthesised circuit. Defaults to all values. It is recommended to follow the default
	// unless a custom list is really needed.
	AllowedControlBitValues []int

	// UseVisibilityExplorer:
	//
	// - If set to `true`, visibility will be taken from all possible applications of the considered
	//   bits computed as the visibility area
	//
	// - If set to `false`, only known applications of the bits from the pheromones graph will be taken
	//   into consideration. This approach appears to be far more effective in the tested use-cases
	UseVisibilityExplorer bool
}
