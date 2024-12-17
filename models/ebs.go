package models

// 云硬盘费用
type EBS struct {
	BillMode    int     `json:"billMode"`    // 按量(0)/包年月(1)
	Region      string  `json:"region"`      // 区域
	SyshdType   string  `json:"syshdType"`   // 系统盘类型
	InstanceCnt int     `json:"instanceCnt"` // 系统盘容量
	UnitPrice   float32 `json:"unitPrice"`   // 单价
	CycleCount  int     `json:"cycleCount"`  // 包月时长
	Count       int     `json:"count"`       // 数量
}
