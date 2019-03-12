package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// To print the error message
func HandleError(err error, errMsg string) bool{
	if err != nil {
		fmt.Println("Error: " + errMsg)
		fmt.Println(err)
		return false
	}
	return true
}

// Read a json file and return the byte data
func ReadJson(file string) InputData{
	fileData, err := os.Open(file)
	HandleError(err, file + " file not found.")
	defer fileData.Close()
	byteData , err := ioutil.ReadAll(fileData)
	HandleError(err, "Error while reading file")
	jsonData := ByteToJson(byteData)
	return jsonData
}

// Parse a json file
func ByteToJson(byteData []byte) InputData{
	var jsonData InputData
	err := json.Unmarshal(byteData, &jsonData)
	HandleError(err, "There was an error decoding the json.")
	return jsonData
}
