package circuit

func equal(f0, f1 []int) bool {
	if len(f0) != len(f1) {
		return false
	}

	for i := 0; i < len(f0); i++ {
		if f0[i] != f1[i] {
			return false
		}
	}
	return true
}

type loop struct {
	start int
	end   int
}

func detectLoops(states []TruthTable) []loop {
	loops := make([]loop, 0)
	for i := 0; i < len(states); i++ {
		start := -1
		end := -1

		for j := 0; j < len(states); j++ {
			if equal(states[i].ToVector().Vector, states[j].ToVector().Vector) {
				if start < 0 {
					start = j
					continue
				}
				end = j
			}
		}

		if end > 0 {
			loops = append(loops, loop{start: start, end: end})
		}
	}
	return loops
}

func removeLoops(states []TruthTable, gates []ToffoliGate, loops []loop) ([]TruthTable, []ToffoliGate) {
	markedStates := make([]TruthTable, len(states))
	copy(markedStates, states)
	markedGates := make([]ToffoliGate, len(gates))
	copy(markedGates, gates)

	for _, loop := range loops {
		for i := loop.start; i < loop.end; i++ {
			markedStates[i] = TruthTable{Rows: nil}
			markedGates[i] = ToffoliGate{TargetBit: -1}
		}
	}

	updatedStates := make([]TruthTable, 0)
	updatedGates := make([]ToffoliGate, 0)
	for i := 0; i < len(markedStates); i++ {
		if markedStates[i].Rows != nil {
			updatedStates = append(updatedStates, markedStates[i])
		}

		if i < len(markedStates)-1 && markedGates[i].TargetBit != -1 {
			updatedGates = append(updatedGates, markedGates[i])
		}
	}

	return updatedStates, updatedGates
}

// CutLoops detects and removes unnecessary operations that lead back to the state circuit already had.
// f.e. "A -> B -> C -> D -> B -> E -> F" will be simplified to "A -> B -> E -> F"
func CutLoops(states []TruthTable, gates []ToffoliGate) ([]TruthTable, []ToffoliGate) {
	loops := detectLoops(states)
	if len(loops) > 0 {
		return removeLoops(states, gates, loops)
	}
	return states, gates
}

// Trim removes redundant operations after the final result has been found for the first time.
// f.e. "A -> B -> C -> D -> B" will turn into "A -> B"
func Trim(states []TruthTable, gates []ToffoliGate) ([]TruthTable, []ToffoliGate) {
	finalVector := states[len(states)-1].ToVector().Vector

	trimTo := -1
	for i := 0; i < len(states)-1; i++ {
		if equal(finalVector, states[i].ToVector().Vector) {
			trimTo = i + 1
		}
	}

	if trimTo >= 0 {
		return states[:trimTo], gates[:trimTo-1]
	}
	return states, gates
}
