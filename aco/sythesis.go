package aco

import (
	"math/rand"

	"github.com/Drofff/revsynth/circuit"
	"github.com/Drofff/revsynth/logging"
	"github.com/Drofff/revsynth/utils"
)

type Synthesizer struct {
	conf        Config
	gateFactory circuit.GateFactory
	log         logging.Logger
}

// Pheromones key - `FromState` and `ToState` concatenated into a string.
type Pheromones map[string]PheromoneDeposit

// PheromoneDeposit represents pheromone left by an ant on an edge between two states (truth tables).
type PheromoneDeposit struct {
	FromState       circuit.TruthTable
	ToState         circuit.TruthTable
	UsedGate        circuit.Gate
	PheromoneAmount float64
}

type SynthesisResult struct {
	// Complexity is a distance between the last achieved state and zero func.
	Complexity int
	// States contains all truth table changes of each Gates including the (initial) desired truth table
	// and the closest func to zero func that was reached (or maybe even zero func itself).
	States []circuit.TruthTable
	// Gates sequentially transforming the desired truth table to zero func truth table.
	// Please reverse this list before building a circuit.
	Gates []circuit.Gate
}

func NewSynthesizer(conf Config, gateFactory circuit.GateFactory, log logging.Logger) *Synthesizer {
	if conf.AllowedControlBitValues == nil {
		conf.AllowedControlBitValues = circuit.ControlBitValues
	}
	return &Synthesizer{conf: conf, gateFactory: gateFactory, log: log}
}

func (s *Synthesizer) calcTargetBitWeight(desiredState circuit.TruthTable, tt circuit.TruthTable, pheromones Pheromones, tb int) float64 {
	pheromonesSum := 0.0
	bestComplexity := 0
	setComplexity := false

	for _, pheromone := range pheromones {
		if pheromone.FromState.Equal(tt) && utils.ContainsInt(pheromone.UsedGate.TargetBits(), tb) {
			pheromonesSum += pheromone.PheromoneAmount

			complexity := CalcComplexity(pheromone.ToState, desiredState)
			if !setComplexity {
				bestComplexity = complexity
				setComplexity = true
			} else if complexity < bestComplexity {
				bestComplexity = complexity
			}
		}
	}

	return pheromonesSum*s.conf.Alpha + float64(bestComplexity)*s.conf.Beta
}

func (s *Synthesizer) selectTargetBits(desiredState circuit.TruthTable, tt circuit.TruthTable, pheromones Pheromones) []int {
	targetBits := make([]int, 0)

	for tbi := 0; tbi < s.gateFactory.TargetBitsCount; tbi++ {
		tbWeights := make([]float64, 0)
		for tbVal := 0; tbVal < len(tt.Rows[0].Input); tbVal++ {
			tbWeights = append(tbWeights, s.calcTargetBitWeight(desiredState, tt, pheromones, tbVal))
		}

		selectedBit := -1

		weightsSum := sumFloat64(tbWeights)
		if weightsSum == 0.0 {
			for selectedBit == -1 || utils.ContainsInt(targetBits, selectedBit) {
				selectedBit = rand.Intn(len(tt.Rows[0].Input))
			}
			targetBits = append(targetBits, selectedBit)
			continue
		}

		tbProbabilities := make([]float64, 0)
		for _, tbWeight := range tbWeights {
			tbProbabilities = append(tbProbabilities, tbWeight/weightsSum)
		}

		for selectedBit == -1 || utils.ContainsInt(targetBits, selectedBit) {
			selectedBit = chooseRand(tbProbabilities)
		}
		targetBits = append(targetBits, selectedBit)
	}

	return targetBits
}

func (s *Synthesizer) calcControlBitWeight(desiredState circuit.TruthTable, tt circuit.TruthTable, pheromones Pheromones, tb []int, cb int, cbValue int) float64 {
	pheromonesSum := 0.0
	bestComplexity := 0
	setComplexity := false

	for _, pheromone := range pheromones {
		if pheromone.FromState.Equal(tt) && haveSameElements(pheromone.UsedGate.TargetBits(), tb) && pheromone.UsedGate.ControlBits()[cb] == cbValue {
			pheromonesSum += pheromone.PheromoneAmount

			complexity := CalcComplexity(pheromone.ToState, desiredState)
			if !setComplexity {
				bestComplexity = complexity
				setComplexity = true
			} else if complexity < bestComplexity {
				bestComplexity = complexity
			}
		}
	}

	return pheromonesSum*s.conf.Alpha + float64(bestComplexity)*s.conf.Beta
}

func (s *Synthesizer) selectControlBits(desiredState circuit.TruthTable, tt circuit.TruthTable, pheromones Pheromones, tb []int) []int {
	controlBits := make([]int, 0)
	for cb := 0; cb < len(tt.Rows[0].Input); cb++ {
		if utils.ContainsInt(tb, cb) {
			controlBits = append(controlBits, circuit.ControlBitIgnore)
			continue
		}

		cbWeights := make([]float64, 0)
		for _, cbValue := range s.conf.AllowedControlBitValues {
			cbWeights = append(cbWeights, s.calcControlBitWeight(desiredState, tt, pheromones, tb, cb, cbValue))
		}

		weightsSum := sumFloat64(cbWeights)
		if weightsSum == 0.0 {
			randControlBitValue := s.conf.AllowedControlBitValues[rand.Intn(len(s.conf.AllowedControlBitValues))]
			controlBits = append(controlBits, randControlBitValue)
			continue
		}

		cbValueProbs := make([]float64, 0)
		for _, cbWeight := range cbWeights {
			cbValueProbs = append(cbValueProbs, cbWeight/weightsSum)
		}

		selectedCBValue := s.conf.AllowedControlBitValues[chooseRand(cbValueProbs)]
		controlBits = append(controlBits, selectedCBValue)
	}
	return controlBits
}

