package main

// Structure of dataset.json and types.json
type InputData struct {
	Data []string `json:"data"`
}

// Structure of submission.json
type OutputData struct {
	Data []Output `json:"DataAvatarGo"`
}

// Structure of single entry is submission.json
type Output struct {
	Text   string `json:"text"`
	Entity string `json:"entity"`
	Types  string `json:"types"`
}
