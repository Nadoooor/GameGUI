package JSON

import (
	"encoding/json"
	"os"
)

type His struct {
	Name   string `json:"name"`
	Exe    string `json:"exe"`
	Banner string `json:"banner"`
}

func Save(things []His) {
	data, _ := json.MarshalIndent(things, "", " ")
	os.WriteFile("./json/His.json", data, 0644)
}

func Load() []His {
	data, _ := os.ReadFile("./json/His.json")
	var his []His
	json.Unmarshal(data, &his)
	return his
}
