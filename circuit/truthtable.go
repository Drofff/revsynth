package circuit

import (
	"fmt"
	"strconv"
)

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
	cRows := make([]TruthTableRow, 0)
	for _, row := range tt.Rows {
		cRows = append(cRows, row.Copy())
	}
	return TruthTable{Rows: cRows}
}

func binToDec(bins []int) int {
	binStr := ""
	for _, bin := range bins {
		binStr += strconv.Itoa(bin)
	}
	dec, err := strconv.ParseInt(binStr, 2, 32)
	if err != nil {
		panic(err)
	}
	return int(dec)
}

func (tt TruthTable) ToVector() TruthVector {

	ins := make([][]int, 0)
	v := make([]int, 0)

	for _, row := range tt.Rows {
		ins = append(ins, row.Input)
		v = append(v, binToDec(row.Output))
	}

	return TruthVector{Inputs: ins, Vector: v}
}

func (tt TruthTable) Key() string {
	return tt.ToVector().Key()
}

func (tt TruthTable) Equal(otherTt TruthTable) bool {
	return tt.ToVector().Equal(otherTt.ToVector())
}

func (ttr TruthTableRow) Copy() TruthTableRow {
	cin := make([]int, len(ttr.Input))
	copy(cin, ttr.Input)
	cout := make([]int, len(ttr.Output))
	copy(cout, ttr.Output)
	return TruthTableRow{Input: cin, Output: cout}
}

func decToBin(dec int, binLen int) []int {
	bin := make([]int, binLen)

	binStr := strconv.FormatInt(int64(dec), 2)
	binI := binLen - 1
	for strI := len(binStr) - 1; strI >= 0; strI-- {
		if binStr[strI] == '1' {
			bin[binI] = 1
		}

		binI--
	}

	return bin
}

func (tv TruthVector) ToTable() TruthTable {

	tt := TruthTable{Rows: make([]TruthTableRow, 0)}
	outputLen := len(tv.Inputs[0])

	for i := range tv.Vector {
		input := tv.Inputs[i]
		output := decToBin(tv.Vector[i], outputLen)

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

func (tv TruthVector) Equal(otherTv TruthVector) bool {
	return equal(tv.Vector, otherTv.Vector)
}
