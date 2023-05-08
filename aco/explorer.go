package aco

import (
	"github.com/Drofff/revsynth/circuit"
	"github.com/Drofff/revsynth/utils"
)

type visibility struct {
	targetBits  []int
	controlBits []int
	distance    int
}

var visibilityCache = map[string][]visibility{}

func generateTargetBits(numOfBits, numOfLines int) [][]int {
	possibleBitValues := make([]int, 0)
	for line := 0; line < numOfLines; line++ {
		possibleBitValues = append(possibleBitValues, line)
	}

	prevGen := make([][]int, 0)
	for _, val := range possibleBitValues {
		prevGen = append(prevGen, []int{val})
	}

	for bit := 1; bit < numOfBits; bit++ {
		newGen := make([][]int, 0)
		for _, prevGenRow := range prevGen {
			for _, val := range possibleBitValues {
				if !utils.ContainsInt(prevGenRow, val) {
					newGen = append(newGen, append(prevGenRow, val))
				}
			}
		}
		prevGen = newGen
	}
	return prevGen
}

func toPositiveNegativeControls(simpleCombs [][]int, linesCount int) [][]int {
	controlBits := make([][]int, 0)
	for _, simpleComb := range simpleCombs {

		scRes := make([][]int, 0)

		initLine := make([]int, 0)
		for lc := 0; lc < linesCount; lc++ {
			initLine = append(initLine, circuit.ControlBitIgnore)
		}
		scRes = append(scRes, initLine)

		for _, sci := range simpleComb {
			newGen := make([][]int, 0)
			for _, scResRow := range scRes {
				combPos := make([]int, len(scResRow))
				copy(combPos, scResRow)
				combPos[sci] = circuit.ControlBitPositive
				newGen = append(newGen, combPos)

				combNeg := make([]int, len(scResRow))
				copy(combNeg, scResRow)
				combNeg[sci] = circuit.ControlBitNegative
				newGen = append(newGen, combNeg)
			}
			scRes = newGen
		}

		for _, scResRow := range scRes {
			controlBits = append(controlBits, scResRow)
		}
	}
	return controlBits
}

func generateFixedControlBits(numOfBits, numOfLines int) [][]int {
	if numOfBits == 0 {
		res := make([]int, numOfLines)
		for i := 0; i < len(res); i++ {
			res[i] = circuit.ControlBitIgnore
		}
		return [][]int{res}
	}

	simpleCombs := generateTargetBits(numOfBits, numOfLines)
	return toPositiveNegativeControls(simpleCombs, numOfLines)
}

func generateControlBits(numOfBitsLimit, numOfLines int) [][]int {
	controlBits := make([][]int, 0)
	for cbCount := 0; cbCount <= numOfBitsLimit; cbCount++ {
		cbCountBits := generateFixedControlBits(cbCount, numOfLines)
		for _, cbCountBitsRow := range cbCountBits {
			controlBits = append(controlBits, cbCountBitsRow)
		}
	}
	return controlBits
}

func intersect(targetBits, controlBits []int) bool {
	for _, tb := range targetBits {
		if controlBits[tb] != circuit.ControlBitIgnore {
			return true
		}
	}
	return false
}

func exploreVisibility(state, desiredState circuit.TruthTable, gateFactory circuit.GateFactory) []visibility {
	cacheKey := state.Key() + gateFactory.GateType
	cacheValue, cacheFound := visibilityCache[cacheKey]
	if cacheFound {
		return cacheValue
	}

	linesCount := len(state.Rows[0].Input)

	targetBits := generateTargetBits(gateFactory.TargetBitsCount, linesCount)

	cbLimit := gateFactory.ControlBitsLimit
	if cbLimit == circuit.ControlBitsNoLimit {
		cbLimit = linesCount
	}
	controlBits := generateControlBits(cbLimit, linesCount)

	gates := make([]circuit.Gate, 0)
	for tbi := 0; tbi < len(targetBits); tbi++ {
		for cbi := 0; cbi < len(controlBits); cbi++ {
			if intersect(targetBits[tbi], controlBits[cbi]) {
				continue
			}

			gates = append(gates, gateFactory.NewGateFunc(targetBits[tbi], controlBits[cbi]))
		}
	}

	visibilities := make([]visibility, 0)
	for _, gate := range gates {
		nextState := gate.Apply(state)
		visibilities = append(visibilities, visibility{
			targetBits:  gate.TargetBits(),
			controlBits: gate.ControlBits(),
			distance:    CalcComplexity(nextState, desiredState),
		})
	}

	visibilityCache[cacheKey] = visibilities
	return visibilities
}
