package boostrap

import (
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/routes"

	"github.com/gin-gonic/gin"
)

func BoostrapApp() {
	app := gin.Default()

	routes.InitRoute(app)

	app.Run(app_config.PORT)
}
