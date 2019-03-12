/*************************************************
*
* This is the main file, where
* the application start.
*
* Author: Raman Tehlan<ramantehlan@gmail.com>
* Date of creation: 12/03/2019 6:21:59 PM
*
**************************************************/
package main

import (
	"encoding/json"
	"fmt"
	"os"
)

var dataFileLoc string = "data/sliceDataset.json"

//var typeFileLoc string = "/data/types.json"

// Structure of dataset.json
type InputData struct {
	inputData []string `json:"inputData"`
}

// Structure of types.json
type TypeData struct {
	typeData []string `json:"typeData"`
}

// Structure of submission.json
type OutputData struct {
	outputData []output `json:"outputData"`
}

// Structure of single entry is submission.json
type Output struct {
	text   string `json:"text"`
	entity string `json:"entity"`
	types  string `json:"types"`
}

func main() {
	fmt.Println("DataAvatarGo Started!\n")
	defer fmt.Println("\nDataAvatarGo Ending!")

	// Step1 :Load dataset
	dataFile, err := os.Open(dataFileLoc)
	// If error in opening the file, handle here
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Success in opening data file: " + dataFileLoc)
	}
	// Close the file in end
	defer dataFile.Close()

	// Step 2: Parse dataset
}
