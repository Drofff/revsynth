package cli

import (
	"fmt"

	"drofff.com/revsynth/circuit"
)

func printRow(printFunc func()) {
	fmt.Print("|")
	printFunc()
	fmt.Print("|\n")
}

func printRowDelim(insCount int) {
	printRow(func() {
		for i := 0; i < insCount; i++ {
			fmt.Print("--------")
		}
		fmt.Printf("-")
	})
}

// DrawTruthTable is sometimes useful for debugging.
func DrawTruthTable(tt circuit.TruthTable) {
	insCount := len(tt.Rows[0].Input)
	printRowDelim(insCount)

	printRow(func() {
		for i := 0; i < insCount; i++ {
			fmt.Printf(" x%v ", i+1)
		}
		fmt.Print("|")
		for i := 0; i < insCount; i++ {
			fmt.Printf(" y%v ", i+1)
		}
	})

	printRowDelim(insCount)

	for i := 0; i < len(tt.Rows); i++ {
		printRow(func() {
			for in := 0; in < insCount; in++ {
				fmt.Printf(" %v  ", tt.Rows[i].Input[in])
			}
			fmt.Print("|")
			for out := 0; out < insCount; out++ {
				fmt.Printf(" %v  ", tt.Rows[i].Output[out])
			}
		})
	}

	printRowDelim(insCount)
}

func DrawCircuit(inputsCount int, gates []circuit.ToffoliGate) {

	for i := 0; i < inputsCount; i++ {
		line := fmt.Sprintf("x%v --", i+1)

		for gateIndex := len(gates) - 1; gateIndex >= 0; gateIndex-- {
			gate := gates[gateIndex]

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
