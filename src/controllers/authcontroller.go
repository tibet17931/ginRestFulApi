package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"ginRestFulApi/src/models/entity"
)

// func Login(c *gin.Context) {

// 	var loginInfo entity.User
// 	if err := c.ShouldBindJSON(&loginInfo); err != nil {
// 		c.AbortWithStatusJSON(400, gin.H{"error": err.Error()})
// 		return
// 	}
// 	//TODO
// 	userservice := service.Userservice{}
// 	user, errf := userservice.Find(&loginInfo)
// 	if errf != nil {
// 		c.AbortWithStatusJSON(401, gin.H{"error": "Not found"})
// 		return
// 	}

// 	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInfo.Password))
// 	if err != nil {
// 		c.AbortWithStatusJSON(402, gin.H{"error": "Email or password is invalid."})
// 		return
// 	}

// 	token, err := user.GetJwtToken()
// 	if err != nil {
// 		c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
// 		return
// 	}
// 	//-------
// 	c.JSON(200, gin.H{
// 		"token": token,
// 	})
// }

//Profile is to provide current user info

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
	// hash, err := bcrypt.GenerateFromPassword([]byte(info.Password), bcrypt.MinCost)
	// if err != nil {
	// 	log.Fatal(err)
	// 	return
	// }

	// user.Password = string(hash)
	// user.Name = info.Name
	// userservice := service.Userservice{}
	// err = userservice.Create(&user)
	// if err != nil {
	// 	c.AbortWithStatusJSON(500, gin.H{"error": err.Error()})
	// } else {
	// 	c.JSON(200, gin.H{"result": "ok"})
	// }
	// return
}
