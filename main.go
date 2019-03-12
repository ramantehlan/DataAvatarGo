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
	"math/rand"
	"time"
)

var dataFileLoc string = "data/dataset.json"
var typeFileLoc string = "data/types.json"
var submissionFileLoc string = "submission.json"

func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	fmt.Println("DataAvatarGo Started!\n")
	defer fmt.Println("\nDataAvatarGo Ending!")

	// Step 1: Read the input json file
	// conver to the byte code
	datasetJson := ReadJson(dataFileLoc)
	typesJson := ReadJson(typeFileLoc)
	var s []Output
	patterns := []string{
		"XXXX XXXX XXXX XXXX",
		"XXXX XXXX XXXX",
		"XXXX XXXX'",
		"XXXX",
		"XX/XX/XXXX",
		"XX/XX/"}

	for i := 0; i < len(datasetJson.Data); i++ {
		var obj Output
		var entity string
		var types string
		found := false
		entry := datasetJson.Data[i]
		fmt.Println("Entry No: ", i)

		doc, _ := prose.NewDocument(entry)
		for _, ent := range doc.Entities() {

			if InArray(ent.Text, patterns) {

				switch ent.Label {
				case "PERSON":
					types = typesJson.Data[5]
					entity = RandomStr(typesJson.Data[5])
				case "ORGANIZATION":
					types = typesJson.Data[19] + " " + typesJson.Data[20]
					entity = RandomStr(typesJson.Data[19])
				case "GPE":
					types = typesJson.Data[19] + " " + typesJson.Data[20]
					entity = RandomStr(typesJson.Data[19])
				}
				found = true
			}
		}

		if !found {
			ran := rand.Intn(len(typesJson.Data))
			types = typesJson.Data[ran]
			entity = RandomStr(types)
		}

		obj.Text = Replace(entry, entity)
		obj.Entity = entity
		obj.Types = types
		s = append(s, obj)
	}

	CreateJson(s)
}
