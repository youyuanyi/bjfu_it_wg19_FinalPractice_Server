package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _AreasMgr struct {
	*_BaseMgr
}

// AreasMgr open func
func AreasMgr(db *gorm.DB) *_AreasMgr {
	if db == nil {
		panic(fmt.Errorf("AreasMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_AreasMgr{_BaseMgr: &_BaseMgr{DB: db.Table("areas"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_AreasMgr) GetTableName() string {
	return "areas"
}

// Reset 重置gorm会话
func (obj *_AreasMgr) Reset() *_AreasMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_AreasMgr) Get() (result Areas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_AreasMgr) Gets() (results []*Areas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_AreasMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Areas{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_AreasMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithAreaName area_name获取
func (obj *_AreasMgr) WithAreaName(areaName string) Option {
	return optionFunc(func(o *options) { o.query["area_name"] = areaName })
}

// GetByOption 功能选项模式获取
func (obj *_AreasMgr) GetByOption(opts ...Option) (result Areas, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_AreasMgr) GetByOptions(opts ...Option) (results []*Areas, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_AreasMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Areas, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Areas{}).Where(options.query)
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
func (obj *_AreasMgr) GetFromID(id uint) (result Areas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_AreasMgr) GetBatchFromID(ids []uint) (results []*Areas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromAreaName 通过area_name获取内容
func (obj *_AreasMgr) GetFromAreaName(areaName string) (results []*Areas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Where("`area_name` = ?", areaName).Find(&results).Error

	return
}

// GetBatchFromAreaName 批量查找
func (obj *_AreasMgr) GetBatchFromAreaName(areaNames []string) (results []*Areas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Where("`area_name` IN (?)", areaNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_AreasMgr) FetchByPrimaryKey(id uint) (result Areas, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Areas{}).Where("`id` = ?", id).First(&result).Error

	return
}