func (s *Synthesizer) selectGate(desiredState circuit.TruthTable, tt circuit.TruthTable, pheromones Pheromones) circuit.Gate {
	targetBits := s.selectTargetBits(desiredState, tt, pheromones)
	controlBits := s.selectControlBits(desiredState, tt, pheromones, targetBits)
	return s.gateFactory.NewGateFunc(targetBits, controlBits)
}

func (s *Synthesizer) depositPheromone(pheromones Pheromones, states []circuit.TruthTable, gates []circuit.Gate, dist int) {
	var amount float64
	if dist == 0 {
		amount = s.conf.DepositStrength
	} else {
		amount = s.conf.DepositStrength / float64(dist)
	}

	for i := 0; i < len(states)-1; i++ {
		fromState := states[i]
		toState := states[i+1]

		linkKey := fromState.Key() + toState.Key()

		pheromone, exists := pheromones[linkKey]
		if !exists {
			pheromone = PheromoneDeposit{FromState: fromState, ToState: toState, UsedGate: gates[i], PheromoneAmount: 0}
		}

		gatePosition := len(states) - (i + 1) // position from end state, increases pheromone for gates closer to start state.

		pheromone.PheromoneAmount += amount * float64(gatePosition)
		pheromones[linkKey] = pheromone
	}
}

func (s *Synthesizer) updatePheromones(pheromones Pheromones, newDeposits Pheromones) {
	for key, deposit := range pheromones {
		deposit.PheromoneAmount *= 1.0 - s.conf.EvaporationRate
		pheromones[key] = deposit
	}

	for key, newDeposit := range newDeposits {

		deposit, exists := pheromones[key]
		if !exists {
			deposit = PheromoneDeposit{
				FromState:       newDeposit.FromState,
				ToState:         newDeposit.ToState,
				UsedGate:        newDeposit.UsedGate,
				PheromoneAmount: 0,
			}
		}

		deposit.PheromoneAmount += newDeposit.PheromoneAmount
		pheromones[key] = deposit
	}
}

// Synthesise uses desiredVector as a starting point and "zero-state" as the target state.
func (s *Synthesizer) Synthesise(desiredVector circuit.TruthVector) SynthesisResult {

	targetState := circuit.InitZeroTruthTable(desiredVector.Inputs)
	pheromones := Pheromones{}

	bestStates := make([]circuit.TruthTable, 0)
	bestGates := make([]circuit.Gate, 0)
	bestDist := CalcComplexity(desiredVector.ToTable(), targetState)

	s.log.LogDebug("initial state defined..")

	for iteration := 0; iteration < s.conf.NumOfIterations; iteration++ {

		s.log.LogDebug("iteration %v", iteration+1)

		iterationDeposits := Pheromones{}

		for ant := 0; ant < s.conf.NumOfAnts; ant++ {

			s.log.LogDebug("ant %v", ant+1)

			tourTruthTable := desiredVector.ToTable().Copy()
			tourStates := []circuit.TruthTable{tourTruthTable}

			tourGates := make([]circuit.Gate, 0)
			tourDist := CalcComplexity(tourTruthTable, targetState)

			for localLoop := 0; localLoop < s.conf.LocalLoops; localLoop++ {

				if tourDist == 0 {
					// ant has arrived to the final state
					break
				}

				nextGate := s.selectGate(targetState, tourTruthTable, pheromones)

				localGates := make([]circuit.Gate, len(tourGates))
				copy(localGates, tourGates)
				localGates = append(localGates, nextGate)

				localTruthTable := tourTruthTable.Copy()
				localTruthTable = nextGate.Apply(localTruthTable)
				localStates := make([]circuit.TruthTable, len(tourStates))
				copy(localStates, tourStates)
				localStates = append(localStates, localTruthTable)

				localDist := CalcComplexity(localTruthTable, targetState)

				for depth := 0; depth < s.conf.SearchDepth; depth++ {

					nextGate := s.selectGate(targetState, localTruthTable, pheromones)

					localGates = append(localGates, nextGate)
					localTruthTable = nextGate.Apply(localTruthTable)
					localStates = append(localStates, localTruthTable)
					localDist = CalcComplexity(localTruthTable, targetState)

					if localDist < tourDist {
						localStatesOpt, localGatesOpt := circuit.Trim(localStates, localGates)
						localStatesOpt, localGatesOpt = circuit.CutLoops(localStatesOpt, localGatesOpt)

						tourTruthTable = localTruthTable
						tourStates = localStatesOpt
						tourGates = localGatesOpt
						tourDist = localDist
					}

				}

			}

			s.depositPheromone(iterationDeposits, tourStates, tourGates, tourDist)

			if (tourDist < bestDist) || (tourDist == bestDist && len(tourGates) < len(bestGates)) {
				tourStatesOpt, tourGatesOpt := circuit.CutLoops(tourStates, tourGates)

				bestStates = tourStatesOpt
				bestGates = tourGatesOpt
				bestDist = tourDist
			}

		}

		s.updatePheromones(pheromones, iterationDeposits)

		s.log.LogInfof(".")

	}

	s.log.LogInfof("\n")
	return SynthesisResult{States: bestStates, Gates: bestGates, Complexity: bestDist}
}
