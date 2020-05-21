package main
//Neural Net driver function

import (
	"fmt"
)
import . "testProject/network"

const maximumRowSize = 1024
const debug = true

func main() {
	var sizes []int
	var name string
	var settings Settings

	name = parseName()
	sizes = parseSizes()
	settings = parseOptions()

	trainingFileName, testFileName := parseFileNames()
	trainingData, testData := readFromFile(trainingFileName, testFileName)
	if debug {
		fmt.Println(trainingData)
		fmt.Println(testData)
	}

	neuralNet := New(name, sizes, settings)				//make
	neuralNet.PrintNet()

	neuralNet.Learn(trainingData, testData)				//train

	input := []float64 {0.8, 0.7, 0.2}
	outputGraph, _ := neuralNet.FeedForward(input)		//decide
	fmt.Println("Output Graph: ")
	Print2dArray(outputGraph)

	DataShuffle(trainingData)

	if debug {
		fmt.Println(trainingData)
	}

	//Learn()

	//I/O decision loop
}


