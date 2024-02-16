package db

import (
	"reblog/internal/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInterface *gorm.DB

func LoadDB() {
	db, err := gorm.Open(sqlite.Open("reblog.db"), &gorm.Config{})

	if err != nil {
		panic("无法连接数据库")
	}

	db.AutoMigrate(&model.Site{}, &model.Article{}, &model.User{})

	dbInterface = db
}

func DB() *gorm.DB {
	if dbInterface == nil {
		LoadDB()
	}

	return dbInterface
}
