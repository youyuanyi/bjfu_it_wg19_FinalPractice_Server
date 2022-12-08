package model

// Userequipments 用户-设备信息绑定表
type Userequipments struct {
	ID  uint `gorm:"primary_key"`
	Uid uint `gorm:"not null"`
	Eid uint `gorm:"not null"`
}
