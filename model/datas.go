package model

type Data struct {
	ID            uint    `gorm:"primary_key"`
	AreaID        uint    `gorm:"not null"`
	NodeID        uint    `gorm:"not null"`
	Date          Time    `gorm:"type:datetime"`
	Temperature   float32 //温度
	Humidity      float32 //湿度
	Rainfall      float32 //降雨量
	Altitude      float32 //海拔
	PM2Dot5       float32 // PM2.5
	WindDirection float32 //风向
	WindSpeed     float32 //风速
	PM10          float32 //PM 10
	Pressure      float32 // 压强
}
