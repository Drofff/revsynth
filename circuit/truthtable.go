package circuit

type TruthTable struct {
	Rows []TruthTableRow
}

type TruthTableRow struct {
	Input  []int
	Output []int
}
