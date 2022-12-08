package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type _DatasMgr struct {
	*_BaseMgr
}

// DatasMgr open func
func DatasMgr(db *gorm.DB) *_DatasMgr {
	if db == nil {
		panic(fmt.Errorf("DatasMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_DatasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("datas"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_DatasMgr) GetTableName() string {
	return "datas"
}

// Reset 重置gorm会话
func (obj *_DatasMgr) Reset() *_DatasMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_DatasMgr) Get() (result Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_DatasMgr) Gets() (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_DatasMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Datas{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_DatasMgr) WithID(id int) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAreaID areaID获取
func (obj *_DatasMgr) WithAreaID(areaID int) Option {
	return optionFunc(func(o *options) { o.query["areaID"] = areaID })
}

// WithNodeID nodeID获取
func (obj *_DatasMgr) WithNodeID(nodeID int) Option {
	return optionFunc(func(o *options) { o.query["nodeID"] = nodeID })
}

// WithDate date获取
func (obj *_DatasMgr) WithDate(date time.Time) Option {
	return optionFunc(func(o *options) { o.query["date"] = date })
}

// WithData1 data1获取
func (obj *_DatasMgr) WithData1(data1 float32) Option {
	return optionFunc(func(o *options) { o.query["data1"] = data1 })
}

// WithData2 data2获取
func (obj *_DatasMgr) WithData2(data2 float32) Option {
	return optionFunc(func(o *options) { o.query["data2"] = data2 })
}

// WithData3 data3获取
func (obj *_DatasMgr) WithData3(data3 float32) Option {
	return optionFunc(func(o *options) { o.query["data3"] = data3 })
}

// WithData4 data4获取
func (obj *_DatasMgr) WithData4(data4 float32) Option {
	return optionFunc(func(o *options) { o.query["data4"] = data4 })
}

// WithData5 data5获取
func (obj *_DatasMgr) WithData5(data5 float32) Option {
	return optionFunc(func(o *options) { o.query["data5"] = data5 })
}

// WithData6 data6获取
func (obj *_DatasMgr) WithData6(data6 float32) Option {
	return optionFunc(func(o *options) { o.query["data6"] = data6 })
}

// WithData7 data7获取
func (obj *_DatasMgr) WithData7(data7 float32) Option {
	return optionFunc(func(o *options) { o.query["data7"] = data7 })
}

// WithData8 data8获取
func (obj *_DatasMgr) WithData8(data8 float32) Option {
	return optionFunc(func(o *options) { o.query["data8"] = data8 })
}

// WithData9 data9获取
func (obj *_DatasMgr) WithData9(data9 float32) Option {
	return optionFunc(func(o *options) { o.query["data9"] = data9 })
}

// GetByOption 功能选项模式获取
func (obj *_DatasMgr) GetByOption(opts ...Option) (result Datas, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_DatasMgr) GetByOptions(opts ...Option) (results []*Datas, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_DatasMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Datas, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Datas{}).Where(options.query)
	query.Count(&count)
	resultPage.SetTotal(count)
	if len(page.GetOrederItemsString()) > 0 {
		query = query.Order(page.GetOrederItemsString())
	}
	err = query.Limit(int(page.GetSize())).Offset(int(page.Offset())).Find(&results).Error

	resultPage.SetRecords(results)
	return
}

//////////////////////////enume case ////////////////////////////////////////////

// GetFromID 通过id获取内容
func (obj *_DatasMgr) GetFromID(id int) (result Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_DatasMgr) GetBatchFromID(ids []int) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAreaID 通过areaID获取内容
func (obj *_DatasMgr) GetFromAreaID(areaID int) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`areaID` = ?", areaID).Find(&results).Error

	return
}

// GetBatchFromAreaID 批量查找
func (obj *_DatasMgr) GetBatchFromAreaID(areaIDs []int) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`areaID` IN (?)", areaIDs).Find(&results).Error

	return
}

