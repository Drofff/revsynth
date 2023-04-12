package circuit

import "log"

type GateFactory struct {
	NewGateFunc func(targetBits []int, controlBits []int) Gate
	// TargetBitsCount - how many target bits the underlying gate anticipates.
	TargetBitsCount int
}

// Gate is a general representation of a quantum circuit gate element.
type Gate interface {
	// TargetBits returns a slice of position indexes that indicate what circuit lines
	// the resulting bit modification should be applied to.
	TargetBits() []int
	// ControlBits returns a slice in which position of each element corresponds to its circuit line index
	// and the value at that position indicates how the line should be treated using one of `ControlBitValues`.
	ControlBits() []int
	// Apply updates the provided truth table with an additional operation defined by this gate.
	// Returns a new table without modifying the input table.
	Apply(tt TruthTable) TruthTable
}

const (
	// ControlBitPositive - `1` - votes to apply the gate change to the target bits, `0` - votes not to apply.
	ControlBitPositive = 0
	// ControlBitNegative - `0` - votes to apply the gate change to the target bits, `1` - votes not to apply.
	ControlBitNegative = 1
	// ControlBitIgnore - indicates that the line should not be included into the gate's decision.
	ControlBitIgnore = 2
)

// ControlBitValues indicate how to include a bit on a line into the gate's decision.
var ControlBitValues = []int{ControlBitPositive, ControlBitNegative, ControlBitIgnore}

func evalControlBits(state []int, controlBits []int) bool {
	for controlBit, bitMode := range controlBits {
		switch bitMode {
		case ControlBitIgnore:
			continue
		case ControlBitPositive:
			if state[controlBit] == 0 {
				return false
			}
		case ControlBitNegative:
			if state[controlBit] == 1 {
				return false
			}
		default:
			log.Fatalln("unexpected control bit mode:", bitMode)
		}
	}
	return true
}

func updateTruthTable(calcNewOutput func(output []int) []int, tt TruthTable) TruthTable {
	res := TruthTable{
		Rows: make([]TruthTableRow, 0),
	}

	for _, row := range tt.Rows {
		newOutput := calcNewOutput(row.Output)
		res.Rows = append(res.Rows, TruthTableRow{Input: row.Input, Output: newOutput})
	}

	return res
}
