package network
//Math and Algorithms

import (
	"math"
)

func dotProduct(array1, array2 []float64) float64{						//trusts that arrays are equal
	var sum float64
	for i := 0; i < len(array1); i++{
		sum += array1[i] * array2[i]
	}
	return sum
}

func sigmoid(input float64) float64{
	return 1/(1+math.Exp((-1)*input)) 									// 1 / ( 1 + e^-z )
}

func sigmoid_prime(input float64) float64{
	return sigmoid(input)*(1-sigmoid(input))							// derivative of sigmoid
}