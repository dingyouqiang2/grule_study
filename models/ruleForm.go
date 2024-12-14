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
	RuleName       string   `form:"ruleName" json:"rule_name"`
	RuleDesc       string   `form:"ruleDesc" json:"rule_desc"`
	RuleSalience   string   `form:"ruleSalience" json:"rule_salience"`
	RuleConditions []string `form:"ruleCondition" json:"rule_conditions"`
	RuleLogics      []string `form:"ruleLogic" json:"rule_logics"`
}

func (rf *RuleForm) GetFormatGrule() string {
	return fmt.Sprintf(
		gruleFormatString, 
		rf.RuleName, 
		rf.RuleDesc, 
		rf.RuleSalience, 
		strings.Join(rf.RuleConditions, " && "), 
		strings.Join(rf.RuleLogics, ";\n\t\t"),
	)
}