// GetFromNodeID 通过nodeID获取内容
func (obj *_DatasMgr) GetFromNodeID(nodeID int) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`nodeID` = ?", nodeID).Find(&results).Error

	return
}

// GetBatchFromNodeID 批量查找
func (obj *_DatasMgr) GetBatchFromNodeID(nodeIDs []int) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`nodeID` IN (?)", nodeIDs).Find(&results).Error

	return
}

// GetFromDate 通过date获取内容
func (obj *_DatasMgr) GetFromDate(date time.Time) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`date` = ?", date).Find(&results).Error

	return
}

// GetBatchFromDate 批量查找
func (obj *_DatasMgr) GetBatchFromDate(dates []time.Time) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`date` IN (?)", dates).Find(&results).Error

	return
}

// GetFromData1 通过data1获取内容
func (obj *_DatasMgr) GetFromData1(data1 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data1` = ?", data1).Find(&results).Error

	return
}

// GetBatchFromData1 批量查找
func (obj *_DatasMgr) GetBatchFromData1(data1s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data1` IN (?)", data1s).Find(&results).Error

	return
}

// GetFromData2 通过data2获取内容
func (obj *_DatasMgr) GetFromData2(data2 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data2` = ?", data2).Find(&results).Error

	return
}

// GetBatchFromData2 批量查找
func (obj *_DatasMgr) GetBatchFromData2(data2s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data2` IN (?)", data2s).Find(&results).Error

	return
}

// GetFromData3 通过data3获取内容
func (obj *_DatasMgr) GetFromData3(data3 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data3` = ?", data3).Find(&results).Error

	return
}

// GetBatchFromData3 批量查找
func (obj *_DatasMgr) GetBatchFromData3(data3s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data3` IN (?)", data3s).Find(&results).Error

	return
}

// GetFromData4 通过data4获取内容
func (obj *_DatasMgr) GetFromData4(data4 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data4` = ?", data4).Find(&results).Error

	return
}

// GetBatchFromData4 批量查找
func (obj *_DatasMgr) GetBatchFromData4(data4s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data4` IN (?)", data4s).Find(&results).Error

	return
}

// GetFromData5 通过data5获取内容
func (obj *_DatasMgr) GetFromData5(data5 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data5` = ?", data5).Find(&results).Error

	return
}

// GetBatchFromData5 批量查找
func (obj *_DatasMgr) GetBatchFromData5(data5s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data5` IN (?)", data5s).Find(&results).Error

	return
}

// GetFromData6 通过data6获取内容
func (obj *_DatasMgr) GetFromData6(data6 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data6` = ?", data6).Find(&results).Error

	return
}

// GetBatchFromData6 批量查找
func (obj *_DatasMgr) GetBatchFromData6(data6s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data6` IN (?)", data6s).Find(&results).Error

	return
}

// GetFromData7 通过data7获取内容
func (obj *_DatasMgr) GetFromData7(data7 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data7` = ?", data7).Find(&results).Error

	return
}

// GetBatchFromData7 批量查找
func (obj *_DatasMgr) GetBatchFromData7(data7s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data7` IN (?)", data7s).Find(&results).Error

	return
}

// GetFromData8 通过data8获取内容
func (obj *_DatasMgr) GetFromData8(data8 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data8` = ?", data8).Find(&results).Error

	return
}

// GetBatchFromData8 批量查找
func (obj *_DatasMgr) GetBatchFromData8(data8s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data8` IN (?)", data8s).Find(&results).Error

	return
}

// GetFromData9 通过data9获取内容
func (obj *_DatasMgr) GetFromData9(data9 float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data9` = ?", data9).Find(&results).Error

	return
}

// GetBatchFromData9 批量查找
func (obj *_DatasMgr) GetBatchFromData9(data9s []float32) (results []*Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`data9` IN (?)", data9s).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_DatasMgr) FetchByPrimaryKey(id int) (result Datas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Datas{}).Where("`id` = ?", id).First(&result).Error

	return
}
