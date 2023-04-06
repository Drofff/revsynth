package aco

import (
	"log"
	"math/rand"
)

func sumFloat64(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func probToInt(prob float64) int {
	return int(prob * 100000.0)
}

func chooseRand(probs []float64) int {
	probsSum := 0
	for _, prob := range probs {
		probsSum += probToInt(prob)
	}

	selector := rand.Intn(probsSum)
	for i, prob := range probs {
		selector -= probToInt(prob)
		if selector < 0 {
			return i
		}
	}

	log.Fatalln("unexpected random select result for:", probs)
	return -1
}
