package cli

import (
	"fmt"

	"drofff.com/revsynth/circuit"
)

func DrawCircuit(inputsCount int, gates []circuit.ToffoliGate) {

	for i := 0; i < inputsCount; i++ {
		line := fmt.Sprintf("x%v --", i+1)

		for _, gate := range gates {
			if gate.TargetBit == i {
				line += "o--"
			} else if gate.ControlBits[i] == 0 {
				line += "x--"
			} else if gate.ControlBits[i] == 1 {
				line += "X--"
			} else {
				line += "---"
			}
		}

		fmt.Println(line)
	}

	fmt.Println("\nLegend: x - positive control, X - negative control, o - target bit")
}
