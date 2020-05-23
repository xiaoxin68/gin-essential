package router

import (
	"gin-essential/controller"
	"gin-essential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.Use(middleware.CORSMiddleware(), middleware.RecoveryMiddleware())
	r.POST("/api/auth/register", controller.Register)
	r.POST("/api/auth/login", controller.Login)
	r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)

	categoryRoutes := r.Group("/categories")
	categoryController := controller.NewCategoryController()
	categoryRoutes.POST("", categoryController.Create)
	categoryRoutes.PUT("/:id", categoryController.Update) //替换
	categoryRoutes.GET("/:id", categoryController.Show)
	categoryRoutes.GET("", categoryController.PageList)
	categoryRoutes.DELETE("/:id", categoryController.Delete)

	postRoutes := r.Group("/posts")
	//需要登陆后才可以
	postRoutes.Use(middleware.AuthMiddleware())
	postController := controller.NewPostController()
	postRoutes.POST("", postController.Create)
	postRoutes.PUT("/:id", postController.Update) //替换
	postRoutes.GET("/:id", postController.Show)
	postRoutes.GET("", postController.PageList)
	postRoutes.DELETE("/:id", postController.Delete)

	return r
}
