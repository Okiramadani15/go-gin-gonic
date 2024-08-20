package boostrap

import (
	"gin-gonic-gorm/configs"
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/database"
	"gin-gonic-gorm/routes"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BoostrapApp() {

	//  LOAD .env FILE
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	//  INIT CONFIG
	configs.InitConfig()

	// DATABASE CONNECTION

	database.ConnectDatabase()

	// INIT GIN ENGINE

	app := gin.Default()

	//  INIT ROUTE

	routes.InitRoute(app)

	// INIT RUN APP

	app.Run(app_config.PORT)
}
