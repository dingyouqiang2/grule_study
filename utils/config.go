package utils

import (
	"encoding/json"
	"grule_study/models"
	"io/ioutil"
	"log"
)

type Grules struct {
	Grules []models.RuleForm `json:"grules"`
}

func ReadConfig() {
	
}

func CreateNode() {
}

func CreateGrule() {
	data, err := ioutil.ReadFile("config.json")
	if err != nil {
		log.Println("Error reading file:", err)
		return
	}
	var rawMap map[string]json.RawMessage
	err = json.Unmarshal(data, &rawMap)
	if err != nil {
		log.Println("Error unmarshalling JSON:", err)
		return
	}
	for k, v := range rawMap {
		var g Grules
		if err := json.Unmarshal(v, &g); err != nil {
			log.Printf("Failed to parse key '%s': %v\n", k, v)
			continue
		}
		log.Println(g)
	}
}

func EditGrule() {
}