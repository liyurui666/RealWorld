package service

import (
	"RealWorld/common"
	"RealWorld/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Register 用户注册逻辑
func Register(context *gin.Context) {
	var body struct {
		User models.User `json:"user"`
	}
	//获取前端的json转为结构体
	context.BindJSON(&body)

	//TODO 数据校验（暂没做）

	//注册前检查邮箱是否已存在
	result := models.QueryEmail(body.User)
	if result.Id != 0 {
		context.JSON(200, gin.H{
			"code":    422,
			"message": "邮箱已被注册",
		})
		return
	}
	//注册
	models.UserRegistration(&body.User)
	//返回用户
	context.JSON(200, gin.H{"user": body.User})
}

// Login 用户登录链接
func Login(context *gin.Context) {
	var body struct {
		User models.User `json:"user"`
	}
	//获取前端的json转为结构体
	context.BindJSON(&body)
	//TODO 数据校验（暂没做）

	//查询数据库 登录
	err := models.Login(&body.User)
	//判断是否登录成功
	if err != nil {
		context.JSON(200, gin.H{
			"code": -1,
			"msg":  "邮箱或者密码错误",
		})
		return
	}
	//TODO JWT
	token, err := common.GetToken(body.User.Username, body.User.Email)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"code": 500,
			"msg":  "系统错误",
		})
		return
	}

	//返回用户
	context.JSON(200, gin.H{"user": body.User, "token": token})
}

// CurrentUser 获取当前用户
func CurrentUser(context *gin.Context) {
	user, _ := context.Get("user")

	context.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": gin.H{"user": user},
	})

}

// UpdateUserInfo 更新用户信息
func UpdateUserInfo(context *gin.Context) {
	var body struct {
		User models.User `json:"user"`
	}
	//获取前端的json转为结构体
	context.BindJSON(&body)
	//更新数据库
	models.UpdateUserInfo(&body.User)
	context.JSON(200, gin.H{"user": body.User})
}
