package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	c.JSON(http.StatusOK, gin.H{"status": 200, "firstname": firstname, "lastname": lastname})
}
