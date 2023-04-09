package circuit

import "log"

type ToffoliGate struct {
	ControlBits []int
	TargetBit   int
}

const (
	ControlBitPositive = 0
	ControlBitNegative = 1
	ControlBitIgnore   = 2
)

func (tg ToffoliGate) Apply(state []int) []int {
	updatedState := make([]int, len(state))
	copy(updatedState, state)

	invert := true
	for controlBit, bitMode := range tg.ControlBits {
		switch bitMode {
		case ControlBitIgnore:
			continue
		case ControlBitPositive:
			if state[controlBit] == 0 {
				invert = false
			}
		case ControlBitNegative:
			if state[controlBit] == 1 {
				invert = false
			}
		default:
			log.Fatalln("unexpected control bit mode:", bitMode)
		}
	}

	if invert {
		updatedState[tg.TargetBit] = (state[tg.TargetBit] + 1) % 2
	}

	return updatedState
}

func UpdateTruthTable(tt TruthTable, gate ToffoliGate) TruthTable {
	res := TruthTable{
		Rows: make([]TruthTableRow, 0),
	}

	for _, row := range tt.Rows {
		resOutput := gate.Apply(row.Output)
		res.Rows = append(res.Rows, TruthTableRow{Input: row.Input, Output: resOutput})
	}

	return res
}
