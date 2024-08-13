package user_controller

import "github.com/gin-gonic/gin"

func GetAllUser(ctx *gin.Context) {

	isValidated := true

	if !isValidated {

		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "bad request, some field not valid",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"hello": " hello user",
	})
}
