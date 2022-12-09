package model

type Data struct {
	ID            uint    `gorm:"primary_key" json:"id"`
	AreaID        uint    `gorm:"not null" json:"area_id"`
	NodeID        uint    `gorm:"not null" json:"node_id"`
	Date          Time    `gorm:"type:datetime"  json:"date"`
	Temperature   float32 `json:"temperature"`   //温度
	Humidity      float32 `json:"humidity"`      //湿度
	Rainfall      float32 `json:"rainfall"`      //降雨量
	Altitude      float32 `json:"altitude"`      //海拔
	PM2Dot5       float32 `json:"pm2dot5"`       // PM2.5
	WindDirection float32 `json:"windDirection"` //风向
	WindSpeed     float32 `json:"windSpeed"`     //风速
	PM10          float32 `json:"pm10"`          //PM 10
	Pressure      float32 `json:"pressure"`      // 压强
}
