package main

import (
	"ginRestFulApi/src/routers"
	"ginRestFulApi/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	router := routers.InitRoute() // apply api router
	gin.SetMode(utils.EnvVar("GIN_MODE", "debug"))
	port := utils.EnvVar("SERVER_PORT", ":8080")
	router.Run(port) // listen to given port
}
