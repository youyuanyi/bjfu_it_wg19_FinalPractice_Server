package controller

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"WeatherServer/response"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetPhy
// @Tags 物理量管理
// @Summary 展示9种物理量的信息
// @Produce  json
// @Success 200 {object} SwaggerResponseData "返回物理量列表"
// @Router /phy [get]
func GetPhy(c *gin.Context) {
	db := common.GetDB()
	// 获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "9"))

	// 先获取当前用户的id
	uid := c.Params.ByName("id")
	// 从userEquipments表中查询和该用户相关的设备
	var equipmentList []model.Userequipments
	db.Where("uid= ? ", uid).Find(&equipmentList)

	// 取出所有eid
	var eidList []uint
	for _, equip := range equipmentList {
		eidList = append(eidList, equip.Eid)
	}

	var phyList []model.Physical
	db.Table("physicals").Select("*").Where("`node_id` IN (?)", eidList).Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&phyList)
	var count int
	db.Model(model.Physical{}).Where("`node_id` IN (?)", eidList).Count(&count)
	response.Success(c, gin.H{
		"phyList": phyList,
		"count":   count,
	}, "查找成功")

}

// EditPhy
// @Tags 物理量管理
// @Summary 管理员修改物理量信息
// @Description
// @Param node body  model.Physical true "物理量信息"
// @Param id path int true "具体修改的物理量的ID"
// @Accept json
// @Produce json
// @Success 200 {object} SwaggerResponse
// @Failure 400 {object} SwaggerResponseData "物理量不存在"
// @Router /phy/:id [put]
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

	newPhy := model.Physical{
		NodeID:      requestPhy.NodeID,
		DataID:      requestPhy.DataID,
		DataName:    requestPhy.DataName,
		DataMeaning: requestPhy.DataMeaning,
		Conversion:  requestPhy.Conversion,
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

func GetPhyByNodeID(c *gin.Context) {
	db := common.GetDB()
	node_id := c.Params.ByName("id")
	var phy []model.Physical
	db.Where("node_id = ?", node_id).Find(&phy)
	count := len(phy)

	response.Success(c, gin.H{"phy": phy, "count": count}, "查找成功")
}
