package network
//Display and Access

import (
	"fmt"
)

func Print2dArray(inputArray [][]float64){
	for i := 0; i < len(inputArray); i++ {
		fmt.Println(inputArray[i])
	}
}

/*
Network Print Method
*/
func (net *network) PrintNet() {
	net.PrintName()
	net.PrintEpochs()
	net.PrintLearningSettings()
	net.PrintLayerSizes()
	net.PrintTotalLayers()
	net.PrintBiases()
	net.PrintWeights()
}

func (net *network) PrintName(){
	fmt.Println("Name: ", net.name)
}

func (net *network) PrintEpochs(){
	fmt.Println("Completed Epochs: ", net.completedEpochs)
}

func (net *network) PrintLearningSettings(){
	fmt.Println("Learning Rate: ", net.learningRate)
	fmt.Println("Epochs Per Dataset: ", net.epochsPerDataset)
	fmt.Println("Batch Size: ", net.batchSize)
}

func (net *network) PrintLayerSizes(){
	fmt.Println("Layer sizes: ", net.sizes)
}

func (net *network) PrintTotalLayers(){
	fmt.Println("Total Layers: ", net.numLayers)
}

func (net *network) PrintBiases(){
	fmt.Println("Bias Graph, top to bottom: ")
	for i := 0; i < net.numLayers; i++ {
		fmt.Println(net.biases[i])
	}
}

func (net *network) PrintWeights(){
	fmt.Println("Weight Graph, top to bottom: ")
	for i := 0; i < net.numLayers -1; i++ {
		fmt.Println(net.weights[i])
	}
}