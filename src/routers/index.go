package routers

import (
	"fmt"
	"ginRestFulApi/libs/middlewares"
	"ginRestFulApi/src/controllers"

	"github.com/gin-gonic/gin"
)

func setUpRoute(router *gin.Engine) {
	// authController := new(controllers.he)

	fmt.Println("SET AUTH")
	public := router.Group("/api/v1")
	public.POST("/testpost", controllers.TestPost)
	public.POST("/signup", controllers.Signup)

	private := router.Group("/api/v1")
	private.Use(middlewares.Authentication())
	private.GET("/ping", controllers.Ping)

}

// InitRoute ..
func InitRoute() *gin.Engine {

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	setUpRoute(router)

	// fmt.Println(router)
	return router
}
