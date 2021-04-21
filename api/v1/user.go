package v1

import (
	//"dt-go-public/model"
	//"dt-go-public/utils/errmsg"
	"dt-go-public/model"
	"dt-go-public/utils/errmsg"
	"github.com/gin-gonic/gin"
	//"github.com/go-playground/validator/v10"
	"net/http"
	//"github.com/go-playground/validator/v10"
	//"net/http"
)

var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	//var msg string
	//var validCode int
	_ = c.ShouldBindJSON(&data)

	//msg, validCode = validator.Validate(&data)
	//if validCode != errmsg.SUCCSE {
	//	c.JSON(
	//		http.StatusOK, gin.H{
	//			"status":  validCode,
	//			"message": msg,
	//		},
	//	)
	//	c.Abort()
	//	return
	//}

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCSE {
		model.CreateUser(&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		code = errmsg.ERROR_USERNAME_USED
	}

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询用户是否存在
func UserExist(c *gin.Context) {

}

// 查询用户列表
func GetUsers(c *gin.Context) {

}

// 查询单个用户

// 编辑用户
func EditUser(c *gin.Context) {

}

// 删除用户
func DeleteUser(c *gin.Context) {

}
