package config

import (
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(mysql.Open("root:@tcp(localhost:3306)/test_golang?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				SlowThreshold:             time.Second, // Slow SQL threshold
				LogLevel:                  logger.Info, // Log level
				IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
				Colorful:                  true,        // Disable color
			},
		),
	})

	if err != nil {
		panic(err)
	}

	DB = db
}

func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
