package circuit

import "log"

type toffoliGate struct {
	controlBits []int
	targetBit   int
}

const (
	toffoliTargetBitsCount = 1
	toffoliTypeName        = "toffoli"
)

func NewToffoliGateFactory() GateFactory {
	return GateFactory{
		GateType: toffoliTypeName,
		NewGateFunc: func(targetBits []int, controlBits []int) Gate {
			if len(targetBits) != toffoliTargetBitsCount {
				log.Fatalln("unexpected target bits count:", len(targetBits))
			}
			return toffoliGate{targetBit: targetBits[0], controlBits: controlBits}
		},
		TargetBitsCount: toffoliTargetBitsCount,
	}
}

func (tg toffoliGate) TypeName() string {
	return toffoliTypeName
}

func (tg toffoliGate) TargetBits() []int {
	return []int{tg.targetBit}
}

func (tg toffoliGate) ControlBits() []int {
	return tg.controlBits
}

func (tg toffoliGate) calcNewOutput(output []int) []int {
	updatedOutput := make([]int, len(output))
	copy(updatedOutput, output)

	if evalControlBits(output, tg.controlBits) {
		updatedOutput[tg.targetBit] = (updatedOutput[tg.targetBit] + 1) % 2
	}

	return updatedOutput
}

func (tg toffoliGate) Apply(tt TruthTable) TruthTable {
	return updateTruthTable(tg.calcNewOutput, tt)
}
