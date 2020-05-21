package main
//Input Parser

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

import . "testProject/network"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func assert(condition bool, reason string){
	if !condition{
		panic(errors.New(reason))

	}
}

func readFromFile(trainingFileName string, testFileName string) ([]Datum, []Datum){
	trainingData := make([]Datum, 0)
	testData := make([]Datum, 0)

	trainingFile, err := os.Open(trainingFileName)
	check(err)
	defer trainingFile.Close()
	trainingReader := bufio.NewReader(trainingFile)

	for {
		line, _, err := trainingReader.ReadLine()
		if err == io.EOF {
			break
		}
		trainingData = append(trainingData, parseDataEntry(line))
	}

	testFile, err := os.Open(testFileName)
	check(err)
	defer testFile.Close()
	testReader := bufio.NewReader(testFile)

	for {
		line, _, err := testReader.ReadLine()
		if err == io.EOF {
			break
		}
		testData = append(testData, parseDataEntry(line))
	}

	return trainingData, testData
}

func parseDataEntry(unParsed []byte) Datum {
	var inputs []float64
	var output int
	for i := 0; i < len(unParsed); i++ {			//input section, stops before '= x'
		if unParsed[i] == 61 {
			assert(unParsed[i+1] == 32, "expected space character following semicolon")
			output64, err := strconv.ParseInt(string(unParsed[i+2:]), 10, 64)
			output = int(output64)
			check(err)
			break
		}
		if unParsed[i] == 32 {
			continue
		}
		num := string(unParsed[i])
		for unParsed[i+1] != 32 { 						//if next char is not a string
		 num = num + string(unParsed[i+1])
		 i++
		}
		input, err := strconv.ParseFloat(num, 64)
		check(err)
		inputs = append(inputs, input)
	}
	dataPoint := Datum{inputs, output}
	return dataPoint
}

func parseFileNames() (string, string){
	scanner := bufio.NewScanner(os.Stdin)
	var learningFile string
	var testFile string
	fmt.Println("Input learning data file name: ")
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("invalid filename input - %s", err)
		}else{
			learningFile = scanner.Text()
			break
		}
	}
	fmt.Println("Input test data file name: ")
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("invalid filename input - %s", err)
		}else{
			testFile = scanner.Text()
			break
		}
	}
	return learningFile, testFile
}

func parseName() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Input name for new network: ")
	for {
		scanner.Scan()
		if err := scanner.Err(); err != nil {
			fmt.Printf("invalid name input - %s", err)
		} else {
			return scanner.Text()
		}
	}
}

func parseSizes() []int {
	scanner := bufio.NewScanner(os.Stdin)
	sizeInputs := make([]int, 0)
	fmt.Println("Input sizes by layer, return when finished: ")
	for {
		scanner.Scan()
		var input= scanner.Text()
		if input == "" {
			break
		}
		rowSize, err := strconv.Atoi(input)
		if rowSize <= 0 || rowSize > maximumRowSize || err != nil {
			fmt.Println("invalid row size")
		} else {
			sizeInputs = append(sizeInputs, rowSize)
			fmt.Println(sizeInputs)
		}
	}
	return sizeInputs
}

func parseOptions() Settings {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Read settings from file? (settings.json) [y/n]")
	for {
		scanner.Scan()
		if (scanner.Text() == "y") || (scanner.Text() == "Y") {
			return getFileOptions()
		}else if (scanner.Text() == "n") || (scanner.Text() == "N") {
			return getCustomOptions()
		}else if err := scanner.Err(); err != nil {
			fmt.Printf("invalid 'y/n' input - %s\n", err)
		}else{
			fmt.Println("invalid 'y/n' input")
		}
	}
}

func getCustomOptions() Settings {
	settings := Settings{}
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Learning Rate: (return for default)")
	for {
		scanner.Scan()
		if scanner.Text() == ""{
			break
		}
		input, err := strconv.ParseFloat(scanner.Text(), 64)
		if input <= 0 || err != nil {
			fmt.Println("invalid learning rate")
		}else {
			settings.LearningRate = input
			break
		}
	}
	fmt.Println("Number of Epochs for each Dataset: (return for default)")
	for {
		scanner.Scan()
		if scanner.Text() == ""{
			break
		}
		input, err := strconv.Atoi(scanner.Text())
		if input <= 0 || err != nil {
			fmt.Println("invalid number of epochs")
		}else {
			settings.EpochsPerDataset = input
			break
		}
	}
	fmt.Println("Data Batch Size: (return for default)")
	for {
		scanner.Scan()
		if scanner.Text() == ""{
			break
		}
		input, err := strconv.Atoi(scanner.Text())
		if input <= 0 || err != nil {
			fmt.Println("invalid batch size")
		}else {
			settings.BatchSize = input
			break
		}
	}
	return settings
}

func getFileOptions() Settings {
	settings := Settings{}
	file, err := os.Open("settings.json")
	check(err)
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&settings)
	check(err)
	return settings
}