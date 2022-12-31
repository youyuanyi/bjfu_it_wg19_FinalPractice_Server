package model

type Data struct {
	ID     uint    `gorm:"primary_key" json:"id"`
	AreaID uint    `gorm:"not null" json:"area_id"`
	NodeID uint    `gorm:"not null" json:"node_id"`
	Date   Time    `gorm:"type:datetime"  json:"date"`
	Data1  float32 `json:"data1"` //温度
	Data2  float32 `json:"data2"` //湿度
	Data3  float32 `json:"data3"` //降雨量
	Data4  float32 `json:"data4"` //海拔
	Data5  float32 `json:"data5"` // PM2.5
	Data6  float32 `json:"data6"` //风向
	Data7  float32 `json:"data7"` //风速
	Data8  float32 `json:"data8"` //PM 10
	Data9  float32 `json:"data9"` // 压强
}
