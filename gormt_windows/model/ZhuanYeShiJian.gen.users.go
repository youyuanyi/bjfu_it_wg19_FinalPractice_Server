package model

import (
	"context"
	"fmt"
	"gorm.io/gorm"
)

type _UsersMgr struct {
	*_BaseMgr
}

// UsersMgr open func
func UsersMgr(db *gorm.DB) *_UsersMgr {
	if db == nil {
		panic(fmt.Errorf("UsersMgr need init by db"))
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &_UsersMgr{_BaseMgr: &_BaseMgr{DB: db.Table("users"), isRelated: globalIsRelated, ctx: ctx, cancel: cancel, timeout: -1}}
}

// GetTableName get sql table name.获取数据库名字
func (obj *_UsersMgr) GetTableName() string {
	return "users"
}

// Reset 重置gorm会话
func (obj *_UsersMgr) Reset() *_UsersMgr {
	obj.New()
	return obj
}

// Get 获取
func (obj *_UsersMgr) Get() (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).First(&result).Error

	return
}

// Gets 获取批量结果
func (obj *_UsersMgr) Gets() (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Find(&results).Error

	return
}

// //////////////////////////////// gorm replace /////////////////////////////////
func (obj *_UsersMgr) Count(count *int64) (tx *gorm.DB) {
	return obj.DB.WithContext(obj.ctx).Model(Users{}).Count(count)
}

//////////////////////////////////////////////////////////////////////////////////

//////////////////////////option case ////////////////////////////////////////////

// WithID id获取
func (obj *_UsersMgr) WithID(id uint) Option {
	return optionFunc(func(o *options) { o.query["id"] = id })
}

// WithUserName user_name获取
func (obj *_UsersMgr) WithUserName(userName string) Option {
	return optionFunc(func(o *options) { o.query["user_name"] = userName })
}

// WithPassword password获取
func (obj *_UsersMgr) WithPassword(password string) Option {
	return optionFunc(func(o *options) { o.query["password"] = password })
}

// WithAvatar avatar获取
func (obj *_UsersMgr) WithAvatar(avatar string) Option {
	return optionFunc(func(o *options) { o.query["avatar"] = avatar })
}

// WithRole role获取
func (obj *_UsersMgr) WithRole(role uint) Option {
	return optionFunc(func(o *options) { o.query["role"] = role })
}

// GetByOption 功能选项模式获取
func (obj *_UsersMgr) GetByOption(opts ...Option) (result Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).First(&result).Error

	return
}

// GetByOptions 批量功能选项模式获取
func (obj *_UsersMgr) GetByOptions(opts ...Option) (results []*Users, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}

	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query).Find(&results).Error

	return
}

// SelectPage 分页查询
func (obj *_UsersMgr) SelectPage(page IPage, opts ...Option) (resultPage IPage, err error) {
	options := options{
		query: make(map[string]interface{}, len(opts)),
	}
	for _, o := range opts {
		o.apply(&options)
	}
	resultPage = page
	results := make([]Users, 0)
	var count int64 // 统计总的记录数
	query := obj.DB.WithContext(obj.ctx).Model(Users{}).Where(options.query)
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
func (obj *_UsersMgr) GetFromID(id uint) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` = ?", id).First(&result).Error

	return
}

// GetBatchFromID 批量查找
func (obj *_UsersMgr) GetBatchFromID(ids []uint) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` IN (?)", ids).Find(&results).Error

	return
}

// GetFromUserName 通过user_name获取内容
func (obj *_UsersMgr) GetFromUserName(userName string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`user_name` = ?", userName).Find(&results).Error

	return
}

// GetBatchFromUserName 批量查找
func (obj *_UsersMgr) GetBatchFromUserName(userNames []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`user_name` IN (?)", userNames).Find(&results).Error

	return
}

// GetFromPassword 通过password获取内容
func (obj *_UsersMgr) GetFromPassword(password string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`password` = ?", password).Find(&results).Error

	return
}

// GetBatchFromPassword 批量查找
func (obj *_UsersMgr) GetBatchFromPassword(passwords []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`password` IN (?)", passwords).Find(&results).Error

	return
}

// GetFromAvatar 通过avatar获取内容
func (obj *_UsersMgr) GetFromAvatar(avatar string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`avatar` = ?", avatar).Find(&results).Error

	return
}

// GetBatchFromAvatar 批量查找
func (obj *_UsersMgr) GetBatchFromAvatar(avatars []string) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`avatar` IN (?)", avatars).Find(&results).Error

	return
}

// GetFromRole 通过role获取内容
func (obj *_UsersMgr) GetFromRole(role uint) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`role` = ?", role).Find(&results).Error

	return
}

// GetBatchFromRole 批量查找
func (obj *_UsersMgr) GetBatchFromRole(roles []uint) (results []*Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`role` IN (?)", roles).Find(&results).Error

	return
}

//////////////////////////primary index case ////////////////////////////////////////////

// FetchByPrimaryKey primary or index 获取唯一内容
func (obj *_UsersMgr) FetchByPrimaryKey(id uint) (result Users, err error) {
	err = obj.DB.WithContext(obj.ctx).Model(Users{}).Where("`id` = ?", id).First(&result).Error

	return
}
