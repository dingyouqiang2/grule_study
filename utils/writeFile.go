package utils

import (
	"fmt"
	"io/ioutil"
	"log"
)

func WriteGruleFile(ruleName string, grule string) error {
	err := ioutil.WriteFile(fmt.Sprintf("grule/%s.grl", ruleName), []byte(grule), 0644)
	if err != nil {
		log.Println(err)
	}
	return nil
}