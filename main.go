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
	"regexp"
)

var dataFileLoc string = "data/sliceDataset.json"
var typeFileLoc string = "data/types.json"
var submissionFileLoc string = "submission.json"

func main() {
	fmt.Println("DataAvatarGo Started!\n")
	defer fmt.Println("\nDataAvatarGo Ending!")

	// Step 1: Read the input json file
	// conver to the byte code
	dataset := ReadJson(dataFileLoc)
	types := ReadJson(typeFileLoc)
	var s []Output
	patterns := []string {
											"XXXX XXXX XXXX XXXX",
										  "XXXX XXXX XXXX",
										  "XXXX XXXX'",
										  "XXXX",
										  "XX/XX/XXXX",
										  "XX/XX/",}

	for i := 0; i < len(dataset.Data); i++ {
			var obj Output
			entry := dataset.Data[i]

			var re = regexp.MustCompile(`XXXX`)
			replace := re.ReplaceAllString(entry, `YYYYYYYYYYYYY`)

			doc, _ := prose.NewDocument(entry)
	 		for _, ent := range doc.Entities() {

			 if InArray(ent.Text, patterns) {
				 switch ent.Label {
				 case "PERSON":
					 	obj.Types = types.Data[5]
				 case "ORGANIZATION":
					 obj.Types = types.Data[19] + " " + types.Data[20]
				 case "GPE":
					 obj.Types = types.Data[19] + " " + types.Data[20]
				 }


			 }else{
				  obj.Types = "Types"
			 }

			}

			obj.Text = replace
			obj.Entity = "Entity"

			s = append(s, obj)
	}

	CreateJson(s)
}
