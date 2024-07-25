package db

import (
	"fmt"
	"log"
	"reblog/internal/config"
	"reblog/internal/model"

	"gorm.io/driver/mysql"
	// "gorm.io/driver/mongodb"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var dbInstance *gorm.DB

func init() {
	dbInstance = NewDB()
}

func NewDB() *gorm.DB {
	config := config.NewFromFile().DB

	var db *gorm.DB
	var err error

	gormLogger := &GormLogger{}

	gormConfig := &gorm.Config{
		Logger: gormLogger,
	}

	switch config.Type {
	case "sqlite3":
		dsn := config.Name
		db, err = gorm.Open(sqlite.Open(dsn), gormConfig)
	case "mysql":
		dsn := fmt.Sprint(config.User, ":", config.Password, "@tcp(", config.Host, ":", config.Port, ")/", config.Name, "?charset=utf8mb4&parseTime=True&loc=Local")
		db, err = gorm.Open(mysql.Open(dsn), gormConfig)
	// case "mongodb":
	// 	db, err = gorm.Open(mongodb.Open(config.DB_URI), &gorm.Config{})
	// Q: 为什么没有MongoDB? A: 因为MongoDB无法进行自动迁移!!!!!!!!!!
	case "postgres":
		var sslmode string
		if config.SSL {
			sslmode = "require"
		} else {
			sslmode = "disable"
		}
		dsn := fmt.Sprint(
			"host=", config.Host,
			" port=", config.Port,
			" user=", config.User,
			" password=", config.Password,
			" dbname=", config.Name,
			" sslmode=", sslmode,
		)
		db, err = gorm.Open(postgres.Open(dsn), gormConfig)
	default:
		log.Panic("[DB] 不支持的数据库类型")
	}

	if err != nil {
		log.Panicf("[DB] 无法连接数据库")
	}

	db.AutoMigrate(
		&model.Site{},
		&model.Article{},
		&model.User{},
		&model.Friend{},
	)

	return db
}

func DB() *gorm.DB {
	if dbInstance == nil {
		dbInstance = NewDB()
	}

	return dbInstance
}
