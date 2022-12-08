package controller

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"WeatherServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

func GetPhy(c *gin.Context) {
	db := common.GetDB()
	var phyList []model.Physical
	db.Find(&phyList)
	response.Success(c, gin.H{
		"phyList": phyList,
	}, "查找成功")
}

func EditPhy(c *gin.Context) {
	db := common.GetDB()
	id := c.Params.ByName("id")
	var phy model.Physical
	if db.Where("id = ?", id).First(&phy).RecordNotFound() {
		response.Fail(c, nil, "该物理量不存在")
		return
	}
	var requestPhy model.Physical
	c.Bind(&requestPhy)
	fmt.Println(requestPhy.PhysicalName)
	fmt.Println(requestPhy.Meaning)
	fmt.Println(requestPhy.Conversion)
	newPhy := model.Physical{
		PhysicalName: requestPhy.PhysicalName,
		Meaning:      requestPhy.Meaning,
		Conversion:   requestPhy.Conversion,
	}
	if err := db.Model(&phy).Update(&newPhy).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
