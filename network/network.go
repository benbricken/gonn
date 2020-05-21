package network
//Neural Net Structure, Constructor, and Decision

import (
	"fmt"
	"math/rand"
	"time"
)

const defaultLearningRate = 1.0
const defaultEpochs = 25
const defaultBatchSize = 10

/*
Network Structure
 */
type network struct{
	name             	string
	completedEpochs		int
	sizes            	[]int
	numLayers        	int
	weights          	[][][]float64
	biases           	[][]float64
	learningRate     	float64
	epochsPerDataset 	int
	batchSize        	int
}

type Settings struct {
	LearningRate		float64
	EpochsPerDataset 	int
	BatchSize        	int
}

/*
Network "Constructor" Function
 */
func New(name string, sizes []int, settings Settings) *network{
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	if settings.LearningRate == 0 {
		settings.LearningRate = defaultLearningRate
	}
	if settings.EpochsPerDataset == 0 {
		settings.EpochsPerDataset = defaultEpochs
	}
	if settings.BatchSize == 0 {
		settings.BatchSize = defaultBatchSize
	}
	newNet := network{	name:				name,
						sizes:          	sizes,
						learningRate:     	settings.LearningRate,
						epochsPerDataset: 	settings.EpochsPerDataset,
						batchSize:        	settings.BatchSize}

	fmt.Println("Created new network")
	newNet.numLayers =  len(newNet.sizes)
	for i := 0; i < newNet.numLayers; i++ {
		newNet.biases = append(newNet.biases, make([]float64, sizes[i]))
		if i == 0 {continue}												//input layer receives no bias
		for j := 0; j < sizes[i]; j++ {
			newNet.biases[i][j] = r.NormFloat64()							//set the biases to random mean:0 stddev:1
		}
	}

	newNet.weights = make([][][]float64, newNet.numLayers-1)

	/*
	The first loop (i) sets the length of each 2d array to the number of nodes in the "connectee" column, or the ends of
	each weighted connection

	The second loop (j) sets the length of each array within each 2d array to the number of "connectors" leading to each
	"connectee", or the beginnings of each weighted connection

	The third loop (k) initializes each weight based on a distribution around 0

	The final data structure represents :
		weights [Layer] [Next node] [Previous node] = weight of connection
	*/
	for i := 0; i < len(newNet.weights); i++ {
		newNet.weights[i] = make([][]float64, sizes[i+1])
		for j := 0; j < sizes[i+1]; j++ {
			newNet.weights[i][j] = make([]float64, sizes[i])
			for k := 0; k < sizes[i]; k++ {
				newNet.weights[i][j][k] = r.NormFloat64()
			}
		}
	}

	return &newNet
}

/*
Network Decision Method
*/
func (net *network) FeedForward(networkInput []float64) ([][]float64, []float64) {		//returns: Graph, Output Layer
	networkGraph :=  make([][]float64, len(net.biases))
	for i := range net.biases{
		networkGraph[i] = make([]float64, len(net.biases[i]))
	}

	if len(networkInput) == len(networkGraph[0]){
		copy(networkGraph[0], networkInput)
	}else{
		fmt.Println("invalid Feedforward Input")
		return nil, nil
	}

	for i := 0; i < len(net.weights); i++ {							//for each layer transition
		for j := 0; j < len(net.weights[i]); j++ {					//for each node in the next layer
			networkGraph[i+1][j] =
				sigmoid( dotProduct(networkGraph[i], net.weights[i][j]) + net.biases[i+1][j] )	//x' = sig(w*x + b)
		}
	}

	return networkGraph, networkGraph[len(networkGraph)-1]
}