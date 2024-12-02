package example

import (
    "testing"

    "github.com/hyperjumptech/grule-rule-engine/ast"
    "github.com/hyperjumptech/grule-rule-engine/builder"
    "github.com/hyperjumptech/grule-rule-engine/engine"
    "github.com/hyperjumptech/grule-rule-engine/pkg"
    "github.com/stretchr/testify/assert"
)

const (
    rule = `
rule CalculateCost "Calculate the total cost based on CPU and Memory usage" {
    when 
        CloudResource.CPU > 0 &&
        CloudResource.Memory > 0
    then
        costPerCPUHour = 0.05;
        costPerGBMemoryHour = 0.01;
        cpuCost = CloudResource.CPU * costPerCPUHour;
        memoryCost = CloudResource.Memory * costPerGBMemoryHour;
        CloudResource.TotalCost = cpuCost + memoryCost;
        Retract("CalculateCost");
}
`
)

type CloudResource struct {
    CPU       float64
    Memory    float64
    TotalCost float64
}

func TestCloudResource(t *testing.T) {
    myResource := &CloudResource{
        CPU:    2,
        Memory: 4,
    }
    dataContext := ast.NewDataContext()
    err := dataContext.Add("CloudResource", myResource)
    if err != nil {
        t.Fatal(err)
    }
    lib := ast.NewKnowledgeLibrary()
    ruleBuilder := builder.NewRuleBuilder(lib)
    err = ruleBuilder.BuildRuleFromResource("Test", "0.1.1", pkg.NewBytesResource([]byte(rule)))
    assert.NoError(t, err)
    kb, err := lib.NewKnowledgeBaseInstance("Test", "0.1.1")
    assert.NoError(t, err)
    eng1 := &engine.GruleEngine{MaxCycle: 1}
    err = eng1.Execute(dataContext, kb)
    assert.NoError(t, err)

    // 檢查計算出的總成本
    expectedTotalCost := (2 * 0.05) + (4 * 0.01)
    assert.Equal(t, expectedTotalCost, myResource.TotalCost)
}