package controller

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"WeatherServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func ShowData(c *gin.Context) {
	db := common.GetDB()

	// 获取分页参数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "7"))

	// 获取设备节点，数据种类，时间段参数
	nodeID := c.DefaultQuery("nodeID", "")
	dataName := c.DefaultQuery("dataName", "*")
	timeRange := c.DefaultQuery("time", "")
	var startTime string
	var endTime string
	var t1 time.Time
	var t2 time.Time
	var hasT1 bool = false
	var hasT2 bool = false
	// 若时间范围在
	if timeRange != "" {
		time_split := strings.Split(timeRange, ",")
		startTime = time_split[0]
		endTime = time_split[1]
		t1, _ = time.ParseInLocation("2006-01-02 15:04:05", startTime, time.Local)
		t2, _ = time.ParseInLocation("2006-01-02 15:04:05", endTime, time.Local)
		hasT1 = true
		hasT2 = true
	}

	if dataName == "" {
		dataName = "*"
	}

	var querys []string
	var args []string
	// 若设备节点在
	if nodeID != "null" {
		querys = append(querys, "node_id = ?")
		args = append(args, nodeID)
	}

	// 先获取当前用户的id
	uid := c.Params.ByName("id")
	// 从userEquipments表中查询和该用户相关的设备
	var equipmentList []model.Userequipments
	db.Where("uid= ? ", uid).Find(&equipmentList)
	// 取出所有eid
	var eidList []uint
	// 构造eid-eName列表,作为前端所需的options
	var equipList []model.Node
	if nodeID == "0" {
		for _, equip := range equipmentList {
			eidList = append(eidList, equip.Eid)
		}
	} else {
		iNodeID, _ := strconv.Atoi(nodeID)
		eidList = append(eidList, uint(iNodeID))
	}

	db.Table("nodes").Where("`id` IN (?)", eidList).Find(&equipList)
	// 取出该eidList所对应的所有数据
	var allData []model.Data

	// 先取出所有的
	// 如果没有指定字段
	if dataName == "*" {
		if err := db.Table("data").Find(&allData).Error; err != nil {
			response.Fail(c, nil, "查询数据失败")
			return
		}

	} else {
		// 指定了字段
		if err := db.Table("data").Select("id,area_id,node_id,date," + dataName).Find(&allData).Error; err != nil {
			response.Fail(c, nil, "查询数据失败")
			return
		}
	}
	fmt.Println("len(allData):", len(allData))
	// 取出所有数据后,过滤时间
	var filterTimeList []model.Data
	if hasT1 && hasT2 {
		for _, value := range allData {
			curTime := time.Time(value.Date)
			if t1.Before(curTime) && t2.After(curTime) {
				filterTimeList = append(filterTimeList, value)
			}
		}
	} else {
		filterTimeList = allData
	}
	if len(filterTimeList) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "时间超出范围",
		})
		return
	}

	// allData，判断当前res的node_id是否在eidList中
	var filterList []model.Data
	for _, value := range filterTimeList {
		eid := value.NodeID
		if exists := isInSlice(eid, eidList); exists {
			filterList = append(filterList, value)
		}
	}
	start := (pageNum - 1) * pageSize
	// 当前页数要取得
	var resList []model.Data = filterList[start : start+pageSize]
	var count int
	count = len(filterList)

	// 获取物理量
	var phyList []model.Physical
	db.Find(&phyList)

	response.Success(c, gin.H{
		"phyList":   phyList,
		"equipList": equipList,
		"dataList":  resList,
		"count":     count,
	}, "查找成功")

}

func isInSlice(e uint, eList []uint) bool {
	for _, value := range eList {
		if e == value {
			return true
		}
	}
	return false
}
