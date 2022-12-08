package model

// User 和数据库要一起绑定
type User struct {
	ID       uint   `gorm:"primary_key"`
	UserName string `gorm:"varchar(20);not null"`
	Password string `gorm:"size:255;not null"`
	Avatar   string `gorm:"size:255;not null"`
	Role     uint   `gorm:"not null;"` // 0:管理员 , 1: 普通用户
}

type UserInfo struct {
	ID         uint   `json:"id"`
	UserName   string `json:"userName"`
	Password   string `json:"password"`
	Avatar     string `json:"avatar"`
	Role       uint   `json:"role"`
	Equipments string `json:"equipments"`
}
