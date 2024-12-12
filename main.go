package main

import (
	"grule_study/controller"
	"net/http"

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


func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{})
	})
	// 填写grule表单保存到文件里
	r.GET("/grule/form/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "grule.tmpl", gin.H{})
	})
	r.POST("/grule/form/", controller.WriteGruleForm)
	r.GET("/ebs/", func(c *gin.Context) {

	})
	r.POST("/ebs/", func(c *gin.Context) {
		
	})
	// 这里测试前端直接传递grule给后端执行
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
