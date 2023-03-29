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
}

type AntTour struct {
	// Path is an ordered list of nodes ant went through.
	Path []int
	// Length defines distance (cost) of the Path.
	Length float64
}

type Optimizer struct {
	conf Config
}

func NewOptimizer(conf Config) *Optimizer {
	return &Optimizer{conf: conf}
}

func (o *Optimizer) Optimize(distances [][]float64) {

}
