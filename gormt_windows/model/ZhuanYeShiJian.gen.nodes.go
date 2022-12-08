package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _NodesMgr struct {
	*_BaseMgr
}

// NodesMgr open func
func NodesMgr(db *gorm.DB) *_NodesMgr {
	if db == nil {
		panic(fmt.Errorf("NodesMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_NodesMgr{_BaseMgr: &_BaseMgr{DB: db.Table("nodes"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_NodesMgr) GetTableName() string {
	return "nodes"
}

// Reset 重置gorm会话
func (obj *_NodesMgr) Reset() *_NodesMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_NodesMgr) Get() (result Nodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_NodesMgr) Gets() (results []*Nodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_NodesMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Nodes{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_NodesMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithNodeName node_name获取
func (obj *_NodesMgr) WithNodeName(nodeName string) Option {
	return optionFunc(func(o *options) { o.query["node_name"] = nodeName })
}

// GetByOption 功能选项模式获取
func (obj *_NodesMgr) GetByOption(opts ...Option) (result Nodes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_NodesMgr) GetByOptions(opts ...Option) (results []*Nodes, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_NodesMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Nodes, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where(options.query)
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
func (obj *_NodesMgr) GetFromID(id uint) (result Nodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_NodesMgr) GetBatchFromID(ids []uint) (results []*Nodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromNodeName 通过node_name获取内容
func (obj *_NodesMgr) GetFromNodeName(nodeName string) (results []*Nodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where("`node_name` = ?", nodeName).Find(&results).Error

	return
}

// GetBatchFromNodeName 批量查找
func (obj *_NodesMgr) GetBatchFromNodeName(nodeNames []string) (results []*Nodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where("`node_name` IN (?)", nodeNames).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_NodesMgr) FetchByPrimaryKey(id uint) (result Nodes, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Nodes{}).Where("`id` = ?", id).First(&result).Error

	return
}
