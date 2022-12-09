package model

// 每种数据的物理量及其含义
type Physical struct {
	ID           uint   `gorm:"primary_key" json:"id"`
	PhysicalName string `gorm:"varchar(20);not null" json:"physicalName"`
	ChineseName  string `gorm:"varchar(20);not null" json:"chinese_name"`
	Meaning      string `gorm:"varchar(20);not null" json:"meaning"`
	Conversion   string `gorm:"varchar(20);not null" json:"conversion"`
}
