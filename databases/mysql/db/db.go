package db

import (
	"fmt"
	"go-clean-architecture/config"
	"go-clean-architecture/databases/mysql/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Connect(cfg *config.Config) *gorm.DB {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MYSQL_USER, cfg.MYSQL_PASSWORD, cfg.MYSQL_HOST, cfg.MYSQL_PORT, cfg.MYSQL_DATABASE,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}

	migrate(db)

	return db
}

func migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&models.TaskModel{},
	)
	if err != nil {
		log.Fatal(err.Error())
	}
}
