package aco

import "drofff.com/revsynth/circuit"

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

	LocalLoops  int
	SearchDepth int
}

type AntTour struct {
	// Path is an ordered list of nodes ant went through.
	Path []int
	// Length defines distance (cost) of the Path.
	Length float64
}

type Synth struct {
	conf Config
}

// Pheromones key is FromState and ToState combined into a string
type Pheromones map[string]PheromoneDeposit

type PheromoneDeposit struct {
	FromState       circuit.TruthVector
	ToState         circuit.TruthVector
	PheromoneAmount int
}

type SynthesisResult struct {
	TruthTable circuit.TruthTable
	Gates      []circuit.ToffoliGate
	Complexity int
}

func NewSynth(conf Config) *Synth {
	return &Synth{conf: conf}
}

func (s *Synth) selectGate(tt circuit.TruthTable, pheromones Pheromones) circuit.ToffoliGate {

	// check for each possible truth bit what pheromone + cost it is

	// then same for control

	// then build a gate using the selected bits above

	return circuit.ToffoliGate{}
}

func (s *Synth) updatePheromones(pheromones Pheromones) Pheromones {
	// чим ближче ми добрались в турі до бажаного результату і чим менше при тому використали гейтів тим більше феромонів залишаємо
}

// Synthesise uses desiredVector as a starting point and "zero-state" as the target state.
func (s *Synth) Synthesise(desiredVector circuit.TruthVector) SynthesisResult {

	targetState := circuit.InitZeroTruthTable(desiredVector.Inputs)
	pheromones := Pheromones{}

	bestTruthTable := desiredVector.ToTable()
	bestGates := make([]circuit.ToffoliGate, 0)
	bestDist := CalcComplexity(bestTruthTable, targetState)

	for iteration := 0; iteration < s.conf.NumOfIterations; iteration++ {

		for ant := 0; ant < s.conf.NumOfAnts; ant++ {

			tourTruthTable := desiredVector.ToTable().Copy()
			tourGates := make([]circuit.ToffoliGate, 0)
			tourDist := CalcComplexity(tourTruthTable, targetState)

			for localLoop := 0; localLoop < s.conf.LocalLoops; localLoop++ {

				if tourDist == 0 {
					// ant has arrived to the final state
					break
				}

				nextGate := s.selectGate(tourTruthTable, pheromones)

				var localGates []circuit.ToffoliGate
				copy(localGates, tourGates)
				localGates = append(localGates, nextGate)

				localTruthTable := tourTruthTable.Copy()
				localTruthTable = circuit.UpdateTruthTable(localTruthTable, nextGate)

				localDist := CalcComplexity(localTruthTable, targetState)

				for depth := 0; depth < s.conf.SearchDepth; depth++ {

					nextGate := s.selectGate(localTruthTable, pheromones)

					localGates = append(localGates, nextGate)
					localTruthTable = circuit.UpdateTruthTable(localTruthTable, nextGate)
					localDist = CalcComplexity(localTruthTable, targetState)

					if localDist < tourDist {
						tourTruthTable = localTruthTable
						tourGates = localGates
						tourDist = localDist
					}

				}

			}

			if (tourDist < bestDist) || (tourDist == bestDist && len(tourGates) < len(bestGates)) {
				bestTruthTable = tourTruthTable
				bestGates = tourGates
				bestDist = tourDist
			}

		}

		pheromones = s.updatePheromones(pheromones)

	}

	return SynthesisResult{TruthTable: bestTruthTable, Gates: bestGates, Complexity: bestDist}
}
