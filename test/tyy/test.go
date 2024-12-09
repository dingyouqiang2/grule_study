package main

import (
	"log"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

func executeEbsGrule(ebsCost *EBSCost) {
	dataContext := ast.NewDataContext()
	dataContext.Add("EBSCost", ebsCost)

	lib := ast.NewKnowledgeLibrary()
	rb := builder.NewRuleBuilder(lib)
	rb.BuildRuleFromResource("TestEBSCost", "1.0.0", pkg.NewBytesResource([]byte(EbsSsdGenericCostRule)))
	rb.BuildRuleFromResource("TestEBSCost", "1.0.0", pkg.NewBytesResource([]byte(EbsSsdCostRule)))
	rb.BuildRuleFromResource("TestEBSCost", "1.0.0", pkg.NewBytesResource([]byte(EbsSataCostRule)))
	rb.BuildRuleFromResource("TestEBSCost", "1.0.0", pkg.NewBytesResource([]byte(EbsSasCostRule)))
	kb, _ := lib.NewKnowledgeBaseInstance("TestEBSCost", "1.0.0")
	
	eng1 := &engine.GruleEngine{MaxCycle: 5}
	eng1.Execute(dataContext, kb)
}

type EcsCost struct {
	BillMode   int    // 按量/包年月
	Region     string // 可用区
	CPU        int    // CPU核数
	Mem        int    // 内存 GB
	FlavorType string // 规格
	SyshdType  string // 磁盘类型
	Syshd      int    // 磁盘容量
	OrderNum   int    // 数量
	CycleCount int    // 时长
	TotalCost  int    // 总费用
}

/*
		单价获取: 将磁盘容量加1G, 获得每G的价格, 然后乘以40GB就是基本价格
		计算公式: 
		    包月费用 = 包月时长 * 单价 * 磁盘容量
			按量费用 = 按量小时 * 单价 * 磁盘容量

	    通用型SSD磁盘(SSD-generic) 
			0.7/1GB/月 (40GB -- 28元)
			0.00097/1GB/小时 (40GB -- 0.0388)
	    超高IO(SSD) 
			1.2/1GB月 146 (40GB -- 48元)
			0.0017/1GB/小时 (40GB -- 0.068)
	    普通IO(SATA)
			0.3/1GB月 36 (40GB -- 12元)
			0.0005/1GB/小时 (40GB -- 0.02)
	    高IO(SAS)
			0.4/1GB月 48 (40GB -- 16元)
			0.0009/1GB/小时 (40GB -- 0.036)
*/

type EBSCost struct { // 云硬盘费用
	BillMode    int     // 按量(0)/包年月(1)
	SyshdType   string  // 系统盘类型
	InstanceCnt int     // 系统盘容量
	CycleCount  int     // 包月时长
	Hour 		int     // 按量小时
	Cost        float32 // 总费用
}

const (
	EbsSsdGenericCostRule = `
		rule EbsSsdGenericCostRule "通用型SSD磁盘(SSD-generic) 计费规则" salience 10 {
			when
				EBSCost.BillMode == 1 && EBSCost.SyshdType == "SSD-generic"
			Then
				SSD_GENERIC = 0.7;
				EBSCost.Cost = SSD_GENERIC * EBSCost.InstanceCnt * EBSCost.CycleCount;
				Retract("EbsSsdGenericCostRule");
		}
	`
	EbsSsdCostRule = `
		rule EbsSsdCostRule "超高IO(SSD) 计费规则" salience 10 {
			when
				EBSCost.BillMode == 1 && EBSCost.SyshdType == "SSD"
			Then
				SSD = 1.2;
				EBSCost.Cost = SSD * EBSCost.InstanceCnt * EBSCost.CycleCount;
				Retract("EbsSsdCostRule");
		}
	`
	EbsSataCostRule = `
		rule EbsSataCostRule "普通IO(SATA) 计费规则" salience 10 {
			when
				EBSCost.BillMode == 1 && EBSCost.SyshdType == "SATA"
			Then
				SATA = 0.3;
				EBSCost.Cost = SATA * EBSCost.InstanceCnt * EBSCost.CycleCount;
				Retract("EbsSataCostRule");
		}
	`
	EbsSasCostRule = `
		rule EbsSasCostRule "高IO(SAS) 计费规则" salience 10 {
			when
				EBSCost.BillMode == 1 && EBSCost.SyshdType == "SAS"
			Then
				SAS = 0.4;
				EBSCost.Cost = SAS * EBSCost.InstanceCnt * EBSCost.CycleCount;
				Retract("EbsSasCostRule");
		}
	`
)

func main() {
	ebsCost := &EBSCost{
		BillMode:    1,
		SyshdType:   "SSD-generic",
		CycleCount:  3,
		InstanceCnt: 40,
	}
	executeEbsGrule(ebsCost)
	log.Println(ebsCost)
}
