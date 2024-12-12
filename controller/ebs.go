package controller

import (
	"grule_study/models"
	"grule_study/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func WriteGruleForm(c *gin.Context) {
	var form models.RuleForm
	c.ShouldBind(&form)
	grule := form.GetFormatGrule()
	err := utils.WriteGruleFile(form.RuleName, grule)
	if err != nil {
		log.Println(err)
	}
	c.Redirect(http.StatusFound, "/grule/form/")
}