package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/tomyalberdi/go-rest-api/controller"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		users := v1.Group("/user")
		{
			users.GET("/", controller.GetUsers)
			users.POST("/", controller.CreateUser)
			users.DELETE("/:id", controller.DeleteUser)

		}
		posts := v1.Group("/post")
		{
			posts.GET("/", controller.GetPosts)
			posts.POST("/", controller.CreatePost)
			posts.DELETE("/:id", controller.DeletePost)
		}
	}
	return r
}
