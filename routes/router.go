package routes

import (
	v1 "dt-go-public/api/v1"
	"dt-go-public/middleware"
	"dt-go-public/utils"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	gin.SetMode(utils.AppMode)
	r := gin.Default()

	/*
		后台管理路由接口
	*/
	auth := r.Group("api/v1")
	auth.Use(middleware.JwtToken())
	{
		// 用户模块路由接口

		auth.PUT("user/:id", v1.EditUser)
		auth.DELETE("user/:id", v1.DeleteUser)
		// 分类模块的路由接口
		auth.POST("category/add", v1.AddCategory)
		auth.PUT("category/:id", v1.EditCate)
		auth.DELETE("category/:id", v1.DeleteCate)
		// 文章模块路由接口
		auth.POST("art/add", v1.AddArticle)
		auth.PUT("art/:id", v1.EditArt)
		auth.DELETE("art/:id", v1.DeleteArt)
	}

	/*
		前端展示页面接口
	*/
	router := r.Group("api/v1")
	{
		// 用户信息模块
		router.POST("user/add", v1.AddUser)
		router.GET("user/:id", v1.GetUserInfo)
		router.GET("users", v1.GetUsers)

		// 文章分类信息模块
		router.GET("category", v1.GetCate)
		router.GET("category/:id", v1.GetCateInfo)

		// 文章模块
		router.GET("art", v1.GetArt)
		router.GET("art/list/:id", v1.GetCateArt)
		router.GET("art/info/:id", v1.GetArtInfo)

		// 登录控制模块
		router.POST("login", v1.Login)
		router.POST("loginfront", v1.LoginFront)

		//// 获取个人设置信息
		//router.GET("profile/:id", v1.GetProfile)
		//
		//// 评论模块
		//router.POST("addcomment", v1.AddComment)
		//router.GET("comment/info/:id", v1.GetComment)
		//router.GET("commentfront/:id", v1.GetCommentListFront)
		//router.GET("commentcount/:id", v1.GetCommentCount)
	}

	r.Run(utils.HttpPort)
}
