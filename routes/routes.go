package routes

import (
	"github.com/gin-gonic/gin"
	"go-server-example/handlers"
	"go-server-example/middleware"
)

// SetupRoutes 设置路由
func SetupRoutes() *gin.Engine {
	router := gin.New()
	
	// 使用中间件
	router.Use(middleware.Logger())
	router.Use(middleware.Recovery())
	router.Use(middleware.CORS())
	router.Use(middleware.RateLimit())
	
	// 健康检查
	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
			"message": "Server is running",
		})
	})
	
	// API v1 路由组
	v1 := router.Group("/api/v1")
	{
		// 用户相关路由
		users := v1.Group("/users")
		{
			userHandler := handlers.NewUserHandler()
			users.GET("", userHandler.GetUsers)
			users.GET("/:id", userHandler.GetUser)
			users.POST("", userHandler.CreateUser)
			users.PUT("/:id", userHandler.UpdateUser)
			users.DELETE("/:id", userHandler.DeleteUser)
		}
		
		// 文章相关路由
		posts := v1.Group("/posts")
		{
			postHandler := handlers.NewPostHandler()
			posts.GET("", postHandler.GetPosts)
			posts.GET("/:id", postHandler.GetPost)
			posts.POST("", postHandler.CreatePost)
			posts.PUT("/:id", postHandler.UpdatePost)
			posts.DELETE("/:id", postHandler.DeletePost)
			posts.GET("/user/:user_id", postHandler.GetPostsByUser)
		}
	}
	
	return router
} 