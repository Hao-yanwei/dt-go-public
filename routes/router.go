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
		// 用户模块路由接口
		router.POST("user/add", v1.AddUser)
		router.GET("user", v1.GetUsers)
		router.PUT("user/:id", v1.EditUser)
		router.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		router.POST("category/add", v1.AddCategory)
		router.GET("category", v1.GetCate)
		router.PUT("category/:id", v1.EditCate)
		router.DELETE("category/:id", v1.DeleteCate)
		// 文章模块路由接口
		router.POST("art/add", v1.AddArticle)
		router.GET("art", v1.GetCateArt)
		router.PUT("art/:id", v1.EditArt)
		router.DELETE("art/:id", v1.DeleteArt)
	}
	r.Run(utils.HttpPort)
}
