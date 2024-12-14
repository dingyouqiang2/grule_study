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

func CreateGrule(key string, newRule models.RuleForm) {
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
	if _, exists := rawMap[key]; !exists {
		g := Grules{}
		g.Grules = append(g.Grules, newRule)
	}
	updatedJSON, err := json.MarshalIndent(g, "", "  ")
	if err != nil {
		log.Println("Error marshalling updated grules:", err)
		return
	}
	rawMap[key] = updatedJSON
	finalJSON, err := json.MarshalIndent(rawMap, "", "  ")
	if err != nil {
		log.Panicln("Error marshalling final JSON:", err)
		return
	}
	err = ioutil.WriteFile("config.json", finalJSON, 0644)
	if err != nil {
		log.Panicln("Error writing to file:", err)
		return
	}
}

func EditGrule() {
}