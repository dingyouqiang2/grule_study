package main

import (
	"log"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type ExponentData struct {
	Check float64
	Set   float64
}

const ExponentRule = `
rule  ExponentCheck  "User Related Rule"  salience 10 {
	when 
		ExponentData.Check == 6.67428e-11
	Then
		ExponentData.Set = .12345E+5;
		Retract("ExponentCheck");
}
`

func main() {
	exponent := &ExponentData{
		Check: 6.67428e-11,
		Set:   0,
	}
	dataContext := ast.NewDataContext()
	dataContext.Add("ExponentData", exponent)
	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	rb.BuildRuleFromResource("TestExponent", "1.0.0", pkg.NewBytesResource([]byte(ExponentRule)))
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	kb, _ := lib.NewKnowledgeBaseInstance("TestExponent", "1.0.0")
	eng1.Execute(dataContext, kb)
	log.Println(exponent)
}
