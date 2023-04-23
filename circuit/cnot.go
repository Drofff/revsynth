package circuit

import "log"

type cnotGate struct {
	targetBit int
	// controlBits must only have one or zero positive or negative control bits set.
	controlBits []int
}

const (
	cnotTargetBitsCount  = 1
	cnotControlBitsLimit = 1
	cnotTypeName         = "cnot"
)

func validateControlBits(controlBits []int) {
	count := 0
	for _, cb := range controlBits {
		if cb == ControlBitPositive || cb == ControlBitNegative {
			count++
		}
	}

	if count > cnotControlBitsLimit {
		log.Fatalln("invalid control bits for CNOT:", controlBits)
	}
}

func NewCnotGateFactory() GateFactory {
	return GateFactory{
		GateType: cnotTypeName,
		NewGateFunc: func(targetBits []int, controlBits []int) Gate {
			if len(targetBits) != cnotTargetBitsCount {
				log.Fatalln("unexpected target bits count:", len(targetBits))
			}
			validateControlBits(controlBits)
			return cnotGate{targetBit: targetBits[0], controlBits: controlBits}
		},
		TargetBitsCount:  cnotTargetBitsCount,
		ControlBitsLimit: cnotControlBitsLimit,
	}
}

func (cg cnotGate) TypeName() string {
	return cnotTypeName
}

func (cg cnotGate) TargetBits() []int {
	return []int{cg.targetBit}
}

func (cg cnotGate) ControlBits() []int {
	return cg.controlBits
}

func (cg cnotGate) calcNewOutput(output []int) []int {
	updatedOutput := make([]int, len(output))
	copy(updatedOutput, output)

	if evalControlBits(output, cg.controlBits) {
		updatedOutput[cg.targetBit] = (output[cg.targetBit] + 1) % 2
	}
	return updatedOutput
}

func (cg cnotGate) Apply(tt TruthTable) TruthTable {
	return updateTruthTable(cg.calcNewOutput, tt)
}
