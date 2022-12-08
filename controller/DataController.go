package controller

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"github.com/gin-gonic/gin"
)

func ShowData(c *gin.Context) {
	db := common.GetDB()
	var dataList []model.Data
	db.Find(&dataList)
	c.JSON(200, gin.H{
		"dataList": dataList,
		"msg":      "OK",
	})
}
