package configs

import (
	"gin-gonic-gorm/configs/app_config"
	"gin-gonic-gorm/configs/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()

}
