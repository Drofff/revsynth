package aco

import (
	"log"
	"math/rand"

	"github.com/Drofff/revsynth/utils"
)

func sumFloat64(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	return sum
}

func probToInt(prob float64) int {
	if prob < 0.0001 {
		return 1
	}

	return int(prob * 10000)
}

func chooseRand(probs []float64) int {
	probsSum := 0
	for _, prob := range probs {
		probsSum += probToInt(prob)
	}

	if probsSum == 0 {
		log.Fatalln("probabilities sum must not be 0:", probs)
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

func haveSameElements(uniqueNums0, uniqueNums1 []int) bool {
	if len(uniqueNums0) != len(uniqueNums1) {
		return false
	}

	for _, uniqueNum0 := range uniqueNums0 {
		if !utils.ContainsInt(uniqueNums1, uniqueNum0) {
			return false
		}
	}
	return true
}
