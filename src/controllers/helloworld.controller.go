package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// "ginRestFulApi/libs/common/common"
type User struct {
	Username string `json:"username" binding:"required,min=4,max=255"`
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func TestPost(ctx *gin.Context) {
	var user User

	if err := ctx.ShouldBindJSON(&user); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	// fmt.Println(user)

	// var loginCmd LoginCommand
	// c.BindJSON(&loginCmd)

	// fmt.Println(loginCmd)

	ctx.JSON(http.StatusOK, gin.H{"status": 200, "data": user})
	return
}

func ParamTest(c *gin.Context) {

	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)

}
