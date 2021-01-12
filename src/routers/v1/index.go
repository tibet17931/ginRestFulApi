package v1

import (
	"github.com/gin-gonic/gin"

	"ginRestFulApi/src/controllers"
)

func ApplyRoutes(r *gin.RouterGroup) {
	v1 := r.Group("/v1")
	{

		v1.GET("/welcome", controllers.Ping)
		v1.POST("/post", controllers.TestPost)
		v1.GET("/paramTest/:name", controllers.ParamTest)
		// v1.GET("/welcome")
		// auth.ApplyRoutes(v1)
		// posts.ApplyRoutes(v1)
	}
}
