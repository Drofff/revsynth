package circuit

import "log"

type fredkinGate struct {
	controlBits []int
	// targetBits must always contain two bits.
	targetBits []int
}

const fredkinTargetBitsCount = 2

func NewFredkinGateFactory() GateFactory {
	return GateFactory{
		NewGateFunc: func(targetBits []int, controlBits []int) Gate {
			if len(targetBits) != fredkinTargetBitsCount {
				log.Fatalln("unexpected target bits count:", len(targetBits))
			}
			return fredkinGate{controlBits: controlBits, targetBits: targetBits}
		},
		TargetBitsCount: fredkinTargetBitsCount,
	}
}

func (fg fredkinGate) TargetBits() []int {
	return fg.targetBits
}

func (fg fredkinGate) ControlBits() []int {
	return fg.controlBits
}

func (fg fredkinGate) calcNewOutput(output []int) []int {
	updatedOutput := make([]int, len(output))
	copy(updatedOutput, output)

	if evalControlBits(output, fg.controlBits) {
		updatedOutput[fg.targetBits[0]] = output[fg.targetBits[1]]
		updatedOutput[fg.targetBits[1]] = output[fg.targetBits[0]]
	}

	return updatedOutput
}

func (fg fredkinGate) Apply(tt TruthTable) TruthTable {
	return updateTruthTable(fg.calcNewOutput, tt)
}
