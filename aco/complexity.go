package aco

func maxMin(n1, n2 int) (int, int) {
	if n1 > n2 {
		return n1, n2
	}
	return n2, n1
}

func CalcComplexity(actualVector, desiredVector []int) int {
	dist := 0

	elCount := len(desiredVector)
	if len(actualVector) != len(desiredVector) {
		maxL, minL := maxMin(len(actualVector), len(desiredVector))
		dist += maxL - minL
		elCount = minL
	}

	for i := 0; i < elCount; i++ {
		if actualVector[i] != desiredVector[i] {
			dist++
		}
	}

	return dist
}
