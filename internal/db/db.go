package db

import (
	"reblog/internal/config"
	"reblog/internal/model"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/mongodb"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInterface *gorm.DB

func LoadDB() {
	var db *gorm.DB
	var err error

	switch config.DB_TYPE {
	case "sqlite3":
		db, err = gorm.Open(sqlite.Open("reblog.db"), &gorm.Config{})
	case "mysql":
		db, err = gorm.Open(mysql.Open(config.DB_DSN), &gorm.Config{})
	// case "mongodb":
	// 	db, err = gorm.Open(mongodb.Open(config.DB_URI), &gorm.Config{})
	// Q: 为什么没有MongoDB? A: 因为MongoDB无法进行自动迁移!!!!!!!!!!
	case "postgres":
		db, err = gorm.Open(postgres.Open(config.DB_DSN), &gorm.Config{})
	default:
		panic("不支持的数据库类型")
	}

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
