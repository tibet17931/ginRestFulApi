package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidatorJson(ctx *gin.Context, data interface{}) {
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Please check your json",
		})
		return
	}
}
