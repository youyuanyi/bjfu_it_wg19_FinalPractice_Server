package model

type Node struct {
	ID       uint   `gorm:"primary_key" json:"id"`
	NodeName string `gorm:"varchar(20);not null" json:"nodeName"`
	State    int    `gorm:"not null" json:"state"`
	Duration int    `gorm:"not null" json:"duration"`
}
