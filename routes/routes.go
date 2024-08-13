package routes

import (
	"gin-gonic-gorm/controllers/book_controller"
	"gin-gonic-gorm/controllers/user_controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(app *gin.Engine) {

	route := app

	route.GET("/user", user_controller.GetAllUser)
	route.GET("/book", book_controller.GetAllBook)

}
