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
	"fmt"
	"gopkg.in/jdkato/prose.v2"
)

var dataFileLoc string = "data/sliceDataset.json"
var typeFileLoc string = "data/types.json"

func main() {
	fmt.Println("DataAvatarGo Started!\n")
	defer fmt.Println("\nDataAvatarGo Ending!")

	// Step 1: Read the input json file
	// conver to the byte code
	dataset := ReadJson(dataFileLoc)
	//types := ReadJson(typeFileLoc)

	for i := 0; i < len(dataset.Data); i++ {
			entry := dataset.Data[i]
			doc, _ := prose.NewDocument(entry)
	 		for _, ent := range doc.Entities() {
			 fmt.Println(ent.Text, ent.Label)
	 		}
	}

}
