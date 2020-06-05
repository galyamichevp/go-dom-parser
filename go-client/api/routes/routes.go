package routes

import (
	"github.com/gin-gonic/gin"
)

//SetupRouter - initialize routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// v1 := r.Group("/v1")
	// {
	//deprecated
	// v1.GET("todo", controllers.GetTodos)

	// v1.GET("resources/:id", controllers.GetResource)
	// v1.GET("resources/", controllers.GetResources)
	// v1.GET("resources/result/:id", controllers.GetResourceResult)
	// v1.POST("resources/download", controllers.PostResourceDownload)
	// v1.POST("resources/parse/:id", controllers.PostResourceParse)
	// }
	return r
}
