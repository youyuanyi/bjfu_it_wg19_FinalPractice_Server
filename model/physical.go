package model

// 每种数据的物理量及其含义
//type Physical struct {
//	ID           uint   `gorm:"primary_key" json:"id"`
//	DataTag      string `gorm:"varchar(20);not null" json:"DataTag"`
//	PhysicalName string `gorm:"varchar(20);not null" json:"physicalName"`
//	ChineseName  string `gorm:"varchar(20);not null" json:"chinese_name"`
//	Meaning      string `gorm:"varchar(20);not null" json:"meaning"`
//	Conversion   string `gorm:"varchar(20);not null" json:"conversion"`
//}

type Physical struct {
	ID          uint   `gorm:"primary_key" json:"id"`
	NodeID      uint   `gorm:"not null" json:"node_id"`                 // 设备ID
	DataID      uint   `gorm:"not null" json:"data_id"`                 // 数据ID:1-9
	DataName    string `gorm:"varchar(20);not null" json:"dataName"`    // 该数据的物理名字
	DataMeaning string `gorm:"varchar(20);not null" json:"dataMeaning"` // 物理意义
	Conversion  string `gorm:"varchar(20);not null" json:"conversion"`  //	换算方式
}
type SystemTime struct {
	SysTime string `json:"sysTime"`
}
