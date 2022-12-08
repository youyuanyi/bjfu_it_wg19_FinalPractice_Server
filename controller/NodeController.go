package controller

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"WeatherServer/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ListNodes 展示设备
func ListNodes(c *gin.Context) {
	db := common.GetDB()
	userId := c.Params.ByName("id")
	var user model.User
	// 查找该用户
	if db.Where("id= ? ", userId).First(&user).RecordNotFound() {
		response.Fail(c, nil, "用户不存在")
		return
	}
	// 去userEquipments表中查询该用户的node
	var equipments []model.Userequipments
	db.Where("uid= ? ", userId).Find(&equipments)
	// 获取了user-node，根据node_id去遍历nodes表
	var nodes []model.Node
	for _, value := range equipments {
		eid := value.Eid
		var node model.Node
		db.Where("id= ? ", eid).First(&node)
		nodes = append(nodes, node)
	}
	response.Success(c, gin.H{
		"nodes": nodes,
	}, "查找成功")
}
func GetAllNode(c *gin.Context) {
	db := common.GetDB()
	var nodeList []model.Node
	db.Find(&nodeList)
	response.Success(c, gin.H{
		"nodeList": nodeList,
	}, "登录设备信息成功")
}

func AddNode(c *gin.Context) {
	db := common.GetDB()
	var node model.Node
	c.Bind(&node)
	nodeName := node.NodeName
	// 先去Node表中查找设备是否已存在
	db.Where("node_name = ?", nodeName).First(&node)
	if node.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 442,
			"msg":  "设备已存在",
		})
		return
	}
	newNode := model.Node{
		NodeName: nodeName,
		State:    node.State,
		Duration: node.Duration,
	}
	if err := db.Create(&newNode).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 442,
			"msg":  "添加新设备失败",
		})
		return
	}
	db.Where("node_name= ? ", newNode.NodeName).First(&newNode)
	newUserEquip := model.Userequipments{
		Uid: 1,
		Eid: newNode.ID,
	}
	db.Create(&newUserEquip)
	//
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

func DeleteNode(c *gin.Context) {
	db := common.GetDB()
	eid := c.Param("id")
	var node model.Node
	// 查找该设备
	if db.Where("id= ? ", eid).First(&node).RecordNotFound() {
		response.Fail(c, nil, "设备不存在")
		return
	}
	// 从nodes表和userequipments表中删除和该用户相关的记录
	if err := db.Delete(&node).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	// 从userrequipments表中删除
	if err := db.Where("eid= ?", eid).Delete(&model.Userequipments{}).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}

// EditNode 修改设备信息
func EditNode(c *gin.Context) {
	db := common.GetDB()
	eid := c.Params.ByName("id")
	var node model.Node
	if db.Where("id = ?", eid).First(&node).RecordNotFound() {
		response.Fail(c, nil, "设备信息不存在")
		return
	}

	var requestNode model.Node
	c.Bind(&requestNode)
	newNode := model.Node{
		NodeName: requestNode.NodeName,
		State:    requestNode.State,
		Duration: requestNode.Duration,
	}
	if err := db.Model(&node).Update(newNode).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}
