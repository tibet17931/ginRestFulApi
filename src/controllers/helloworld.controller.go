package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginCommand struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func TestPost(c *gin.Context) {

	var loginCmd LoginCommand

	c.BindJSON(&loginCmd)
	c.JSON(http.StatusOK, gin.H{"status": 200, "data": loginCmd})
}

func ParamTest(c *gin.Context) {

	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)

}
