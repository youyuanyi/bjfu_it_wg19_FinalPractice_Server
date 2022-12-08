package model

import (
	"time"
)

// Areas [...]
type Areas struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"-"`
	AreaName string `gorm:"column:area_name" json:"areaName"`
}

// TableName get sql table name.获取数据库表名
func (m *Areas) TableName() string {
	return "areas"
}

// AreasColumns get sql column name.获取数据库列名
var AreasColumns = struct {
	ID       string
	AreaName string
}{
	ID:       "id",
	AreaName: "area_name",
}

// Datas [...]
type Datas struct {
	ID     int       `gorm:"primaryKey;column:id" json:"-"`
	AreaID int       `gorm:"column:areaID" json:"areaId"`
	NodeID int       `gorm:"column:nodeID" json:"nodeId"`
	Date   time.Time `gorm:"column:date" json:"date"`
	Data1  float32   `gorm:"column:data1" json:"data1"`
	Data2  float32   `gorm:"column:data2" json:"data2"`
	Data3  float32   `gorm:"column:data3" json:"data3"`
	Data4  float32   `gorm:"column:data4" json:"data4"`
	Data5  float32   `gorm:"column:data5" json:"data5"`
	Data6  float32   `gorm:"column:data6" json:"data6"`
	Data7  float32   `gorm:"column:data7" json:"data7"`
	Data8  float32   `gorm:"column:data8" json:"data8"`
	Data9  float32   `gorm:"column:data9" json:"data9"`
}

// TableName get sql table name.获取数据库表名
func (m *Datas) TableName() string {
	return "datas"
}

// DatasColumns get sql column name.获取数据库列名
var DatasColumns = struct {
	ID     string
	AreaID string
	NodeID string
	Date   string
	Data1  string
	Data2  string
	Data3  string
	Data4  string
	Data5  string
	Data6  string
	Data7  string
	Data8  string
	Data9  string
}{
	ID:     "id",
	AreaID: "areaID",
	NodeID: "nodeID",
	Date:   "date",
	Data1:  "data1",
	Data2:  "data2",
	Data3:  "data3",
	Data4:  "data4",
	Data5:  "data5",
	Data6:  "data6",
	Data7:  "data7",
	Data8:  "data8",
	Data9:  "data9",
}

// Nodes [...]
type Nodes struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"-"`
	NodeName string `gorm:"column:node_name" json:"nodeName"`
}

// TableName get sql table name.获取数据库表名
func (m *Nodes) TableName() string {
	return "nodes"
}

// NodesColumns get sql column name.获取数据库列名
var NodesColumns = struct {
	ID       string
	NodeName string
}{
	ID:       "id",
	NodeName: "node_name",
}

// Userequipments [...]
type Userequipments struct {
	ID  uint `gorm:"primaryKey;column:id" json:"-"`
	UID uint `gorm:"column:uid" json:"uid"`
	Eid uint `gorm:"column:eid" json:"eid"`
}

// TableName get sql table name.获取数据库表名
func (m *Userequipments) TableName() string {
	return "userequipments"
}

// UserequipmentsColumns get sql column name.获取数据库列名
var UserequipmentsColumns = struct {
	ID  string
	UID string
	Eid string
}{
	ID:  "id",
	UID: "uid",
	Eid: "eid",
}

// Users [...]
type Users struct {
	ID       uint   `gorm:"primaryKey;column:id" json:"-"`
	UserName string `gorm:"column:user_name" json:"userName"`
	Password string `gorm:"column:password" json:"password"`
	Avatar   string `gorm:"column:avatar" json:"avatar"`
	Role     uint   `gorm:"column:role" json:"role"`
}

// TableName get sql table name.获取数据库表名
func (m *Users) TableName() string {
	return "users"
}

// UsersColumns get sql column name.获取数据库列名
var UsersColumns = struct {
	ID       string
	UserName string
	Password string
	Avatar   string
	Role     string
}{
	ID:       "id",
	UserName: "user_name",
	Password: "password",
	Avatar:   "avatar",
	Role:     "role",
}
