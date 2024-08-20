package database

import (
	"fmt"
	"gin-gonic-gorm/configs/db_config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error

	if db_config.DB_DRIVER == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
			db_config.DB_HOST,
			db_config.DB_USER,
			db_config.DB_PASSWORD,
			db_config.DB_NAME,
			db_config.DB_PORT,
		)

		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatalf("can't connect to database: %v", err)
		}

		log.Println("Connected to database")
	} else {
		log.Fatalf("unsupported database driver: %s", db_config.DB_DRIVER)
	}
}

// import (
// 	"fmt"
// 	"gin-gonic-gorm/configs/db_config"
// 	"log"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// var DB *gorm.DB

// func ConnectDatabase() *gorm.DB {
// 	var db *gorm.DB
// 	var errConnection error

// 	if db_config.DB_DRIVER == "postgres" {
// 		dsn := fmt.Sprintf(
// 			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
// 			db_config.DB_HOST,
// 			db_config.DB_USER,
// 			db_config.DB_PASSWORD,
// 			db_config.DB_NAME,
// 			db_config.DB_PORT,
// 		)

// 		db, errConnection = gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	}

// 	if errConnection != nil {
// 		log.Fatalf("can't connect to database: %v", errConnection)
// 	}

// 	log.Println("Connected to database")
// 	return db
// }

// // func ConnectDatabase() {

// // 	var errConnection error

// // 	if db_config.DB_DRIVER == "postgres" {

// // 		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
// // 			db_config.DB_HOST,
// // 			db_config.DB_USER,
// // 			db_config.DB_PASSWORD,
// // 			db_config.DB_NAME,
// // 			db_config.DB_PORT,
// // 		)

// // 		db, errConnection := gorm.Open(postgres.Open(dsn), &gorm.Config{})

// // 	}

// // 	if errConnection != nil {
// // 		panic("can't connect to database")
// // 	}
// // 	log.Println("connect to database")
// // }
