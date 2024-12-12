package models

type EBSCost struct { // 云硬盘费用
	BillMode    int     // 按量(0)/包年月(1)
	SyshdType   string  // 系统盘类型
	InstanceCnt int     // 系统盘容量
	CycleCount  int     // 包月时长
	Cost        float32 // 费用
}
