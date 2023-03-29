package circuit

type ToffoliGate struct {
	ControlBits []int
	TargetBit   int
}

func (tg ToffoliGate) Apply(state []int) []int {
	invert := true
	for _, controlBit := range tg.ControlBits {
		if state[controlBit] == 0 {
			invert = false
		}
	}

	if invert {
		state[tg.TargetBit] = (state[tg.TargetBit] + 1) % 2
	}

	return state
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
