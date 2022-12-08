package controller

import (
	"WeatherServer/common"
	"WeatherServer/model"
	"WeatherServer/response"
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"strconv"
	"strings"
)

// Register 注册
func Register(c *gin.Context) {
	db := common.GetDB()
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	password := requestUser.Password
	role := requestUser.Role

	// 数据验证
	var user model.User
	db.Where("user_name = ?", userName).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 442,
			"msg":  "用户已存在",
		})
	}
	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 创建用户
	newUser := model.User{
		UserName: userName,
		Password: string(hashedPassword),
		Avatar:   "/images/default_avatar.png",
		Role:     role,
	}
	db.Create(&newUser)
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "注册成功",
	})
}

// Login 登录
func Login(c *gin.Context) {
	db := common.GetDB()
	// 获取参数
	var requestUser model.User
	c.Bind(&requestUser)
	userName := requestUser.UserName
	password := requestUser.Password

	//数据验证
	var user model.User
	db.Where("user_name = ?", userName).First(&user)
	if user.ID == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 442,
			"msg":  "用户不存在",
		})
		return
	}
	// 判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 422,
			"msg":  "密码错误",
		})
		return
	}
	// 发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统异常",
		})
		return
	}
	// 返回结果给前端
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"token": token},
		"msg":  "登录成功",
	})
}

func GetInfo(c *gin.Context) {
	// 获取上下文中的用户信息
	user, _ := c.Get("user")
	// 返回用户信息
	response.Success(c, gin.H{"id": user.(model.User).ID, "avatar": user.(model.User).Avatar, "name": user.(model.User).UserName, "role": user.(model.User).Role}, "登录获取信息成功")
}

// ToStringArray 将自定义类型转化为字符串数组
func ToStringArray(l []string) (a model.Array) {
	for i := 0; i < len(a); i++ {
		l = append(l, a[i])
	}
	return l
}

func GetAllUsers(c *gin.Context) {
	db := common.GetDB()
	// 当前页数
	pageNum, _ := strconv.Atoi(c.DefaultQuery("pageNum", "1"))
	// 每一页的数据条数
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "5"))
	var users []model.User
	var count int

	db.Table("users").Select("id,user_name,role").Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&users)
	db.Model(model.User{}).Count(&count) // 取到的条数

	var userInfos []model.UserInfo
	// 找到所有的用户
	for _, user := range users {
		var userInfo model.UserInfo
		id := user.ID
		var equipments []model.Userequipments
		userInfo.ID = user.ID
		userInfo.UserName = user.UserName
		userInfo.Avatar = user.Avatar
		userInfo.Role = user.Role
		db.Where("uid=?", id).Find(&equipments) // 找到uid所对应的所有eid
		var equip_res string
		for _, eq := range equipments {
			eid := eq.Eid
			equip_res += strconv.Itoa(int(eid)) + " "
		}
		userInfo.Equipments = equip_res
		userInfos = append(userInfos, userInfo)
	}

	response.Success(c, gin.H{
		"users": userInfos,
		"count": count,
	}, "查找成功")

}

// AddUser 管理员添加用户
func AddUser(c *gin.Context) {
	db := common.GetDB()

	var requestUser model.UserInfo
	c.Bind(&requestUser)

	userName := requestUser.UserName
	password := requestUser.Password
	equipment := requestUser.Equipments
	fmt.Println("equipment:", equipment)
	equipmentList := strings.Fields(equipment)
	fmt.Println("equipmentList:", equipmentList)
	// 数据验证
	var user model.User
	db.Where("user_name = ?", userName).First(&user)
	if user.ID != 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 442,
			"msg":  "用户已存在",
		})
		return
	}

	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 创建用户
	newUser := model.User{
		UserName: userName,
		Password: string(hashedPassword),
		Avatar:   "/images/default_avatar.png",
		Role:     1,
	}
	db.Create(&newUser)

	// 先查表，获取uid
	var fuser model.User
	db.Where("user_name = ?", userName).First(&fuser)
	uid := fuser.ID
	for _, value := range equipmentList {
		var UserEquip model.Userequipments
		UserEquip.Uid = uid
		eid, _ := strconv.Atoi(value)
		UserEquip.Eid = uint(eid)
		db.Create(&UserEquip)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "添加成功",
	})
}

// EditUser 管理员修改用户信息
func EditUser(c *gin.Context) {
	db := common.GetDB()
	var requestUser model.UserInfo
	c.Bind(&requestUser)
	// 获取前端传来的数据
	userId := c.Params.ByName("id")
	userName := requestUser.UserName
	password := requestUser.Password
	equipment := requestUser.Equipments
	equipmentList := strings.Fields(equipment)
	// 密码加密
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	// 查找用户
	var user model.User
	if db.Where("id = ?", userId).First(&user).RecordNotFound() {
		response.Fail(c, nil, "用户信息不存在")
		return
	}
	newUser := model.User{
		UserName: userName,
		Password: string(hashedPassword),
		Role:     user.Role,
	}
	// 修改用户信息
	if err := db.Model(&user).Update(newUser).Error; err != nil {
		response.Fail(c, nil, "修改失败")
		return
	}

	// 先删除userEquiments表中所有的和当前用户相关的记录，重新添加
	db.Where("uid= ?", userId).Delete(&model.Userequipments{})
	// 重新添加记录
	for _, value := range equipmentList {
		var UserEquip model.Userequipments
		uid, _ := strconv.Atoi(userId)
		UserEquip.Uid = uint(uid)
		eid, _ := strconv.Atoi(value)
		UserEquip.Eid = uint(eid)
		db.Create(&UserEquip)
	}
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "修改成功",
	})
}

// DelUser 管理员删除用户
func DelUser(c *gin.Context) {
	db := common.GetDB()
	uid := c.Param("id")
	var user model.User
	// 查找该用户
	if db.Where("id= ? ", uid).First(&user).RecordNotFound() {
		response.Fail(c, nil, "用户不存在")
		return
	}
	// 从users表和userequipments表中删除和该用户相关的记录
	if err := db.Delete(&user).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	if err := db.Where("uid= ?", uid).Delete(&model.Userequipments{}).Error; err != nil {
		response.Fail(c, nil, "删除失败")
		return
	}
	response.Success(c, nil, "删除成功")
}
