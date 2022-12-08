package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UserequipmentsMgr struct {
	*_BaseMgr
}

// UserequipmentsMgr open func
func UserequipmentsMgr(db *gorm.DB) *_UserequipmentsMgr {
	if db == nil {
		panic(fmt.Errorf("UserequipmentsMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UserequipmentsMgr{_BaseMgr: &_BaseMgr{DB: db.Table("userequipments"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UserequipmentsMgr) GetTableName() string {
	return "userequipments"
}

// Reset 重置gorm会话
func (obj *_UserequipmentsMgr) Reset() *_UserequipmentsMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UserequipmentsMgr) Get() (result Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UserequipmentsMgr) Gets() (results []*Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UserequipmentsMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UserequipmentsMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUID uid获取
func (obj *_UserequipmentsMgr) WithUID(uid uint) Option {
	return optionFunc(func(o *options) { o.query["uid"] = uid })
}

// WithEid eid获取
func (obj *_UserequipmentsMgr) WithEid(eid uint) Option {
	return optionFunc(func(o *options) { o.query["eid"] = eid })
}

// GetByOption 功能选项模式获取
func (obj *_UserequipmentsMgr) GetByOption(opts ...Option) (result Userequipments, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UserequipmentsMgr) GetByOptions(opts ...Option) (results []*Userequipments, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UserequipmentsMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Userequipments, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where(options.query)
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
func (obj *_UserequipmentsMgr) GetFromID(id uint) (result Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UserequipmentsMgr) GetBatchFromID(ids []uint) (results []*Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUID 通过uid获取内容
func (obj *_UserequipmentsMgr) GetFromUID(uid uint) (results []*Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where("`uid` = ?", uid).Find(&results).Error

	return
}

// GetBatchFromUID 批量查找
func (obj *_UserequipmentsMgr) GetBatchFromUID(uids []uint) (results []*Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where("`uid` IN (?)", uids).Find(&results).Error

	return
}

// GetFromEid 通过eid获取内容
func (obj *_UserequipmentsMgr) GetFromEid(eid uint) (results []*Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where("`eid` = ?", eid).Find(&results).Error

	return
}

// GetBatchFromEid 批量查找
func (obj *_UserequipmentsMgr) GetBatchFromEid(eids []uint) (results []*Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where("`eid` IN (?)", eids).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UserequipmentsMgr) FetchByPrimaryKey(id uint) (result Userequipments, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Userequipments{}).Where("`id` = ?", id).First(&result).Error

	return
}
