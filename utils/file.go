package utils

import (
	"fmt"
	"io/ioutil"
	"log"
)

func WriteGruleFile(ruleName string, grule string) error {
	err := ioutil.WriteFile(
		fmt.Sprintf("grule/%s.grl", ruleName),
	    []byte(grule),
		0644,
	)
	if err != nil {
		log.Println(err)
	}
	return nil
}

func ReadGruleFile(ruleName string) string {
	bytes, err := ioutil.ReadFile(fmt.Sprintf("grule/%s.grl", ruleName))
	if err != nil {
		log.Println(err)
	}
	return string(bytes)
}