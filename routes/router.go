package routes

import (
	v1 "dt-go-public/api/v1"
	"dt-go-public/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	router := r.Group("api/v1")
	{
		router.POST("user/add", v1.AddUser)
		router.GET("user/get", v1.GetUsers)
		router.PUT("user/edit", v1.EditUser)
		router.DELETE("user/del", v1.DeleteUser)
	}
	r.Run(utils.HttpPort)
}
