package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

type ExponentData struct {
	Check float64
	Set   float64
}

type RuleForm struct {
	RuleName       string   `form:"ruleName"`
	RuleDesc       string   `form:"ruleDesc"`
	RuleSalience   string   `form:"ruleSalience"`
	RuleConditions []string `form:"ruleCondition"`
	RuleLogic      []string `form:"ruleLogic"`
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	r.GET("/grule/form/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "grule.tmpl", gin.H{})
	})
	r.POST("/grule/form/", func(c *gin.Context) {
		var form RuleForm
		c.ShouldBind(&form)
		grule := fmt.Sprintf(`
rule %s "%s" salience %s {
    when
        %s
    Then
        %s;
}`, 
        form.RuleName, form.RuleDesc, form.RuleSalience, strings.Join(form.RuleConditions, " && "), strings.Join(form.RuleLogic, ";\n\t\t"))
        err := ioutil.WriteFile(fmt.Sprintf("grule/%s.grl", form.RuleName), []byte(grule), 0644)
        if err != nil {
            log.Println(err)
        }
		c.Redirect(http.StatusFound, "/grule/form/")
	})
	r.POST("/number/", func(c *gin.Context) {
		textareaContent := c.PostForm("textarea")
		exponent := &ExponentData{
			Check: 6.67428e-11,
			Set:   0,
		}
		dataContext := ast.NewDataContext()
		dataContext.Add("ExponentData", exponent)
		lib := ast.NewKnowledgeLibrary()
		rb := builder.NewRuleBuilder(lib)
		rb.BuildRuleFromResource("TestExponent", "1.0.0", pkg.NewBytesResource([]byte(textareaContent)))
		eng1 := &engine.GruleEngine{MaxCycle: 5}
		kb, _ := lib.NewKnowledgeBaseInstance("TestExponent", "1.0.0")
		eng1.Execute(dataContext, kb)
		c.JSON(http.StatusOK, gin.H{
			"textareaContent": exponent.Set,
		})
	})
	r.Run()
}
