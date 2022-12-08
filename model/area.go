package model

type Area struct {
	ID       uint   `gorm:"primary_key"`
	AreaName string `gorm:"varchar(20);not null"`
}
