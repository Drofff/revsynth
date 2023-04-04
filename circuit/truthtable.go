package circuit

import "fmt"

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

func InitZeroTruthTable(inputs [][]int) TruthTable {
	tt := TruthTable{Rows: make([]TruthTableRow, 0)}
	for _, input := range inputs {
		tt.Rows = append(tt.Rows, TruthTableRow{Input: input, Output: input})
	}
	return tt
}

func (tt TruthTable) Copy() TruthTable {
	ctt := TruthTable{Rows: make([]TruthTableRow, 0)}
	for _, row := range tt.Rows {
		ctt.Rows = append(ctt.Rows, row.Copy())
	}
	return ctt
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

func (tt TruthTable) Key() string {
	return tt.ToVector().Key()
}

func (ttr TruthTableRow) Copy() TruthTableRow {
	cin := make([]int, 0)
	copy(cin, ttr.Input)
	cout := make([]int, 0)
	copy(cout, ttr.Output)
	return TruthTableRow{Input: cin, Output: cout}
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

func (tv TruthVector) Key() string {
	return fmt.Sprint(tv.Vector)
}
