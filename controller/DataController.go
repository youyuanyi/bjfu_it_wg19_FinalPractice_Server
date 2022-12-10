package controller

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"WeatherServer/response"
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
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

	// 先获取当前用户的id
	uid := c.Params.ByName("id")
	// 从userEquipments表中查询和该用户相关的设备
	var equipmentList []model.Userequipments
	db.Where("uid= ? ", uid).Find(&equipmentList)

	// 取出所有eid
	var eidList []uint
	// 构造eid-eName列表,作为前端所需的options
	var equipList []model.Node // 用户当前想查看的
	var eidListAll []uint
	var equipListAll []model.Node
	// 如果没有指定设备
	if nodeID == "0" {
		for _, equip := range equipmentList {
			eidList = append(eidList, equip.Eid)
			eidListAll = append(eidListAll, equip.Eid)

		}
	} else {
		// 指定了设备

		iNodeID, _ := strconv.Atoi(nodeID)
		eidList = append(eidList, uint(iNodeID))
		for _, equip := range equipmentList {
			eidListAll = append(eidListAll, equip.Eid)
		}
	}
	db.Table("nodes").Where("`id` IN (?)", eidListAll).Find(&equipListAll)
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
	var count int
	count = len(filterList)
	// 如果0<count<pageSize
	if count > 0 && count < pageSize {
		pageSize = count
	} else if count == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "查询不到相关数据",
		})
		return
	}
	// 当前页数要取得
	var resList []model.Data = filterList[start : start+pageSize]

	// 获取物理量
	var phyList []model.Physical
	db.Find(&phyList)

	response.Success(c, gin.H{
		"phyList":      phyList,
		"equipList":    equipList,
		"equipListAll": equipListAll,
		"dataList":     resList,
		"count":        count,
	}, "查找成功")

}

func DownloadExcel(c *gin.Context) {
	db := common.GetDB()
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

	// 先获取当前用户的id
	uid := c.Params.ByName("id")
	// 从userEquipments表中查询和该用户相关的设备
	var equipmentList []model.Userequipments
	db.Where("uid= ? ", uid).Find(&equipmentList)

	// 取出所有eid
	var eidList []uint
	// 构造eid-eName列表,作为前端所需的options
	var equipList []model.Node // 用户当前想查看的
	var eidListAll []uint
	var equipListAll []model.Node
	// 如果没有指定设备
	if nodeID == "0" {
		for _, equip := range equipmentList {
			eidList = append(eidList, equip.Eid)
			eidListAll = append(eidListAll, equip.Eid)

		}
	} else {
		// 指定了设备

		iNodeID, _ := strconv.Atoi(nodeID)
		eidList = append(eidList, uint(iNodeID))
		for _, equip := range equipmentList {
			eidListAll = append(eidListAll, equip.Eid)
		}
	}
	db.Table("nodes").Where("`id` IN (?)", eidListAll).Find(&equipListAll)
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

	// 把filterList导入到EXCEL表中
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "编号")
	xlsx.SetCellValue("Sheet1", "B1", "区域编号")
	xlsx.SetCellValue("Sheet1", "C1", "设备编号")
	xlsx.SetCellValue("Sheet1", "D1", "日期")
	xlsx.SetCellValue("Sheet1", "E1", "温度")
	xlsx.SetCellValue("Sheet1", "F1", "湿度")
	xlsx.SetCellValue("Sheet1", "G1", "降雨量")
	xlsx.SetCellValue("Sheet1", "H1", "海拔")
	xlsx.SetCellValue("Sheet1", "I1", "PM2.5")
	xlsx.SetCellValue("Sheet1", "J1", "风向")
	xlsx.SetCellValue("Sheet1", "K1", "风速")
	xlsx.SetCellValue("Sheet1", "L1", "PM10")
	xlsx.SetCellValue("Sheet1", "M1", "压强")

	for i, v := range filterList {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.AreaID)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.NodeID)
		xlsx.SetCellValue("Sheet1", "D"+strIndex, v.Date)
		xlsx.SetCellValue("Sheet1", "E"+strIndex, v.Temperature)
		xlsx.SetCellValue("Sheet1", "F"+strIndex, v.Humidity)
		xlsx.SetCellValue("Sheet1", "G"+strIndex, v.Rainfall)
		xlsx.SetCellValue("Sheet1", "H"+strIndex, v.Altitude)
		xlsx.SetCellValue("Sheet1", "I"+strIndex, v.PM2Dot5)
		xlsx.SetCellValue("Sheet1", "J"+strIndex, v.WindDirection)
		xlsx.SetCellValue("Sheet1", "K"+strIndex, v.WindSpeed)
		xlsx.SetCellValue("Sheet1", "L"+strIndex, v.PM10)
		xlsx.SetCellValue("Sheet1", "M"+strIndex, v.Pressure)
	}
	// 返回文件流
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+"result.xlsx")
	c.Header("Content-Transfer-Encoding", "binary")
	_ = xlsx.Write(c.Writer)
}
func isInSlice(e uint, eList []uint) bool {
	for _, value := range eList {
		if e == value {
			return true
		}
	}
	return false
}
