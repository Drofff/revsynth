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

	// LocalLoops defines number of leaps an ant makes attempting to get to the food (zero-state).
	LocalLoops int
	// SearchDepth defines number of gates an ant can use within a leap (local loop).
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

// Pheromones key is FromState and ToState combined into a string.
type Pheromones map[string]PheromoneDeposit

type PheromoneDeposit struct {
	FromState       circuit.TruthTable
	ToState         circuit.TruthTable
	PheromoneAmount float64
}

type SynthesisResult struct {
	States     []circuit.TruthTable
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

func (s *Synth) depositPheromone(pheromones Pheromones, states []circuit.TruthTable, dist int) {
	amount := s.conf.DepositStrength / float64(dist)

	for i := 0; i < len(states)-1; i++ {
		fromState := states[i]
		toState := states[i+1]

		linkKey := fromState.Key() + toState.Key()

		pheromone, exists := pheromones[linkKey]
		if !exists {
			pheromone = PheromoneDeposit{FromState: fromState, ToState: toState, PheromoneAmount: 0}
		}

		gatePosition := len(states) - (i + 1) // position from end state, increases pheromone for gates closer to start state.

		pheromone.PheromoneAmount += amount * float64(gatePosition)
		pheromones[linkKey] = pheromone
	}
}

func (s *Synth) updatePheromones(pheromones Pheromones, newDeposits Pheromones) {
	for key, deposit := range pheromones {
		deposit.PheromoneAmount *= 1.0 - s.conf.EvaporationRate
		pheromones[key] = deposit
	}

	for key, newDeposit := range newDeposits {

		deposit, exists := pheromones[key]
		if !exists {
			deposit = PheromoneDeposit{FromState: newDeposit.FromState, ToState: newDeposit.ToState, PheromoneAmount: 0}
		}

		deposit.PheromoneAmount += newDeposit.PheromoneAmount
		pheromones[key] = deposit
	}
}

// Synthesise uses desiredVector as a starting point and "zero-state" as the target state.
func (s *Synth) Synthesise(desiredVector circuit.TruthVector) SynthesisResult {

	targetState := circuit.InitZeroTruthTable(desiredVector.Inputs)
	pheromones := Pheromones{}

	bestStates := make([]circuit.TruthTable, 0)
	bestGates := make([]circuit.ToffoliGate, 0)
	bestDist := CalcComplexity(desiredVector.ToTable(), targetState)

	for iteration := 0; iteration < s.conf.NumOfIterations; iteration++ {

		iterationDeposits := Pheromones{}

		for ant := 0; ant < s.conf.NumOfAnts; ant++ {

			tourTruthTable := desiredVector.ToTable().Copy()
			tourStates := []circuit.TruthTable{tourTruthTable}

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
				var localStates []circuit.TruthTable
				copy(localStates, tourStates)
				localStates = append(localStates, localTruthTable)

				localDist := CalcComplexity(localTruthTable, targetState)

				for depth := 0; depth < s.conf.SearchDepth; depth++ {

					nextGate := s.selectGate(localTruthTable, pheromones)

					localGates = append(localGates, nextGate)
					localTruthTable = circuit.UpdateTruthTable(localTruthTable, nextGate)
					localStates = append(localStates, localTruthTable)
					localDist = CalcComplexity(localTruthTable, targetState)

					if localDist < tourDist {
						tourTruthTable = localTruthTable
						tourStates = localStates
						tourGates = localGates
						tourDist = localDist
					}

				}

			}

			s.depositPheromone(iterationDeposits, tourStates, tourDist)

			if (tourDist < bestDist) || (tourDist == bestDist && len(tourGates) < len(bestGates)) {
				bestStates = tourStates
				bestGates = tourGates
				bestDist = tourDist
			}

		}

		s.updatePheromones(pheromones, iterationDeposits)

	}

	return SynthesisResult{States: bestStates, Gates: bestGates, Complexity: bestDist}
}
