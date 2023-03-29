package circuit

type TruthTable struct {
	Rows []TruthTableRow
}

type TruthTableRow struct {
	Input  []int
	Output []int
}

type TruthVector struct {
	Inputs [][]int
	Vector []int
}

func (tt TruthTable) ToVector() TruthVector {

	ins := make([][]int, 0)
	v := make([]int, 0)

	for _, outRow := range tt.Rows {

		ins = append(ins, outRow.Input)

		for inRowIndex, inRow := range tt.Rows {

			if equal(outRow.Output, inRow.Input) {
				v = append(v, inRowIndex)
				break
			}

		}

	}

	return TruthVector{Inputs: ins, Vector: v}
}

func (tv TruthVector) ToTable() TruthTable {

	tt := TruthTable{Rows: make([]TruthTableRow, 0)}

	for i := range tv.Vector {
		input := tv.Inputs[i]
		output := tv.Inputs[tv.Vector[i]]

		row := TruthTableRow{
			Input:  input,
			Output: output,
		}
		tt.Rows = append(tt.Rows, row)
	}

	return tt
}
