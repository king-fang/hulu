package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hulujia/config"
	"hulujia/controller/admin"
	"hulujia/middleware"
)


func SetupRouter() *gin.Engine {

	gin.SetMode(config.App.RunMode)

	route := gin.Default()

	route.Use(middleware.Cors(),middleware.LoggerToFile())

	// 欢迎页
	route.GET("/", func(context *gin.Context) {
		context.String(200,fmt.Sprintf("Welcome to %s apis", config.App.Name))
	})

	//################################
	//#                              #
	//#             API              #
	//#                              #
	//################################
	api := route.Group("/api")
	jwtAuth := middleware.JwtAuth(middleware.LoginStandard)
	api.POST("/auth/login", jwtAuth.LoginHandler)

	jwtApi := api.Group("/")
	jwtApi.Use(jwtAuth.MiddlewareFunc())
	jwtApi.POST("/auth/login/refresh", jwtAuth.RefreshHandler)

	//################################
	//#                              #
	//#       Admin API              #
	//#                              #
	//################################
	adminAPI := jwtApi.Group("/admin")

	adminAPI.Use(middleware.AdminRequired())
	{
		// User Controller
		UserController := &admin.UserController{}
		// 获取用户信息
		adminAPI.POST("user/info", UserController.UserInfo)
		adminAPI.GET("user/:id", UserController.Show)
		adminAPI.POST("user", UserController.Create)
		adminAPI.GET("user", UserController.List)
		adminAPI.PUT("user/:id", UserController.Update)
		adminAPI.DELETE("user/:id", UserController.Destroy)

		// Role Controller
		RoleController := &admin.RoleController{}
		adminAPI.GET("/role",RoleController.List)
		adminAPI.POST("/role",RoleController.Create)
		adminAPI.PUT("/role/:id",RoleController.Update)
		adminAPI.DELETE("/role/:id",RoleController.Destroy)
		adminAPI.GET("/role/:id",RoleController.Show)

	}
	return route
}



