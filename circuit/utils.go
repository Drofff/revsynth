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

func detectLoops(circuit [][]int) []loop {
	loops := make([]loop, 0)
	for i := 0; i < len(circuit); i++ {
		start := -1
		end := -1

		for j := 0; j < len(circuit); j++ {
			if equal(circuit[i], circuit[j]) {
				if start < 0 {
					start = j
					continue
				}
				end = j
			}
		}

		if end >= 0 {
			loops = append(loops, loop{start: start, end: end})
		}
	}
	return loops
}

func removeLoops(circuit [][]int, loops []loop) [][]int {
	markedCircuit := make([][]int, len(circuit))
	copy(markedCircuit, circuit)

	for _, loop := range loops {
		for i := loop.start; i < loop.end; i++ {
			markedCircuit[i] = nil
		}
	}

	updatedCircuit := make([][]int, 0)
	for _, el := range markedCircuit {
		if el != nil {
			updatedCircuit = append(updatedCircuit, el)
		}
	}
	return updatedCircuit
}

// CutLoops detects and removes unnecessary operations that lead back to the state circuit already had.
// f.e. "A -> B -> C -> D -> B -> E -> F" will be simplified to "A -> B -> E -> F"
func CutLoops(circuit [][]int) [][]int {
	loops := detectLoops(circuit)
	if len(loops) > 0 {
		return removeLoops(circuit, loops)
	}
	return circuit
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
		return states[:trimTo], gates[:trimTo]
	}
	return states, gates
}
