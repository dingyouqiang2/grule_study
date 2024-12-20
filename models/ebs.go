package models

// 云硬盘费用
type EBS struct {
	BillMode    int     // 按量(0)/包年月(1)
	Region      string  // 区域
	SyshdType   string  // 系统盘类型
	InstanceCnt int     // 系统盘容量
	UnitPrice   float32 // 单价
	CycleCount  int     // 包月时长
	Count       int     // 数量
	TotalCost   float32 //总费用
}

func (e *EBS) GetFieldDescMap() map[string]string {
	fieldDescMap := make(map[string]string)
	fieldDescMap["BillMode"] = "计费模式"
	fieldDescMap["Region"] = "地域"
	fieldDescMap["SyshdType"] = "系统盘类型"
	fieldDescMap["InstanceCnt"] = "系统盘容量"
	fieldDescMap["UnitPrice"] = "单价"
	fieldDescMap["CycleCount"] = "包月时长"
	fieldDescMap["Count"] = "数量"
	return fieldDescMap
}

func (e *EBS) GetFieldDesc(fieldName string) string {
	fieldDescMap := e.GetFieldDescMap()
	return fieldDescMap[fieldName]
}
