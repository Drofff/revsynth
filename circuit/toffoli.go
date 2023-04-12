package circuit

import "log"

type toffoliGate struct {
	controlBits []int
	targetBit   int
}

const toffoliTargetBitsCount = 1

func NewToffoliGateFactory() GateFactory {
	return GateFactory{
		NewGateFunc: func(targetBits []int, controlBits []int) Gate {
			if len(targetBits) != toffoliTargetBitsCount {
				log.Fatalln("unexpected target bits count:", len(targetBits))
			}
			return toffoliGate{targetBit: targetBits[0], controlBits: controlBits}
		},
		TargetBitsCount: toffoliTargetBitsCount,
	}
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

	invert := true
	for controlBit, bitMode := range tg.controlBits {
		switch bitMode {
		case ControlBitIgnore:
			continue
		case ControlBitPositive:
			if output[controlBit] == 0 {
				invert = false
			}
		case ControlBitNegative:
			if output[controlBit] == 1 {
				invert = false
			}
		default:
			log.Fatalln("unexpected control bit mode:", bitMode)
		}
	}

	if invert {
		updatedOutput[tg.targetBit] = (updatedOutput[tg.targetBit] + 1) % 2
	}

	return updatedOutput
}

func (tg toffoliGate) Apply(tt TruthTable) TruthTable {
	res := TruthTable{
		Rows: make([]TruthTableRow, 0),
	}

	for _, row := range tt.Rows {
		rowOutput := tg.calcNewOutput(row.Output)
		res.Rows = append(res.Rows, TruthTableRow{Input: row.Input, Output: rowOutput})
	}

	return res
}
