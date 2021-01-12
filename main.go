package main

import (
	"fmt"
	"ginRestFulApi/src/routers"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	//test synxtax
	var ar []int = []int{1, 2, 3, 4, 7}

	fmt.Println(ar[2])

	// load .env environment variables

	// fmt.Println(gin.ReleaseMode)
	gin.SetMode(gin.ReleaseMode)

	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	// initializes database
	// db, _ := database.Initialize()

	port := os.Getenv("PORT")

	app := gin.Default() // create gin app
	// app.Use(database.Inject(db))
	// app.Use(middlewares.JWTMiddleware())
	routers.ApplyRoutes(app) // apply api router

	fmt.Println(port)
	app.Run(":" + port) // listen to given port
}
