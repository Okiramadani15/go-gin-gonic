package user_controller

import (
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/models"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {

	persons := new([]models.User)

	err := database.DB.Table("person").Find(&persons).Error

	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "internal server error",
		})
	}

	// isValidated := true

	// if !isValidated {

	// 	ctx.AbortWithStatusJSON(400, gin.H{
	// 		"message": "bad request, some field not valid",
	// 	})
	// 	return
	// }

	ctx.JSON(200, gin.H{
		"data": persons,
	})
}
