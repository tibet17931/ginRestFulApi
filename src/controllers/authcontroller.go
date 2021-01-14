package controllers

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"

	"ginRestFulApi/src/models/entity"
	"ginRestFulApi/src/models/service"
)

//Signup is for user signup
func Signup(c *gin.Context) {

	type signupBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
		Name     string `json:"name"`
	}
	var body signupBody

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entity.User{}
	user.Email = body.Email
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)
	if err != nil {
		log.Fatal(err)
		return
	}

	user.Password = string(hash)
	user.Name = body.Name
	// userservice := service.Userservice{}
	res := service.Create(&user)
	// if err != nil {
	// 	c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	// } else {
	// 	c.JSON(200, gin.H{"result": "ok"})
	// }
	c.JSON(200, gin.H{"result": res})
	return
}
