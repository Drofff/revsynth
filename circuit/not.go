package circuit

import "log"

type notGate struct {
	targetBit int
	// controlBits must only contain ControlBitIgnore values.
	controlBits []int
}

const (
	notTargetBitsCount  = 1
	notControlBitsLimit = 0
	notTypeName         = "not"
)

func NewNotGateFactory() GateFactory {
	return GateFactory{
		GateType: notTypeName,
		NewGateFunc: func(targetBits []int, controlBits []int) Gate {
			if len(targetBits) != notTargetBitsCount {
				log.Fatalln("unexpected target bits count for NOT gate:", len(targetBits))
			}

			if CountControls(controlBits) > notControlBitsLimit {
				log.Fatalln("invalid control bits for NOT:", controlBits)
			}
			return notGate{targetBit: targetBits[0], controlBits: controlBits}
		},
		TargetBitsCount:  notTargetBitsCount,
		ControlBitsLimit: notControlBitsLimit,
	}
}

func (ng notGate) TypeName() string {
	return notTypeName
}

func (ng notGate) TargetBits() []int {
	return []int{ng.targetBit}
}

func (ng notGate) ControlBits() []int {
	return ng.controlBits
}

func (ng notGate) calcNewOutput(output []int) []int {
	updatedOutput := make([]int, len(output))
	copy(updatedOutput, output)

	updatedOutput[ng.targetBit] = (output[ng.targetBit] + 1) % 2
	return updatedOutput
}

func (ng notGate) Apply(tt TruthTable) TruthTable {
	return updateTruthTable(ng.calcNewOutput, tt)
}
