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

// ShowData
// @Tags 数据管理
// @Summary 显示当前用户可见的所有数据
// @Description
// @Param pageNum query int  false "当前页数,默认为1"
// @Param pageSize query  int false "每页大小，默认为7"
// @Param id query int true "当前用户id"
// @Param nodeID body  int true "条件筛选: 设备ID"
// @Param dataName body  string true "条件筛选: 物理量"
// @Param time body  string true "条件筛选: 时间:yy.mm.dd hh:mm:ss - yy.mm.dd hh:mm:ss"
// @Accept json
// @Produce json
// @Success 200 {object} SwaggerResponse
// @Router /data [post]
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
		dataName = "data" + dataName[2:]
		if err := db.Table("data").Select("id,area_id,node_id,date," + dataName).Find(&allData).Error; err != nil {
			response.Fail(c, nil, "查询数据失败")
			return
		}
	}
	if len(allData) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 404,
			"msg":  "查询不到相关数据",
		})
		return
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

// DownloadExcel 导出数据
// @Tags 数据管理
// @Summary 导出当前用户的所有可见设备的数据
// @Produce  json
// @Param id query int true "当前用户id"
// @Param nodeID body  int true "条件筛选: 设备ID"
// @Param dataName body  string true "条件筛选: 物理量"
// @Param time body  string true "条件筛选: 时间:yy.mm.dd hh:mm:ss - yy.mm.dd hh:mm:ss"
// @Success 200 {object} SwaggerResponseData "下载成功"
// @Router /data/download/:id [get]
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
		dataName = "data" + dataName[2:]
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

	data1 := "data1"
	data2 := "data2"
	data3 := "data3"
	data4 := "data4"
	data5 := "data5"
	data6 := "data6"
	data7 := "data7"
	data8 := "data8"
	data9 := "data9"
	if nodeID != "0" {
		var phyList []model.Physical
		db.Where("node_id = ?", nodeID).Find(&phyList)
		var fileds [9]string
		for index, v := range phyList {
			fileds[index] = v.DataName
		}
		data1 = fileds[0]
		data2 = fileds[1]
		data3 = fileds[2]
		data4 = fileds[3]
		data5 = fileds[4]
		data6 = fileds[5]
		data7 = fileds[6]
		data8 = fileds[7]
		data9 = fileds[8]
	}
	// 把filterList导入到EXCEL表中
	xlsx := excelize.NewFile()
	xlsx.SetCellValue("Sheet1", "A1", "编号")
	xlsx.SetCellValue("Sheet1", "B1", "区域编号")
	xlsx.SetCellValue("Sheet1", "C1", "设备编号")
	xlsx.SetCellValue("Sheet1", "D1", "日期")
	xlsx.SetCellValue("Sheet1", "E1", data1)
	xlsx.SetCellValue("Sheet1", "F1", data2)
	xlsx.SetCellValue("Sheet1", "G1", data3)
	xlsx.SetCellValue("Sheet1", "H1", data4)
	xlsx.SetCellValue("Sheet1", "I1", data5)
	xlsx.SetCellValue("Sheet1", "J1", data6)
	xlsx.SetCellValue("Sheet1", "K1", data7)
	xlsx.SetCellValue("Sheet1", "L1", data8)
	xlsx.SetCellValue("Sheet1", "M1", data9)

	for i, v := range filterList {
		curIndex := i + 2
		strIndex := strconv.Itoa(curIndex)
		xlsx.SetCellValue("Sheet1", "A"+strIndex, v.ID)
		xlsx.SetCellValue("Sheet1", "B"+strIndex, v.AreaID)
		xlsx.SetCellValue("Sheet1", "C"+strIndex, v.NodeID)
		xlsx.SetCellValue("Sheet1", "D"+strIndex, v.Date)
		xlsx.SetCellValue("Sheet1", "E"+strIndex, v.Data1)
		xlsx.SetCellValue("Sheet1", "F"+strIndex, v.Data2)
		xlsx.SetCellValue("Sheet1", "G"+strIndex, v.Data3)
		xlsx.SetCellValue("Sheet1", "H"+strIndex, v.Data4)
		xlsx.SetCellValue("Sheet1", "I"+strIndex, v.Data5)
		xlsx.SetCellValue("Sheet1", "J"+strIndex, v.Data6)
		xlsx.SetCellValue("Sheet1", "K"+strIndex, v.Data7)
		xlsx.SetCellValue("Sheet1", "L"+strIndex, v.Data8)
		xlsx.SetCellValue("Sheet1", "M"+strIndex, v.Data9)
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
