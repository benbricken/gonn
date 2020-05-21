package network
//Data Parsing

import (
	"math/rand"
	"time"
)

type Datum struct{
	Inputs []float64
	Output int
}

func DataShuffle(input []Datum) []Datum {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(input), func(i, j int) { input[i], input[j] = input[j], input[i] })
	return input
}

func InterpretResult(result []float64) int { //returns the most active node, -1 in case of tie
	var max float64
	var decision int
	var tie bool
	for index, activation := range result {
		if activation == max {
			tie = true
		}
		if activation > max {
			max = activation
			decision = index
			tie = false
		}
	}
	if tie{
		return -1
	}
	return decision
}