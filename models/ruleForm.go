package models

import (
	"fmt"
	"strings"
)

var gruleFormatString = `
rule %s "%s" salience %s {
    when
        %s
    Then
        %s;
}`

type RuleForm struct {
	RuleName       string   `form:"ruleName"`
	RuleDesc       string   `form:"ruleDesc"`
	RuleSalience   string   `form:"ruleSalience"`
	RuleConditions []string `form:"ruleCondition"`
	RuleLogic      []string `form:"ruleLogic"`
}

func (rf *RuleForm) GetFormatGrule() string {
	return fmt.Sprintf(
		gruleFormatString, 
		rf.RuleName, 
		rf.RuleDesc, 
		rf.RuleSalience, 
		strings.Join(rf.RuleConditions, " && "), 
		strings.Join(rf.RuleLogic, ";\n\t\t"),
	)
}