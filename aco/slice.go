package aco

func sumFloat64(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func chooseRand(probs []float64) int {
	// TODO
}
