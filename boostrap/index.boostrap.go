package boostrap

import (
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BoostrapApp() {
	database.ConnectDatabase()

	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
