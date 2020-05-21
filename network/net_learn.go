package network
//Learning

func (net *network) Learn(trainingData []Datum, testData []Datum){
	DataShuffle(trainingData)
	DataShuffle(testData)
	var testing bool
	testLength := len(testData)
	if testLength != 0 {
		testing = true
	}
	if testing {}



}

func (net *network) updateNetwork(batch []Datum){
	//TODO: act on SGD
}

//TODO: func (net *network) backPropagate()

func (net *network) Evaluate(testData []Datum) float64{ //returns ratio correct from a dataset
	var correct int
	for _, test := range testData {
		_, outputs := net.FeedForward(test.Inputs)
		if InterpretResult(outputs) == test.Output {
			correct++
		}
	}
	return float64(correct)/float64(len(testData))
}
